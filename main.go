package main

import (
	"hyperagent/log"
	"hyperagent/rest"
	"hyperagent/web"
	"net/http"
)

func main() {
	log.Debug("Start hyperagent ....")
	log.Debug("========================================================================\n")

	mux := http.NewServeMux()
	web.HandlerWebSite(mux)
	rest.HandlerRestServices(mux)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Error("ListenAndServe: ", err.Error())
	}

	log.Debug("\n========================================================================")
	log.Debug("Hyeragent stared.")
}
