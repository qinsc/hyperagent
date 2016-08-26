package rest

import (
	"encoding/json"
	"hyperagent/gui"
	"hyperagent/heartbeat"
	"hyperagent/host"
	"hyperagent/log"
	"hyperagent/monitor"
	"hyperagent/util"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	//"time"
)

func HandlerRestServices(mux *http.ServeMux) {
	log.Debug("HandlerRestServices")
	mux.HandleFunc("/rest/host/config", safeHandlerRest(getHostConfig))
	mux.HandleFunc("/rest/host/info", safeHandlerRest(getHostInfo))
	mux.HandleFunc("/rest/host/add", safeHandlerRest(addHost))
	mux.HandleFunc("/rest/host/remove", safeHandlerRest(removeHost))
	mux.HandleFunc("/rest/host/logoff", safeHandlerRest(logoffHost))
	mux.HandleFunc("/rest/host/shutdown", safeHandlerRest(shutdownHost))
	mux.HandleFunc("/rest/host/reboot", safeHandlerRest(rebootHost))
	mux.HandleFunc("/rest/host/message", safeHandlerRest(sendMessage))
}

func safeHandlerRest(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, "RestError", http.StatusInternalServerError)
				log.Warn("WARN: panic in %v. - %v", fn, e)
				log.Debug(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func getHostConfig(w http.ResponseWriter, r *http.Request) {
	log.Debug("do getHostConfig, Method = %s", r.Method)
	if r.Method == "GET" {
		hostName := host.GetHostName()
		log.Debug("hostname = %s", hostName)
		m := monitor.GetMonitor()
		log.Debug("monitor = %s", util.ToJson(m))

		var hostConfig = host.HostConfig{
			HostName: hostName,
			Monitor:  m,
		}
		w.Write([]byte(util.ToJson(hostConfig)))
	}
}

func addHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("do addHost, Method = %s", r.Method)
	if r.Method == "POST" {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		log.Debug("body = %s", string(body))

		var m monitor.Monitor
		err = json.Unmarshal(body, &m)
		if err != nil {
			panic(err)
		}
		log.Debug("do addHost, monitor = %s", util.ToJson(m))
		monitor.SaveMonitor(&m)

		hostInfo := host.GetHostInfo()
		w.Header().Set("Content-Type", "json")
		if hostInfo != nil {
			w.Write([]byte(util.ToJson(hostInfo)))
		}

		heartbeat.StartHeartBeat()
	}
}

func removeHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("do removeHost = %s", r.Method)
	if r.Method == "POST" {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		hostId := string(body)
		log.Debug("hostId = %s", hostId)

		m := monitor.GetMonitor()
		if m != nil {
			log.Debug("m.hostId = %s", m.HostId)

			if m.HostId == hostId {
				monitor.RemoveMonitor()
				heartbeat.StopHeartBeat()
			}
		}
	}
}

func getHostInfo(w http.ResponseWriter, r *http.Request) {
	log.Debug("do getHostInfo, Method = %s", r.Method)
	if r.Method == "GET" {
		hostInfo := host.GetHostInfo()
		w.Header().Set("Content-Type", "json")
		if hostInfo != nil {
			w.Write([]byte(util.ToJson(hostInfo)))
		}
	}
}

func logoffHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("do logoffHost, Method = %s", r.Method)
	if r.Method == "POST" {
		log.Debug("Method is post, do logoff")
		go func() {
			//gui.ShowMessage("30秒后系统将注销", false)
			//time.Sleep(30 * time.Second)
			host.LogoffHost()
		}()
	}
}

func shutdownHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("do shutdownHost, Method = %s", r.Method)
	if r.Method == "POST" {
		go func() {
			//gui.ShowMessageAll("30秒后系统将关闭")
			//time.Sleep(30 * time.Second)
			host.ShutdownHost()
		}()
	}
}

func rebootHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("do rebootHost, Method = %s", r.Method)
	if r.Method == "POST" {
		go func() {
			//gui.ShowMessageAll("30秒后系统将重启")
			//time.Sleep(30 * time.Second)
			host.RebootHost()
		}()
	}
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	log.Debug("do sendMessage, Method = %s", r.Method)
	if r.Method == "POST" {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		msg := string(body)
		log.Debug("msg = %s", msg)

		go func() {
			log.Debug("Show Message ...")
			gui.ShowMessageAll(msg)
		}()
	}
}
