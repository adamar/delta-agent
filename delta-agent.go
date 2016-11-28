package main

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/adamar/delta-server/models"
	"github.com/adamar/delta-agent/delta"
	"github.com/cskr/pubsub"
)

func main() {

	gob.Register(models.Response{})
	gob.Register(models.Event{})

	models.PubSub = pubsub.New(20)

	if delta.PassedPreflighChecks() != true {
		os.Exit(1)
	}

	rpc := delta.NewRPClient()
	go delta.StartAuditEngine()
	go delta.StartLogStreamEngine()
	go delta.StartProcFSEngine()

	inbound := models.PubSub.Sub("SystemCall", "Exec", "PathChange", "ConfigChange", "SystemEvent", "LogEvent", "ProcFS")

	for {

		select {
		case in := <-inbound:
			_, err := rpc.Call(in)
			if err != nil {
				log.Fatalf("Error when sending request to server: %s", err)
			}
		}
	}

}
