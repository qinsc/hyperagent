package monitor

import (
	"encoding/json"
	"hyperagent/log"
	"hyperagent/util"
	"io"
	"os"
)

const _MONITOR_JSON string = "monitor.json"

type Monitor struct {
	HostId            string `json:"hostId"`
	MonitorIp         string `json:"monitorIp"`
	EtcdIp            string `json:"etcdIp"`
	EtcdPort          string `json:"etcdPort"`
	HeartBeatInterval string `json:"heartBeatInterval"`
	HeartBeatTimeout  string `json:"heartBeatTimeout"`
	GuacdIp           string `json:"guacdIp"`
	GuacdPort         string `json:"guacdPort"`
}

var monitor *Monitor

func GetMonitor() *Monitor {
	return monitor
}

func SaveMonitor(m *Monitor) {
	file, err := os.OpenFile(_MONITOR_JSON, os.O_CREATE|os.O_TRUNC, 0666) //打开文件
	if err != nil {
		log.Error("Error while open monitor.json to write")
		return
	}
	_, err = io.WriteString(file, util.ToJson(m)) //写入文件(字符串)
	if err != nil {
		log.Error("Error while write monitor.json")
		return
	}
	LoadMonitorConfig()
}

func RemoveMonitor() {
	log.Info("Delete monitor.json ..")
	os.Remove(_MONITOR_JSON)
	monitor = nil
}

func LoadMonitorConfig() {
	log.Info("Start load monitor config ....")
	if _, err := os.Stat(_MONITOR_JSON); os.IsNotExist(err) {
		log.Warn("Monitor.json is not exists.")
		return
	}

	log.Debug("Open monitor file ....")
	file, err := os.Open(_MONITOR_JSON)
	if err != nil {
		log.Error("Error while open monitor.json")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var m Monitor
	err = decoder.Decode(&m)
	if err != nil {
		log.Error("Error while decode monitor.json")
		return
	}
	log.Debug("Load monitor config = %s", util.ToJson(m))
	monitor = &m
}
