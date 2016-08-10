package host

import (
	"hyperagent/log"
	"time"

	. "github.com/CodyGuo/win"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type HostInfo struct {
	Hostname          string     `json:"hostname"`
	Uptime            uint64     `json:"upTime"`
	BootTime          uint64     `json:"bootTime"`
	Procs             uint64     `json:"procs"`
	OS                string     `json:"os"`
	OSPlatform        string     `json:"osPlatform"`
	OSPlatformFamily  string     `json:"osPlatformFamily"`
	OSPlatformVersion string     `json:"osPlatformVersion"`
	CPUCores          int32      `json:"cpuCores"`
	CPUModelName      string     `json:"cpuModelName"`
	CPUMhz            float64    `json:"cpuMhz"`
	CPUUsage          float64    `json:"cpuUsage"`
	MemTotal          uint64     `json:"memSize"`
	MemUsedPercent    float64    `json:"memUsage"`
	Nics              []NicInfo  `json:"nicInfos"`
	Disks             []DiskInfo `json:"diskInfos"`
}

func GetHostInfo() *HostInfo {
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
	//	cpuUsages := make([]float64, 1)

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
	return &HostInfo{
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

func LogoffHost() {
	ExitWindowsEx(EWX_LOGOFF, 0)
}

func ShutdownHost() {
	getPrivileges()
	ExitWindowsEx(EWX_SHUTDOWN&EWX_FORCE, 0)
}

func RebootHost() {
	getPrivileges()
	ExitWindowsEx(EWX_REBOOT&EWX_FORCE, 0)
}

func getPrivileges() {
	var hToken HANDLE
	var tkp TOKEN_PRIVILEGES

	OpenProcessToken(GetCurrentProcess(), TOKEN_ADJUST_PRIVILEGES|TOKEN_QUERY, &hToken)
	LookupPrivilegeValueA(nil, StringToBytePtr(SE_SHUTDOWN_NAME), &tkp.Privileges[0].Luid)
	tkp.PrivilegeCount = 1
	tkp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED
	AdjustTokenPrivileges(hToken, false, &tkp, 0, nil, nil)
}
