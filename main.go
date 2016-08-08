package main

import (
	srhost "github.com/shirou/gopsutil/host"
	"html/template"
	"hyperagent/host"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"runtime/debug"
)

func main() {
	fmt.Println("Start hyperagent ....")
	fmt.Println("========================================================================\n")

	// fmt.Println("HostName: ", host.GetHostName())
	// fmt.Println("NicInfos")
	// showNetworkInfo()
	h := getHostInfo()
	fmt.Println(h)

	fmt.Println("\n========================================================================")
	fmt.Println("Hyeragent stared.")
}

// func showNetworkInfo() {
// 	nicInfos := host.GetNicInfos()
// 	for i := 0; i < len(nicInfos); i++ {
// 		fmt.Println("      ", nicInfos[i])
// 	}
// }

func getHostInfo() (h host.Host) {
	hostInfo, _ := srhost.Info()
	h.HostName = hostInfo.Hostname
	h.OS = hostInfo.OS
	h.Platform = hostInfo.Platform
	h.BootTime = hostInfo.BootTime
	h.Uptime = hostInfo.Uptime
	fmt.Println(hostInfo)
	return
}
