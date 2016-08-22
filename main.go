package main

import (
	"hyperagent/gui"
	"hyperagent/log"
	"hyperagent/rest"
	"hyperagent/web"
	"net/http"
	// "os"

	// "golang.org/x/sys/windows/svc"
)

func main() {
	//	const svcName = "myservice"

	//	isIntSess, err := svc.IsAnInteractiveSession()
	//	if err != nil {
	//		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	//	}
	//	if !isIntSess {
	//		runService(svcName, false)
	//		return
	//	}

	//	if len(os.Args) < 2 {
	//		usage("no command specified")
	//	}

	//	cmd := strings.ToLower(os.Args[1])
	//	switch cmd {
	//	case "debug":
	//		runService(svcName, true)
	//		return
	//	case "install":
	//		err = installService(svcName, "my service")
	//	case "remove":
	//		err = removeService(svcName)
	//	case "start":
	//		err = startService(svcName)
	//	case "stop":
	//		err = controlService(svcName, svc.Stop, svc.Stopped)
	//	case "pause":
	//		err = controlService(svcName, svc.Pause, svc.Paused)
	//	case "continue":
	//		err = controlService(svcName, svc.Continue, svc.Running)
	//	default:
	//		usage(fmt.Sprintf("invalid command %s", cmd))
	//	}
	//	if err != nil {
	//		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	//	}
	//	return
	//}

	//func runService() {
	log.Debug("Start hyperagent ....")
	log.Debug("=====================================================getHostName===================\n")

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
