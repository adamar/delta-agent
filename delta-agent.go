package main

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/adamar/delta-agent/delta"
	"github.com/adamar/delta-server/models"
	"github.com/cskr/pubsub"
)

func main() {

	gob.RegisterName("Response", models.Response{})
	gob.RegisterName("Event", models.Event{})

	models.PubSub = pubsub.New(20)

	if delta.PassedPreflighChecks() != true {
		os.Exit(1)
	}

	rpc := delta.NewRPClient()
	go delta.StartAuditEngine()
	go delta.StartLogStreamEngine()
	go delta.StartProcFSEngine()
        go delta.StartiNotifyEngine()

	//inbound := models.PubSub.Sub("SystemCall", "Exec", "PathChange", "ConfigChange", "SystemEvent", "LogEvent", "ProcFS", delta.InotifyChannel)
	inbound := models.PubSub.Sub(delta.InotifyChannel)

	events := models.PubSub.Sub("SystemCall", "Exec", "PathChange", "ConfigChange", "SystemEvent")

	go delta.ParseEvents(events)

	for {

		select {
		case in := <-inbound:
			_, err := rpc.Call(in)
			log.Println(in)
			if err != nil {
				log.Fatalf("Error when sending request to server: %s", err)
			}
		}
	}

}
