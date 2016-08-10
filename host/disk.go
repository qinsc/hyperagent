package host

import (
	"hyperagent/log"

	disk "github.com/shirou/gopsutil/disk"
)

type DiskInfo struct {
	Path        string  `json:"path"`
	Fstype      string  `json:"fsType"`
	Device      string  `json:"device"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func ListDisks() []DiskInfo {
	partitions, err := disk.Partitions(true)
	if err != nil {
		log.Error("Get disks failed.")
		return nil
	}

	disks := make([]DiskInfo, len(partitions))
	for _, partition := range partitions {
		diskUsage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Error("Get disk " + partition.Mountpoint + " usage failed.")
			continue
		}

		log.Debug("partition: %s", partition)

		disk := DiskInfo{
			Path:        partition.Mountpoint,
			Fstype:      partition.Fstype,
			Device:      partition.Device,
			Total:       diskUsage.Total,
			Used:        diskUsage.Used,
			UsedPercent: diskUsage.UsedPercent,
		}
		disks = append(disks, disk)
	}
	return disks
}
