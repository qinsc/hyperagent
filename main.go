package main

import (
	"hyperagent/gui"
	"hyperagent/heartbeat"
	"hyperagent/log"
	"hyperagent/rest"
	"hyperagent/web"
	"net/http"
)

func main() {
	log.Debug("Start hyperagent ....")

	go gui.CreateTray()
	heartbeat.StartHeartBeat()
	startWebServer()

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
