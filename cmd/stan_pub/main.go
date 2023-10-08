package main

import (
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	stanClusterID = "test-cluster"
	stanClientID  = "pub"
	channel       = "orders"
	testCases     = 4
)

func main() {
	sc, err := stan.Connect(stanClusterID, stanClientID)
	stan.SetConnectionLostHandler(func(_ stan.Conn, err error) {
		log.Fatalf("Connection lost, reason: %s", err.Error())
	})
	if err != nil {
		log.Fatalf("Cant connect to claster %s:%s", stanClusterID, err.Error())
	}

	defer sc.Close()

	for i := 1; i < testCases+1; i++ {
		path := "cmd/stan_pub/testcase" + strconv.Itoa(i) + ".json"
		testcase, err := os.Open(path)
		if err != nil {
			log.Fatalf(err.Error())
		}

		data, err := io.ReadAll(testcase)
		if err != nil {
			log.Fatalf(err.Error())
		}
		err = sc.Publish(channel, data)
		if err != nil {
			log.Printf("Failed to publish message:%s", err.Error())
		}
		log.Printf("Message %d send", i)
		time.Sleep(1 * time.Second)
	}
}
