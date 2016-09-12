package main

import (
	"fmt"
	"hyperagent/heartbeat"
	"hyperagent/log"
	"hyperagent/monitor"
	"hyperagent/rest"
	"hyperagent/web"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
)

const (
	name        = "HyperAgent"
	description = "Hyper agent service"
)

//  dependencies that are NOT required by the service, but might be used
var dependencies = []string{"dummy.service"}

// Service has embedded daemon
type Service struct {
	daemon.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	// Do something, call your goroutines, etc

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	go doWork()

	// loop work cycle with accept connections or interrupt
	// by system signal
	for {
		select {
		case killSignal := <-interrupt:
			fmt.Println("Got signal:", killSignal)
			fmt.Println("Stoping daemon ")

			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}

	// never happen, but need to complete code
	return usage, nil
}

func doWork() {
	log.Debug("Start hyperagent ....")

	log.Debug("Load monitor ...")
	monitor.LoadMonitorConfig()
	heartbeat.StartHB()

	log.Debug("Start web and rest ....")
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

func main() {
	srv, err := daemon.New(name, description, dependencies...)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
