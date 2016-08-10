package host

import (
	"encoding/json"
	"hyperagent/log"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type Host struct {
	Hostname          string  `json:"hostname"`
	Uptime            uint64  `json:"upTime"`
	BootTime          uint64  `json:"bootTime"`
	Procs             uint64  `json:"procs"`
	OS                string  `json:"os"`
	OSPlatform        string  `json:"osPlatform"`
	OSPlatformFamily  string  `json:"osPlatformFamily"`
	OSPlatformVersion string  `json:"osPlatformVersion"`
	CPUCores          int32   `json:"cpuCores"`
	CPUModelName      string  `json:"cpuModelName"`
	CPUMhz            float64 `json:"cpuMhz"`
	CPUUsage          float64 `json:"cpuUsage"`
	MemTotal          uint64  `json:"memSize"`
	MemUsedPercent    float64 `json:"memUsage"`
	Nics              []Nic   `json:"nicInfos"`
	Disks             []Disk  `json:"diskInfos"`
}

func (h *Host) String() string {
	s, _ := json.Marshal(h)
	return string(s)
}

func GetHostInfo() *Host {
	log.Debug("Start get HostInfo")
	log.Debug("Call host.Info()")
	hostInfo, err := host.Info()
	if err != nil {
		log.Debug("Get host info failed.")
		return nil
	}

	log.Debug("Call cpu.Info()")
	cpuInfos, err := cpu.Info()
	if err != nil {
		log.Debug("Get cpu info failed.")
		return nil
	}

	var cpuCores int32
	var cpuMhz float64
	var cpuModelName string
	for _, cpuInfo := range cpuInfos {
		cpuCores += cpuInfo.Cores
		cpuMhz = cpuInfo.Mhz
		cpuModelName = cpuInfo.ModelName
	}

	log.Debug("Call cpu.Percent()")
	cpuUsages, err := cpu.Percent(time.Second*5, false)
	if err != nil {
		log.Debug("Get cpu usage failed.")
		return nil
	}

	log.Debug("Call VirtualMemory()")
	memoryInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Debug("Get memory info failed.")
		return nil
	}

	log.Debug("Call ListNics()")
	nics := ListNics()

	log.Debug("Call ListDisks()")
	disks := ListDisks()

	log.Debug("Return HostInfo")
	return &Host{
		Hostname:          hostInfo.Hostname,
		Uptime:            hostInfo.Uptime,
		BootTime:          hostInfo.BootTime,
		Procs:             hostInfo.Procs,
		OS:                hostInfo.OS,
		OSPlatform:        hostInfo.Platform,
		OSPlatformFamily:  hostInfo.PlatformFamily,
		OSPlatformVersion: hostInfo.PlatformVersion,
		CPUCores:          cpuCores,
		CPUMhz:            cpuMhz,
		CPUModelName:      cpuModelName,
		CPUUsage:          cpuUsages[0],
		MemTotal:          memoryInfo.Total,
		MemUsedPercent:    memoryInfo.UsedPercent,
		Nics:              nics,
		Disks:             disks,
	}
}
