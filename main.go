package main

import (
	"engine-lib/mq"
	"time"
)

func main() {
	m := mq.MQ{}
	m.Init("localhost", 5672)
	m.Create("masaruz")
	go m.Receive()
	for {
		go m.Send()
		<-time.After(1 * time.Second)
	}
}
