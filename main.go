package main

import (
	//	"hyperagent/host"
	"hyperagent/log"
	"hyperagent/rest"
	//	"hyperagent/util"
	"hyperagent/gui"
	"hyperagent/web"
	"net/http"
)

func main() {
	log.Debug("Start hyperagent ....")
	log.Debug("========================================================================\n")

	go gui.CreateTray()
	startWebServer()

	//	log.Error(util.ToJson(host.GetHostInfo()))

	log.Debug("\n========================================================================")
	log.Debug("Hyeragent stared.")
}

func startWebServer() {
	mux := http.NewServeMux()
	web.HandlerWebSite(mux)
	rest.HandlerRestServices(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Error("ListenAndServe: ", err.Error())
	}
}
