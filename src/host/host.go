package host

import (
	"encoding/json"
)

type Host struct {
	HostName string
	OS       string
	Platform string
	BootTime uint64
	Uptime   uint64
}

func (h Host) String() string {
	s, _ := json.Marshal(h)
	return string(s)
}
