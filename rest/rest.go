package rest

import (
	"hyperagent/host"
	"hyperagent/log"
	"hyperagent/util"
	"net/http"
	"runtime/debug"
)

func HandlerRestServices(mux *http.ServeMux) {
	log.Debug("HandlerRestServices")

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
		w.Write(util.ToJson(hostInfo))
	}
}
