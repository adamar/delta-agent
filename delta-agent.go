package main

import (
	//"encoding/gob"
	"log"

	"github.com/adamar/delta-agent/delta"
	"github.com/adamar/delta-server/models"
	"github.com/cskr/pubsub"
)

func main() {

	delta.Start()

	models.PubSub = pubsub.New(20)

	rpc := delta.NewRPClient()
	go delta.StartAuditEngine()
	go delta.StartLogStreamEngine()
	go delta.StartProcFSEngine()
        go delta.StartiNotifyEngine()

	inbound := models.PubSub.Sub(delta.InotifyChannel, delta.ProcfsChannel, delta.LogChannel)

	//events := models.PubSub.Sub("SystemCall", "Exec", "PathChange", "ConfigChange", "SystemEvent")

	//go delta.ParseEvents(events)

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
