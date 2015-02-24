package net

import (
	"bytes"
	"path/filepath"
	"strings"
	"text/template"

	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	bosharp "github.com/cloudfoundry/bosh-agent/platform/net/arp"
	boship "github.com/cloudfoundry/bosh-agent/platform/net/ip"
	boshsettings "github.com/cloudfoundry/bosh-agent/settings"
	boshsys "github.com/cloudfoundry/bosh-agent/system"
)

const ubuntuNetManagerLogTag = "ubuntuNetManager"

type ubuntuNetManager struct {
	DefaultNetworkResolver

	cmdRunner          boshsys.CmdRunner
	fs                 boshsys.FileSystem
	ipResolver         boship.Resolver
	addressBroadcaster bosharp.AddressBroadcaster
	logger             boshlog.Logger
}

func NewUbuntuNetManager(
	fs boshsys.FileSystem,
	cmdRunner boshsys.CmdRunner,
	defaultNetworkResolver DefaultNetworkResolver,
	ipResolver boship.Resolver,
	addressBroadcaster bosharp.AddressBroadcaster,
	logger boshlog.Logger,
) Manager {
	return ubuntuNetManager{
		DefaultNetworkResolver: defaultNetworkResolver,
		cmdRunner:              cmdRunner,
		fs:                     fs,
		ipResolver:             ipResolver,
		addressBroadcaster:     addressBroadcaster,
		logger:                 logger,
	}
}

// DHCP Config file - /etc/dhcp/dhclient.conf
// Ubuntu 14.04 accepts several DNS as a list in a single prepend directive
const ubuntuDHCPConfigTemplate = `# Generated by bosh-agent

option rfc3442-classless-static-routes code 121 = array of unsigned integer 8;

send host-name "<hostname>";

request subnet-mask, broadcast-address, time-offset, routers,
	domain-name, domain-name-servers, domain-search, host-name,
	netbios-name-servers, netbios-scope, interface-mtu,
	rfc3442-classless-static-routes, ntp-servers;
{{ if . }}
prepend domain-name-servers {{ . }};{{ end }}
`

type staticInterface struct {
	Name      string
	Address   string
	Netmask   string
	Network   string
	Broadcast string
	Mac       string
	Gateway   string
}

type dhcpInterface struct {
	Name string
}

func (net ubuntuNetManager) SetupNetworking(networks boshsettings.Networks, errCh chan error) error {
	nonVipNetworks := boshsettings.Networks{}
	for networkName, networkSettings := range networks {
		if networkSettings.IsVIP() {
			continue
		}
		nonVipNetworks[networkName] = networkSettings
	}

	staticInterfaces, dhcpInterfaces, err := net.buildInterfaces(nonVipNetworks)
	if err != nil {
		return err
	}

	dnsNetwork, _ := networks.DefaultNetworkFor("dns")
	dnsServers := dnsNetwork.DNS

	interfacesChanged, err := net.writeNetworkInterfaces(dhcpInterfaces, staticInterfaces, dnsServers)
	if err != nil {
		return bosherr.WrapError(err, "Writing network configuration")
	}

	dhcpChanged := false
	if len(dhcpInterfaces) > 0 {
		dhcpChanged, err = net.writeDHCPConfiguration(dnsServers)
		if err != nil {
			return err
		}
	}

	if interfacesChanged || dhcpChanged {
		net.restartNetworkingInterfaces()
	}

	net.broadcastIps(staticInterfaces, dhcpInterfaces, errCh)

	return nil
}

func (net ubuntuNetManager) buildInterfaces(networks boshsettings.Networks) ([]staticInterface, []dhcpInterface, error) {
	var (
		staticInterfaces []staticInterface
		dhcpInterfaces   []dhcpInterface
	)
	interfacesByMacAddress, err := net.detectMacAddresses()
	if err != nil {
		return nil, nil, bosherr.WrapError(err, "Getting network interfaces")
	}

	for _, network := range networks {
		if network.IsDynamic() {
			dhcpInterfaces = append(dhcpInterfaces, dhcpInterface{
				Name: interfacesByMacAddress[network.Mac],
			})
		} else {
			networkAddress, broadcastAddress, err := boshsys.CalculateNetworkAndBroadcast(network.IP, network.Netmask)
			if err != nil {
				return nil, nil, bosherr.WrapError(err, "Calculating Network and Broadcast")
			}
			staticInterfaces = append(staticInterfaces, staticInterface{
				Name:      interfacesByMacAddress[network.Mac],
				Address:   network.IP,
				Netmask:   network.Netmask,
				Network:   networkAddress,
				Broadcast: broadcastAddress,
				Mac:       network.Mac,
				Gateway:   network.Gateway,
			})
		}
	}

	return staticInterfaces, dhcpInterfaces, nil

}

