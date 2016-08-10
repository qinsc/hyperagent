package rest

import (
	"hyperagent/gui"
	"hyperagent/host"
	"hyperagent/log"
	"hyperagent/util"
	"net/http"
	"runtime/debug"
)

func HandlerRestServices(mux *http.ServeMux) {
	log.Debug("HandlerRestServices")
	mux.HandleFunc("/rest/host/info", safeHandlerRest(getHostInfo))
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

func getHostInfo(w http.ResponseWriter, r *http.Request) {
	hostInfo := host.GetHostInfo()
	w.Header().Set("Content-Type", "json")
	if hostInfo != nil {
		w.Write([]byte(util.ToJson(hostInfo)))
	}
}

func logoffHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("Logoff host ...")
	host.LogoffHost()
}

func shutdownHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("Shutdown host ...")
	host.ShutdownHost()
}

func rebootHost(w http.ResponseWriter, r *http.Request) {
	log.Debug("ShowHello host ...")
	host.RebootHost()
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	log.Debug("Show Message ...")
	gui.ShowMessage("来自世界的恶意")
}
