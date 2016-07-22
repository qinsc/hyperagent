package host

import (
	"fmt"
	"net"
)

type NicInfo struct {
	Name    string
	Ip      string
	MacAddr string
	Dns     []string
}

func GetNicInfos() []NicInfo {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Get interface list failed.")
		return nil
	}

	size := len(interfaces)
	nicInfos := make([]NicInfo, size)
	for i := 0; i < size; i++ {
		nicInfos[i] = getNicInfo(interfaces[i])
	}

	return nicInfos
}

func getNicInfo(interFace net.Interface) (nicInfo NicInfo) {
	nicInfo.Name = interFace.Name
	nicInfo.MacAddr = interFace.HardwareAddr.String()

	addrs, err := interFace.Addrs()
	if err != nil {
		fmt.Println("Get interface addrs failed.")
		return
	}

	for i := 0; i < len(addrs); i++ {
		nicInfo.Ip = addrs[i].String()
	}

	return
}
