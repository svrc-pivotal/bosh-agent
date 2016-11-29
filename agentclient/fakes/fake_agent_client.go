// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/agentclient"
	"github.com/cloudfoundry/bosh-agent/agentclient/applyspec"
	"github.com/cloudfoundry/bosh-agent/settings"
)

type FakeAgentClient struct {
	PingStub        func() (string, error)
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 string
		result2 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
	ApplyStub        func(applyspec.ApplySpec) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 applyspec.ApplySpec
	}
	applyReturns struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	GetStateStub        func() (agentclient.AgentState, error)
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct{}
	getStateReturns     struct {
		result1 agentclient.AgentState
		result2 error
	}
	MountDiskStub        func(string) error
	mountDiskMutex       sync.RWMutex
	mountDiskArgsForCall []struct {
		arg1 string
	}
	mountDiskReturns struct {
		result1 error
	}
	UnmountDiskStub        func(string) error
	unmountDiskMutex       sync.RWMutex
	unmountDiskArgsForCall []struct {
		arg1 string
	}
	unmountDiskReturns struct {
		result1 error
	}
	ListDiskStub        func() ([]string, error)
	listDiskMutex       sync.RWMutex
	listDiskArgsForCall []struct{}
	listDiskReturns     struct {
		result1 []string
		result2 error
	}
	MigrateDiskStub        func() error
	migrateDiskMutex       sync.RWMutex
	migrateDiskArgsForCall []struct{}
	migrateDiskReturns     struct {
		result1 error
	}
	CompilePackageStub        func(packageSource agentclient.BlobRef, compiledPackageDependencies []agentclient.BlobRef) (compiledPackageRef agentclient.BlobRef, err error)
	compilePackageMutex       sync.RWMutex
	compilePackageArgsForCall []struct {
		packageSource               agentclient.BlobRef
		compiledPackageDependencies []agentclient.BlobRef
	}
	compilePackageReturns struct {
		result1 agentclient.BlobRef
		result2 error
	}
	DeleteARPEntriesStub        func(ips []string) error
	deleteARPEntriesMutex       sync.RWMutex
	deleteARPEntriesArgsForCall []struct {
		ips []string
	}
	deleteARPEntriesReturns struct {
		result1 error
	}
	SyncDNSStub        func(blobID, sha1 string, version uint64) (string, error)
	syncDNSMutex       sync.RWMutex
	syncDNSArgsForCall []struct {
		blobID  string
		sha1    string
		version uint64
	}
	syncDNSReturns struct {
		result1 string
		result2 error
	}
	UpdateSettingsStub        func(settings settings.UpdateSettings) error
	updateSettingsMutex       sync.RWMutex
	updateSettingsArgsForCall []struct {
		settings settings.UpdateSettings
	}
	updateSettingsReturns struct {
		result1 error
	}
	RunScriptStub        func(scriptName string, options map[string]interface{}) error
	runScriptMutex       sync.RWMutex
	runScriptArgsForCall []struct {
		scriptName string
		options    map[string]interface{}
	}
	runScriptReturns struct {
		result1 error
	}
	SSHStub        func(username string) error
	sSHMutex       sync.RWMutex
	sSHArgsForCall []struct {
		username string
	}
	sSHReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAgentClient) Ping() (string, error) {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1, fake.pingReturns.result2
	}
}

func (fake *FakeAgentClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeAgentClient) PingReturns(result1 string, result2 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) Stop() error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeAgentClient) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeAgentClient) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Apply(arg1 applyspec.ApplySpec) error {
	fake.applyMutex.Lock()
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 applyspec.ApplySpec
	}{arg1})
	fake.recordInvocation("Apply", []interface{}{arg1})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(arg1)
	} else {
		return fake.applyReturns.result1
	}
}

func (fake *FakeAgentClient) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeAgentClient) ApplyArgsForCall(i int) applyspec.ApplySpec {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.applyArgsForCall[i].arg1
}

