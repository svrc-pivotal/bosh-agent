package devicepathresolver

import (
	"fmt"
	"path"
	"strings"
	"time"
	"strconv"

	boshsettings "github.com/cloudfoundry/bosh-agent/settings"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

)

const maxScanRetries = 30

type SCSIVolumeIDDevicePathResolver struct {
	diskWaitTimeout time.Duration
	fs              boshsys.FileSystem

	logger boshlog.Logger
	logTag string

}

func NewSCSIVolumeIDDevicePathResolver(
	diskWaitTimeout time.Duration,
	fs boshsys.FileSystem,
	logger boshlog.Logger,
) SCSIVolumeIDDevicePathResolver {
	return SCSIVolumeIDDevicePathResolver{
		fs:              fs,
		diskWaitTimeout: diskWaitTimeout,
		logTag: "scsiVolIdResolver",
		logger: logger,
	}
}

func (devicePathResolver SCSIVolumeIDDevicePathResolver) GetRealDevicePath(diskSettings boshsettings.DiskSettings) (realPath string, timedOut bool, err error) {
	devicePaths, err := devicePathResolver.fs.Glob("/sys/bus/scsi/devices/*:0:0:0/block/*")
	devicePathResolver.logger.Info(devicePathResolver.logTag, "Detected %d SCSI hosts", len(devicePaths))
	if err != nil {
		return
	}

	var hostID string
	volumeID := diskSettings.VolumeID

	// Assumes the lowest SCSI hostID greater than 0 is the primary controller for the VM and houses all BOSH root/ephemeral/persistent disks 
	for _, rootDevicePath := range devicePaths {
		rootDevicePathSplits := strings.Split(rootDevicePath, "/")
		if len(rootDevicePathSplits) > 5 {
			scsiPath := rootDevicePathSplits[5]
			scsiPathSplits := strings.Split(scsiPath, ":")
			if len(scsiPathSplits) > 0 {
				devicePathResolver.logger.Debug(devicePathResolver.logTag, "Evaluating SCSI host %s", scsiPathSplits[0])
				if (len(hostID) == 0) {
					newHostID, err := strconv.Atoi(scsiPathSplits[0])
					if (err == nil) {
						if (newHostID > 0) {
							hostID = scsiPathSplits[0] 
						}
					}
				} else {
					lastHostID, err := strconv.Atoi(hostID)
					if (err == nil) {
						newHostID, err := strconv.Atoi(scsiPathSplits[0])
						if (err == nil) {
							if (newHostID < lastHostID) {
								hostID = scsiPathSplits[0] 
							}
						}
					}
				}				
			}
		}		
	}
	if len(hostID) == 0 {
		return
	}
    devicePathResolver.logger.Info(devicePathResolver.logTag, "Using SCSI host %s for device path resolution", hostID)

	scanPath := fmt.Sprintf("/sys/class/scsi_host/host%s/scan", hostID)
	err = devicePathResolver.fs.WriteFileString(scanPath, "- - -")
	if err != nil {
		return
	}

	deviceGlobPath := fmt.Sprintf("/sys/bus/scsi/devices/%s:0:%s:0/block/*", hostID, volumeID)

	for i := 0; i < maxScanRetries; i++ {
		devicePaths, err = devicePathResolver.fs.Glob(deviceGlobPath)
		if err != nil || len(devicePaths) == 0 {
			time.Sleep(devicePathResolver.diskWaitTimeout)
			continue
		} else {
			break
		}
	}

	if err != nil || len(devicePaths) == 0 {
		return
	}

	basename := path.Base(devicePaths[0])
	devicePathResolver.logger.Info(devicePathResolver.logTag, "Using SCSI device %s for device path", basename)
	realPath = path.Join("/dev/", basename)

	return
}
