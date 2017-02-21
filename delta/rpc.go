package delta

import (
	"github.com/valyala/gorpc"
	"time"
	"os"
	"log"
)

func (dc *DeltaCore) NewRPClient() {

	server_url := os.Getenv("DELTA_SERVER")
        if server_url == "" {
                server_url = "127.0.0.1:12345"
        }	

	log.Println("Connecting to: ", server_url)

	dc.Rpc = &gorpc.Client{
		Addr:           server_url,
		RequestTimeout: 90000 * time.Second,
	}
	dc.Rpc.Start()

}
