package main

import (
	//"encoding/gob"
	"log"

	"github.com/adamar/delta-agent/delta"
	"github.com/adamar/delta-server/models"
	"github.com/cskr/pubsub"
)

func main() {

	DC, _ := delta.Start()
	
	models.PubSub = pubsub.New(20)

	inbound := models.PubSub.Sub(delta.InotifyChannel, delta.ProcfsChannel, delta.LogChannel)

	//events := models.PubSub.Sub("SystemCall", "Exec", "PathChange", "ConfigChange", "SystemEvent")

	//go delta.ParseEvents(events)

	for {

		select {
		case in := <-inbound:
			_, err := DC.Rpc.Call(in)
			log.Println("RPC Error: " , in)
			if err != nil {
				log.Fatalf("Error when sending request to server: %s", err)
			}
		}
	}

}
