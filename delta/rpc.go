package delta

import (
	"github.com/valyala/gorpc"
	"time"
)

func NewRPClient() *gorpc.Client {

	c := &gorpc.Client{
		Addr:           "127.0.0.1:12345",
		RequestTimeout: 90000 * time.Second,
	}
	c.Start()

	return c

}
