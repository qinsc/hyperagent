package heartbeat

import (
	"hyperagent/log"
	"hyperagent/monitor"
	"hyperagent/util"
	"sync"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

var (
	tick = time.Tick(time.Second)
	hbCH = make(chan int)
	once sync.Once
)

const (
	hb_start = iota
	hb_stop
	hb_quit
)

func runHB(ch chan int) {
loop:
	for {
		select {
		case c := <-ch:
			log.Debug("Get control signal")
			if c == hb_start {
				log.Debug("to start hb")
				StartHB()
			} else if c == hb_stop {
				log.Debug("to stop hb")
				StopHB()
			} else {
				log.Debug("to quit hb")
				QuitHB()
				break loop
			}
		}
	}

	log.Debug("HB controller goroutine finished.")
}

func StartHB() {
	log.Debug("startHB")
	once.Do(doHB)
	hbCH <- hb_start
}

func StopHB() {
	log.Debug("stopHB")
	hbCH <- hb_stop
}

func QuitHB() {
	log.Debug("quitHB")
	hbCH <- hb_quit
}

func doHB() {
	go func() {
		send := true
	loop:
		for {
			select {
			case <-tick:
				log.Debug(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
				if send {
					go sendHB()
				} else {
					log.Debug("not send ...")
				}
			case sendHB := <-hbCH:
				if sendHB == hb_start {
					log.Debug("Set send to true")
					send = true
				} else if sendHB == hb_stop {
					log.Debug("Set send to false")
					send = false
				} else {
					log.Debug("break send loop")
					break loop
				}
				continue loop
			}
		}

		log.Debug("HB worker finished")
	}()
}

type HBInfo struct {
	ManagerIp string
}

func sendHB() {
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

	log.Debug("Send to Etcd")

	kAPI := client.NewKeysAPI(cli)

	hb := HBInfo{
		ManagerIp: m.MonitorIp,
	}

	_, err = kAPI.Set(context.Background(), "/hyper/agent/heartbeat/"+m.HostId, util.ToJson(hb), &client.SetOptions{
		TTL: time.Second * 90,
	})
	if err != nil {
		log.Error("Etcd client put failed, msg = %v", err)
	}
}