func (fake *FakeAgentClient) ApplyReturns(result1 error) {
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Start() error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeAgentClient) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeAgentClient) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) GetState() (agentclient.AgentState, error) {
	fake.getStateMutex.Lock()
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct{}{})
	fake.recordInvocation("GetState", []interface{}{})
	fake.getStateMutex.Unlock()
	if fake.GetStateStub != nil {
		return fake.GetStateStub()
	} else {
		return fake.getStateReturns.result1, fake.getStateReturns.result2
	}
}

func (fake *FakeAgentClient) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *FakeAgentClient) GetStateReturns(result1 agentclient.AgentState, result2 error) {
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 agentclient.AgentState
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) MountDisk(arg1 string) error {
	fake.mountDiskMutex.Lock()
	fake.mountDiskArgsForCall = append(fake.mountDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("MountDisk", []interface{}{arg1})
	fake.mountDiskMutex.Unlock()
	if fake.MountDiskStub != nil {
		return fake.MountDiskStub(arg1)
	} else {
		return fake.mountDiskReturns.result1
	}
}

func (fake *FakeAgentClient) MountDiskCallCount() int {
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	return len(fake.mountDiskArgsForCall)
}

func (fake *FakeAgentClient) MountDiskArgsForCall(i int) string {
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	return fake.mountDiskArgsForCall[i].arg1
}

func (fake *FakeAgentClient) MountDiskReturns(result1 error) {
	fake.MountDiskStub = nil
	fake.mountDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) UnmountDisk(arg1 string) error {
	fake.unmountDiskMutex.Lock()
	fake.unmountDiskArgsForCall = append(fake.unmountDiskArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UnmountDisk", []interface{}{arg1})
	fake.unmountDiskMutex.Unlock()
	if fake.UnmountDiskStub != nil {
		return fake.UnmountDiskStub(arg1)
	} else {
		return fake.unmountDiskReturns.result1
	}
}

func (fake *FakeAgentClient) UnmountDiskCallCount() int {
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	return len(fake.unmountDiskArgsForCall)
}

func (fake *FakeAgentClient) UnmountDiskArgsForCall(i int) string {
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	return fake.unmountDiskArgsForCall[i].arg1
}

func (fake *FakeAgentClient) UnmountDiskReturns(result1 error) {
	fake.UnmountDiskStub = nil
	fake.unmountDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) ListDisk() ([]string, error) {
	fake.listDiskMutex.Lock()
	fake.listDiskArgsForCall = append(fake.listDiskArgsForCall, struct{}{})
	fake.recordInvocation("ListDisk", []interface{}{})
	fake.listDiskMutex.Unlock()
	if fake.ListDiskStub != nil {
		return fake.ListDiskStub()
	} else {
		return fake.listDiskReturns.result1, fake.listDiskReturns.result2
	}
}

func (fake *FakeAgentClient) ListDiskCallCount() int {
	fake.listDiskMutex.RLock()
	defer fake.listDiskMutex.RUnlock()
	return len(fake.listDiskArgsForCall)
}

func (fake *FakeAgentClient) ListDiskReturns(result1 []string, result2 error) {
	fake.ListDiskStub = nil
	fake.listDiskReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) MigrateDisk() error {
	fake.migrateDiskMutex.Lock()
	fake.migrateDiskArgsForCall = append(fake.migrateDiskArgsForCall, struct{}{})
	fake.recordInvocation("MigrateDisk", []interface{}{})
	fake.migrateDiskMutex.Unlock()
	if fake.MigrateDiskStub != nil {
		return fake.MigrateDiskStub()
	} else {
		return fake.migrateDiskReturns.result1
	}
}

func (fake *FakeAgentClient) MigrateDiskCallCount() int {
	fake.migrateDiskMutex.RLock()
	defer fake.migrateDiskMutex.RUnlock()
	return len(fake.migrateDiskArgsForCall)
}

func (fake *FakeAgentClient) MigrateDiskReturns(result1 error) {
	fake.MigrateDiskStub = nil
	fake.migrateDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) CompilePackage(packageSource agentclient.BlobRef, compiledPackageDependencies []agentclient.BlobRef) (compiledPackageRef agentclient.BlobRef, err error) {
	var compiledPackageDependenciesCopy []agentclient.BlobRef
	if compiledPackageDependencies != nil {
		compiledPackageDependenciesCopy = make([]agentclient.BlobRef, len(compiledPackageDependencies))
		copy(compiledPackageDependenciesCopy, compiledPackageDependencies)
	}
	fake.compilePackageMutex.Lock()
	fake.compilePackageArgsForCall = append(fake.compilePackageArgsForCall, struct {
		packageSource               agentclient.BlobRef
		compiledPackageDependencies []agentclient.BlobRef
	}{packageSource, compiledPackageDependenciesCopy})
	fake.recordInvocation("CompilePackage", []interface{}{packageSource, compiledPackageDependenciesCopy})
	fake.compilePackageMutex.Unlock()
	if fake.CompilePackageStub != nil {
		return fake.CompilePackageStub(packageSource, compiledPackageDependencies)
	} else {
		return fake.compilePackageReturns.result1, fake.compilePackageReturns.result2
	}
}

func (fake *FakeAgentClient) CompilePackageCallCount() int {
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	return len(fake.compilePackageArgsForCall)
}

func (fake *FakeAgentClient) CompilePackageArgsForCall(i int) (agentclient.BlobRef, []agentclient.BlobRef) {
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	return fake.compilePackageArgsForCall[i].packageSource, fake.compilePackageArgsForCall[i].compiledPackageDependencies
}

func (fake *FakeAgentClient) CompilePackageReturns(result1 agentclient.BlobRef, result2 error) {
	fake.CompilePackageStub = nil
	fake.compilePackageReturns = struct {
		result1 agentclient.BlobRef
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) DeleteARPEntries(ips []string) error {
	var ipsCopy []string
	if ips != nil {
		ipsCopy = make([]string, len(ips))
		copy(ipsCopy, ips)
	}
	fake.deleteARPEntriesMutex.Lock()
	fake.deleteARPEntriesArgsForCall = append(fake.deleteARPEntriesArgsForCall, struct {
		ips []string
	}{ipsCopy})
	fake.recordInvocation("DeleteARPEntries", []interface{}{ipsCopy})
	fake.deleteARPEntriesMutex.Unlock()
	if fake.DeleteARPEntriesStub != nil {
		return fake.DeleteARPEntriesStub(ips)
	} else {
		return fake.deleteARPEntriesReturns.result1
	}
}

func (fake *FakeAgentClient) DeleteARPEntriesCallCount() int {
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	return len(fake.deleteARPEntriesArgsForCall)
}

func (fake *FakeAgentClient) DeleteARPEntriesArgsForCall(i int) []string {
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	return fake.deleteARPEntriesArgsForCall[i].ips
}

func (fake *FakeAgentClient) DeleteARPEntriesReturns(result1 error) {
	fake.DeleteARPEntriesStub = nil
	fake.deleteARPEntriesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) SyncDNS(blobID string, sha1 string, version uint64) (string, error) {
	fake.syncDNSMutex.Lock()
	fake.syncDNSArgsForCall = append(fake.syncDNSArgsForCall, struct {
		blobID  string
		sha1    string
		version uint64
	}{blobID, sha1, version})
	fake.recordInvocation("SyncDNS", []interface{}{blobID, sha1, version})
	fake.syncDNSMutex.Unlock()
	if fake.SyncDNSStub != nil {
		return fake.SyncDNSStub(blobID, sha1, version)
	} else {
		return fake.syncDNSReturns.result1, fake.syncDNSReturns.result2
	}
}

func (fake *FakeAgentClient) SyncDNSCallCount() int {
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	return len(fake.syncDNSArgsForCall)
}

func (fake *FakeAgentClient) SyncDNSArgsForCall(i int) (string, string, uint64) {
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	return fake.syncDNSArgsForCall[i].blobID, fake.syncDNSArgsForCall[i].sha1, fake.syncDNSArgsForCall[i].version
}

func (fake *FakeAgentClient) SyncDNSReturns(result1 string, result2 error) {
	fake.SyncDNSStub = nil
	fake.syncDNSReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAgentClient) UpdateSettings(settingsA settings.UpdateSettings) error {
	fake.updateSettingsMutex.Lock()
	fake.updateSettingsArgsForCall = append(fake.updateSettingsArgsForCall, struct {
		settings settings.UpdateSettings
	}{settingsA})
	fake.recordInvocation("UpdateSettings", []interface{}{settingsA})
	fake.updateSettingsMutex.Unlock()
	if fake.UpdateSettingsStub != nil {
		return fake.UpdateSettingsStub(settingsA)
	} else {
		return fake.updateSettingsReturns.result1
	}
}

func (fake *FakeAgentClient) UpdateSettingsCallCount() int {
	fake.updateSettingsMutex.RLock()
	defer fake.updateSettingsMutex.RUnlock()
	return len(fake.updateSettingsArgsForCall)
}

func (fake *FakeAgentClient) UpdateSettingsArgsForCall(i int) settings.UpdateSettings {
	fake.updateSettingsMutex.RLock()
	defer fake.updateSettingsMutex.RUnlock()
	return fake.updateSettingsArgsForCall[i].settings
}

func (fake *FakeAgentClient) UpdateSettingsReturns(result1 error) {
	fake.UpdateSettingsStub = nil
	fake.updateSettingsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) RunScript(scriptName string, options map[string]interface{}) error {
	fake.runScriptMutex.Lock()
	fake.runScriptArgsForCall = append(fake.runScriptArgsForCall, struct {
		scriptName string
		options    map[string]interface{}
	}{scriptName, options})
	fake.recordInvocation("RunScript", []interface{}{scriptName, options})
	fake.runScriptMutex.Unlock()
	if fake.RunScriptStub != nil {
		return fake.RunScriptStub(scriptName, options)
	} else {
		return fake.runScriptReturns.result1
	}
}

func (fake *FakeAgentClient) RunScriptCallCount() int {
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	return len(fake.runScriptArgsForCall)
}

func (fake *FakeAgentClient) RunScriptArgsForCall(i int) (string, map[string]interface{}) {
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	return fake.runScriptArgsForCall[i].scriptName, fake.runScriptArgsForCall[i].options
}

func (fake *FakeAgentClient) RunScriptReturns(result1 error) {
	fake.RunScriptStub = nil
	fake.runScriptReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) SSH(username string) error {
	fake.sSHMutex.Lock()
	fake.sSHArgsForCall = append(fake.sSHArgsForCall, struct {
		username string
	}{username})
	fake.recordInvocation("SSH", []interface{}{username})
	fake.sSHMutex.Unlock()
	if fake.SSHStub != nil {
		return fake.SSHStub(username)
	} else {
		return fake.sSHReturns.result1
	}
}

func (fake *FakeAgentClient) SSHCallCount() int {
	fake.sSHMutex.RLock()
	defer fake.sSHMutex.RUnlock()
	return len(fake.sSHArgsForCall)
}

func (fake *FakeAgentClient) SSHArgsForCall(i int) string {
	fake.sSHMutex.RLock()
	defer fake.sSHMutex.RUnlock()
	return fake.sSHArgsForCall[i].username
}

func (fake *FakeAgentClient) SSHReturns(result1 error) {
	fake.SSHStub = nil
	fake.sSHReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAgentClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	fake.mountDiskMutex.RLock()
	defer fake.mountDiskMutex.RUnlock()
	fake.unmountDiskMutex.RLock()
	defer fake.unmountDiskMutex.RUnlock()
	fake.listDiskMutex.RLock()
	defer fake.listDiskMutex.RUnlock()
	fake.migrateDiskMutex.RLock()
	defer fake.migrateDiskMutex.RUnlock()
	fake.compilePackageMutex.RLock()
	defer fake.compilePackageMutex.RUnlock()
	fake.deleteARPEntriesMutex.RLock()
	defer fake.deleteARPEntriesMutex.RUnlock()
	fake.syncDNSMutex.RLock()
	defer fake.syncDNSMutex.RUnlock()
	fake.updateSettingsMutex.RLock()
	defer fake.updateSettingsMutex.RUnlock()
	fake.runScriptMutex.RLock()
	defer fake.runScriptMutex.RUnlock()
	fake.sSHMutex.RLock()
	defer fake.sSHMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAgentClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ agentclient.AgentClient = new(FakeAgentClient)
