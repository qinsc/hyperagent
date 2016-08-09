package main

import (
	"fmt"
	"hyperagent/host"
)

func main() {
	fmt.Println("Start hyperagent ....")
	fmt.Println("========================================================================\n")

	// fmt.Println("HostName: ", host.GetHostName())
	// fmt.Println("NicInfos")
	// showNetworkInfo()
	fmt.Println(host.HostInfo().String())

	fmt.Println("\n========================================================================")
	fmt.Println("Hyeragent stared.")
}

//func getHostInfo() (h host.Host) {
//	hostInfo, _ := srhost.Info()
//	h.HostName = hostInfo.Hostname
//	h.OS = hostInfo.OS
//	h.Platform = hostInfo.Platform
//	h.BootTime = hostInfo.BootTime
//	h.Uptime = hostInfo.Uptime
//	fmt.Println(hostInfo)
//	return
//}
