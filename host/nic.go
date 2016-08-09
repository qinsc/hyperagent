package host

import (
	"log"

	net "github.com/shirou/gopsutil/net"
)

type Nic struct {
	Name        string   `json:"name"`
	Mac         string   `json:"mac"`
	Ips         []string `json:"ip"`
	BytesSent   uint64   `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64   `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64   `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64   `json:"packetsRecv"` // number of packets received
	Errin       uint64   `json:"errin"`       // total number of errors while receiving
	Errout      uint64   `json:"errout"`      // total number of errors while sending
	Dropin      uint64   `json:"dropin"`      // total number of incoming packets which were dropped
	Dropout     uint64   `json:"dropout"`     // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
	Fifoin      uint64   `json:"fifoin"`      // total number of FIFO buffers errors while receiving
	Fifoout     uint64   `json:"fifoout"`     // total number of FIFO buffers errors while sending
}

func ListNics() []Nic {
	itfs, err := net.Interfaces()
	if err != nil {
		log.Println("Get nics failed.")
		return nil
	}

	ioCounterMap := mapIOCounters()
	if ioCounterMap == nil {
		return nil
	}

	nics := make([]Nic, len(itfs))
	for _, itf := range itfs {
		ioCounter, ok := ioCounterMap[itf.Name]
		if ok {
			nic := Nic{
				Name:        itf.Name,
				Mac:         itf.HardwareAddr,
				BytesSent:   ioCounter.BytesSent,
				BytesRecv:   ioCounter.BytesRecv,
				PacketsSent: ioCounter.PacketsSent,
				PacketsRecv: ioCounter.PacketsRecv,
				Errin:       ioCounter.Errin,
				Errout:      ioCounter.Errout,
				Dropin:      ioCounter.Dropin,
				Dropout:     ioCounter.Dropout,
				Fifoin:      ioCounter.Fifoin,
				Fifoout:     ioCounter.Fifoout,
			}
			if itf.Addrs != nil {
				nic.Ips = make([]string, len(itf.Addrs))
				for _, addr := range itf.Addrs {
					nic.Ips = append(nic.Ips, addr.Addr)
				}
			}
			nics = append(nics, nic)
		}
	}
	return nics
}

func mapIOCounters() map[string]net.IOCountersStat {
	ioCounters, err := net.IOCounters(true)
	if err != nil {
		log.Println("Get nic iocounter failed.")
		return nil
	}

	ioCounterMap := make(map[string]net.IOCountersStat, len(ioCounters))
	for _, ioCounter := range ioCounters {
		ioCounterMap[ioCounter.Name] = ioCounter
	}
	return ioCounterMap
}
