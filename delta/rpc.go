package delta

import (
	"github.com/valyala/gorpc"
	"time"
	"os"
)

func NewRPClient() *gorpc.Client {

	server_url := os.Getenv("SERVER")
        if server_url == "" {
                server_url = "127.0.0.1:12345"
        }	

	c := &gorpc.Client{
		Addr:           server_url,
		RequestTimeout: 90000 * time.Second,
	}
	c.Start()

	return c

}
