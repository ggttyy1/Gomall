package mq

import "github.com/nats-io/nats.go"

var (
	Nc  *nats.Conn
	Err error
)

func InitNats() {
	// Connect to NATS
	Nc, Err = nats.Connect(nats.DefaultURL)
	if Err != nil {
		panic(Err)
	}

}
