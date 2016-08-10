package main

import (
	"hyperagent/host"
	"hyperagent/log"
)

func main() {
	log.Debug("Start hyperagent ....")
	log.Debug("========================================================================\n")

	log.Debug(host.GetHostInfo().String())

	//	log.Debug("Output debug ...")
	//	log.Info("Output info ...")
	//	log.Warn("Output warn ...")
	//	log.Error("Output error ...")

	log.Debug("\n========================================================================")
	log.Debug("Hyeragent stared.")
}
