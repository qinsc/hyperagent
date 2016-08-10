package host

import (
	"hyperagent/log"

	disk "github.com/shirou/gopsutil/disk"
)

type Disk struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fsType"`
	Device            string  `json:"device"`
	Total             uint64  `json:"total"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

func ListDisks() []Disk {
	partitions, err := disk.Partitions(true)
	if err != nil {
		log.Error("Get disks failed.")
		return nil
	}

	disks := make([]Disk, len(partitions))
	for _, partition := range partitions {
		diskUsage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Error("Get disk " + partition.Mountpoint + " usage failed.")
			continue
		}

		disk := Disk{
			Path:              partition.Mountpoint,
			Fstype:            partition.Fstype,
			Device:            partition.Device,
			Total:             diskUsage.Total,
			Used:              diskUsage.Used,
			UsedPercent:       diskUsage.UsedPercent,
			InodesTotal:       diskUsage.InodesTotal,
			InodesUsed:        diskUsage.InodesUsed,
			InodesUsedPercent: diskUsage.InodesUsedPercent,
		}
		disks = append(disks, disk)
	}
	return disks
}
