package host

import (
	"hyperagent/log"
	"strings"

	net "github.com/shirou/gopsutil/net"
)

type NicInfo struct {
	Name string   `json:"name"`
	Mac  string   `json:"mac"`
	Ips  []string `json:"ips"`
	//	BytesSent   uint64   `json:"bytesSent"`
	//	BytesRecv   uint64   `json:"bytesRecv"`
	//	PacketsSent uint64   `json:"packetsSent"`
	//	PacketsRecv uint64   `json:"packetsRecv"`
	//	Errin       uint64   `json:"errin"`
	//	Errout      uint64   `json:"errout"`
	//	Dropin      uint64   `json:"dropin"`
	//	Dropout     uint64   `json:"dropout"`
	//	Fifoin      uint64   `json:"fifoin"`
	//	Fifoout     uint64   `json:"fifoout"`
}

func ListNics() []NicInfo {
	itfs, err := net.Interfaces()
	if err != nil {
		log.Error("Get nics failed.")
		return nil
	}

	//	ioCounterMap := mapIOCounters()
	//	if ioCounterMap == nil {
	//		return nil
	//	}

	nics := make([]NicInfo, 0)
	for _, itf := range itfs {
		if len(itf.Name) == 0 || strings.Contains(itf.Name, "Loopback") {
			continue
		}

		ips := make([]string, 0)
		for _, addr := range itf.Addrs {
			if len(addr.Addr) == 0 || strings.Contains(addr.Addr, ":") {
				continue
			}
			ips = append(ips, addr.Addr)
		}

		log.Debug("Nic = %s", itf)

		nic := NicInfo{
			Name: itf.Name,
			Mac:  itf.HardwareAddr,
			Ips:  ips,
		}

		nic.Name = strings.Replace(nic.Name, "Local Area Connection", "本地连接", -1)

		//		ioCounter, ok := ioCounterMap[itf.Name]
		//		if ok {
		//			nic := NicInfo{
		//				Name:        nic.Name,
		//				Mac:         nic.Mac,
		//				Ips:         nic.Ips,
		//				BytesSent:   ioCounter.BytesSent,
		//				BytesRecv:   ioCounter.BytesRecv,
		//				PacketsSent: ioCounter.PacketsSent,
		//				PacketsRecv: ioCounter.PacketsRecv,
		//				Errin:       ioCounter.Errin,
		//				Errout:      ioCounter.Errout,
		//				Dropin:      ioCounter.Dropin,
		//				Dropout:     ioCounter.Dropout,
		//				Fifoin:      ioCounter.Fifoin,
		//				Fifoout:     ioCounter.Fifoout,
		//			}
		//			if itf.Addrs != nil {
		//				nic.Ips = make([]string, 0)
		//				for _, addr := range itf.Addrs {
		//					nic.Ips = append(nic.Ips, addr.Addr)
		//				}
		//			}
		//		}
		nics = append(nics, nic)
	}
	return nics
}

//func mapIOCounters() map[string]net.IOCountersStat {
//	ioCounters, err := net.IOCounters(true)
//	if err != nil {
//		log.Error("Get nic iocounter failed.")
//		return nil
//	}

//	ioCounterMap := make(map[string]net.IOCountersStat, len(ioCounters))
//	for _, ioCounter := range ioCounters {
//		ioCounterMap[ioCounter.Name] = ioCounter
//		log.Debug("net.IOCounterStat = %s", ioCounter)
//	}
//	return ioCounterMap
//}
