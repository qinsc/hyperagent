package host

import (
	"hyperagent/log"
	"hyperagent/monitor"
	"hyperagent/util"
	"io/ioutil"
	"os/exec"
	"time"

	. "github.com/CodyGuo/win"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type HostInfo struct {
	Hostname          string     `json:"hostName"`
	Uptime            uint64     `json:"upTime"`
	BootTime          uint64     `json:"bootTime"`
	OS                string     `json:"os"`
	OSPlatform        string     `json:"osPlatform"`
	OSPlatformFamily  string     `json:"osPlatformFamily"`
	OSPlatformVersion string     `json:"osPlatformVersion"`
	CPUCores          int32      `json:"cpuCores"`
	CPUModelName      string     `json:"cpuModelName"`
	CPUMhz            float64    `json:"cpuMhz"`
	CPUUsage          float64    `json:"cpuUsage"`
	MemTotal          uint64     `json:"memSize"`
	MemUsed           uint64     `json:"memUsed"`
	MemUsedPercent    float64    `json:"memUsage"`
	Nics              []NicInfo  `json:"nicInfos"`
	Disks             []DiskInfo `json:"diskInfos"`
}

type HostConfig struct {
	HostName string           `json:"hostName"`
	Monitor  *monitor.Monitor `json:"monitor"`
}

func init() {
	GetHostDetailInfo()
}

func GetHostDetailInfo() string {
	cmd := exec.Command("systeminfo")
	stdout, err := cmd.StdoutPipe()
	cmd.Start()
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Error("Error while get host detailinfo, err = " + util.ToJson(err))
		return ""
	}
	// log.Debug(string(content))
	return string(content)
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
		OS:                hostInfo.OS,
		OSPlatform:        hostInfo.Platform,
		OSPlatformFamily:  hostInfo.PlatformFamily,
		OSPlatformVersion: hostInfo.PlatformVersion,
		CPUCores:          cpuCores,
		CPUMhz:            cpuMhz,
		CPUModelName:      cpuModelName,
		CPUUsage:          cpuUsages[0],
		MemTotal:          memoryInfo.Total,
		MemUsed:           memoryInfo.Used,
		MemUsedPercent:    memoryInfo.UsedPercent,
		Nics:              nics,
		Disks:             disks,
	}
}

func GetHostName() string {
	hostInfo, err := host.Info()
	if err != nil {
		log.Debug("Get host info failed.")
		return "Unkown hostName"
	}
	return hostInfo.Hostname
}

func LogoffHost() {
	ExitWindowsEx(EWX_LOGOFF, 0)
}

func ShutdownHost() {
	getPrivileges()
	ExitWindowsEx(EWX_SHUTDOWN|EWX_FORCE, 0)
}

func RebootHost() {
	getPrivileges()
	ExitWindowsEx(EWX_REBOOT|EWX_FORCE, 0)
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