func (net ubuntuNetManager) broadcastIps(staticInterfaces []staticInterface, dhcpInterfaces []dhcpInterface, errCh chan error) {
	addresses := []boship.InterfaceAddress{}
	for _, iface := range staticInterfaces {
		addresses = append(addresses, boship.NewSimpleInterfaceAddress(iface.Name, iface.Address))
	}
	for _, iface := range dhcpInterfaces {
		addresses = append(addresses, boship.NewResolvingInterfaceAddress(iface.Name, net.ipResolver))
	}

	go func() {
		net.addressBroadcaster.BroadcastMACAddresses(addresses)
		if errCh != nil {
			errCh <- nil
		}
	}()
}

func (net ubuntuNetManager) restartNetworkingInterfaces() {
	net.logger.Debug(ubuntuNetManagerLogTag, "Restarting network interfaces")

	_, _, _, err := net.cmdRunner.RunCommand("ifdown", "-a", "--no-loopback")
	if err != nil {
		net.logger.Error(ubuntuNetManagerLogTag, "Ignoring ifdown failure: %s", err.Error())
	}

	_, _, _, err = net.cmdRunner.RunCommand("ifup", "-a", "--no-loopback")
	if err != nil {
		net.logger.Error(ubuntuNetManagerLogTag, "Ignoring ifup failure: %s", err.Error())
	}
}

func (net ubuntuNetManager) writeDHCPConfiguration(dnsServers []string) (bool, error) {
	buffer := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("dhcp-config").Parse(ubuntuDHCPConfigTemplate))

	// Keep DNS servers in the order specified by the network
	// because they are added by a *single* DHCP's prepend command
	dnsServersList := strings.Join(dnsServers, ", ")
	err := t.Execute(buffer, dnsServersList)
	if err != nil {
		return false, bosherr.WrapError(err, "Generating config from template")
	}
	dhclientConfigFile := "/etc/dhcp/dhclient.conf"
	changed, err := net.fs.ConvergeFileContents(dhclientConfigFile, buffer.Bytes())

	if err != nil {
		return changed, bosherr.WrapErrorf(err, "Writing to %s", dhclientConfigFile)
	}

	return changed, nil
}

type networkInterfaceConfig struct {
	DNSServers        []string
	StaticInterfaces  []staticInterface
	DhcpInterfaces    []dhcpInterface
	HasDNSNameServers bool
}

func (net ubuntuNetManager) writeNetworkInterfaces(dhcpInterfaces []dhcpInterface, staticInterfaces []staticInterface, dnsServers []string) (bool, error) {
	networkInterfaceValues := networkInterfaceConfig{
		StaticInterfaces:  staticInterfaces,
		DhcpInterfaces:    dhcpInterfaces,
		HasDNSNameServers: true,
		DNSServers:        dnsServers,
	}

	buffer := bytes.NewBuffer([]byte{})

	t := template.Must(template.New("network-interfaces").Parse(networkInterfacesTemplate))

	err := t.Execute(buffer, networkInterfaceValues)
	if err != nil {
		return false, bosherr.WrapError(err, "Generating config from template")
	}

	changed, err := net.fs.ConvergeFileContents("/etc/network/interfaces", buffer.Bytes())
	if err != nil {
		return changed, bosherr.WrapError(err, "Writing to /etc/network/interfaces")
	}

	return changed, nil
}

const networkInterfacesTemplate = `# Generated by bosh-agent
auto lo
iface lo inet loopback
{{ range .DhcpInterfaces }}auto {{ .Name }}
iface {{ .Name }} inet dhcp{{ end }}
{{ range .StaticInterfaces }}auto {{ .Name }}
iface {{ .Name }} inet static
    address {{ .Address }}
    network {{ .Network }}
    netmask {{ .Netmask }}
    broadcast {{ .Broadcast }}
    gateway {{ .Gateway }}{{ end }}
{{ if .DNSServers }}dns-nameservers{{ range .DNSServers }} {{ . }}{{ end }}{{ end }}`

func (net ubuntuNetManager) detectMacAddresses() (map[string]string, error) {
	addresses := map[string]string{}

	filePaths, err := net.fs.Glob("/sys/class/net/*")
	if err != nil {
		return addresses, bosherr.WrapError(err, "Getting file list from /sys/class/net")
	}

	var macAddress string
	for _, filePath := range filePaths {
		macAddress, err = net.fs.ReadFileString(filepath.Join(filePath, "address"))
		if err != nil {
			return addresses, bosherr.WrapError(err, "Reading mac address from file")
		}

		macAddress = strings.Trim(macAddress, "\n")

		interfaceName := filepath.Base(filePath)
		addresses[macAddress] = interfaceName
	}

	return addresses, nil
}
