package stan

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"wb_l0"
	"wb_l0/pkg/repository"
)

const (
	stanClusterId = "test-cluster"
)

func SubscribeToNS(repo *repository.Repository) {

	//db, err := repository.NewPostgresDB(repository.Config{
	//	Host:     "localhost",
	//	Port:     "5432",
	//	Username: "postgres",
	//	Password: "456123789",
	//	DBName:   "postgres",
	//	SSLMode:  "disable",
	//})
	//if err != nil {
	//	log.Fatalf("Cant initialize db: %s", err.Error())
	//}
	//
	//repo := repository.NewRepository(db)

	sc, err := stan.Connect(stanClusterId, "sub",
		stan.SetConnectionLostHandler(func(_ stan.Conn, err error) {
			log.Fatalf("Connection lost, reason: %s", err.Error())
		}))
	if err != nil {
		log.Fatalf("Connection lost, reason: %s", err.Error())
	}

	defer sc.Close()

	log.Printf("Connected to cluster %s", stanClusterId)

	startOpt := stan.DeliverAllAvailable()
	handler := func(msg *stan.Msg) {
		fmt.Printf("Got new message\n")
		parseMsg(repo, msg.Data)
	}

	sub, err := sc.QueueSubscribe("orders", "", handler, startOpt)
	if err != nil {
		log.Fatalf("Couldnt subscribe to cluster %s: %s", stanClusterId, err.Error())
	}

	for sub.IsValid() {

	}
	defer sub.Unsubscribe()
}

func parseMsg(repo *repository.Repository, msg []byte) {
	var order wb_l0.Order
	err := json.Unmarshal(msg, &order)
	if err != nil {
		fmt.Printf("Message is incorrect: %s", string(msg))
	} else {
		log.Println("Message is correct")
		id, err := repo.PushOrder(order)
		if err != nil {
			log.Printf("Couldnt push order to db: %s", err.Error())
		} else {
			log.Printf("Message deployed to db, id: %d", id)
		}
	}
}
