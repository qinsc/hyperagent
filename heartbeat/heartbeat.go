package heartbeat

import (
	"hyperagent/log"
	"hyperagent/monitor"
	"hyperagent/util"
	"strconv"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type HeartBeatInfo struct {
	ManagerIp string
}

var ch chan bool = make(chan bool)
var hbRunning bool = false

func StartHeartBeat() {
	log.Info("Start heartbeat")
	m := monitor.GetMonitor()
	if m != nil && m.EtcdIp != "" {
		log.Info("Etcd ip = %s, port = %s, heartbeat interlval = %s, timeout = %s", m.EtcdIp, m.EtcdPort, m.HeartBeatInterval, m.HeartBeatTimeout)
		startHB()
	} else {
		log.Info("Etcd is not configed, heartbeat is not started.")
	}
}

func startHB() {
	if !hbRunning {
		go heartBeat()
		hbRunning = true
	}
}

func StopHeartBeat() {
	if hbRunning {
		go func() {
			select {
			case ch <- true:
				log.Info("Heart beat stoped")
			case <-time.After(time.Second * 90):
				log.Info("Stop heart beat timeout ...")

			}
		}()
		hbRunning = false
	}
}

func heartBeat() {
	hbInterval, err := strconv.Atoi(monitor.GetMonitor().HeartBeatInterval)
	if err != nil {
		hbInterval = 30
	}
	go sendHBtoEtcd()

HB:
	for {
		select {
		case <-time.After(time.Second * time.Duration(hbInterval)):
			log.Info("Send heart to etcd ...")
			go sendHBtoEtcd()
		case <-ch:
			log.Info("Stop heart beat ...")
			ch <- true
			break HB
		}
	}
}

func sendHBtoEtcd() {
	m := monitor.GetMonitor()
	if m == nil {
		return
	}

	etcdEndPoint := "http://" + m.EtcdIp + ":" + m.EtcdPort
	log.Info("etcdEndPoint = %s", etcdEndPoint)

	cli, err := client.New(client.Config{
		Endpoints:               []string{etcdEndPoint},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	})
	if err != nil {
		log.Error("Etcd client create failed")
		return
	}
	//	defer cli.Close()

	log.Debug("Send to Etcd")

	kAPI := client.NewKeysAPI(cli)

	hb := HeartBeatInfo{
		ManagerIp: m.MonitorIp,
	}

	_, err = kAPI.Set(context.Background(), "/hyper/agent/heartbeat/"+m.HostId, util.ToJson(hb), &client.SetOptions{
		TTL: time.Second * 90,
	})
	if err != nil {
		log.Error("Etcd client put failed, msg = %v", err)
	}
}
