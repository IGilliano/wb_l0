package stan_sub

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"wb_l0"
	"wb_l0/pkg/repository"
)

const (
	ClusterID = "test-cluster"
	ClientID  = "test"
)

type StanSub struct {
	clusterID    string
	clientID     string
	connection   stan.Conn
	subscription stan.Subscription
}

func NewStanSub() *StanSub {
	return &StanSub{clusterID: ClusterID, clientID: ClientID}
}

func (s *StanSub) Run(repo *repository.Repository) {
	con, err := stan.Connect(s.clusterID, s.clientID)
	if err != nil {
		log.Fatalf("Cant connect to nats-streaming server: %s\n", err.Error())
	}
	s.connection = con
	log.Printf("Connected to cluster %s", ClusterID)

	if err = s.Subscribe(repo); err != nil {
		log.Fatalf("Couldnt subscribe to cluster %s: %s\n", ClusterID, err.Error())
	}
}

func (s *StanSub) Close() error {

	if err := s.connection.Close(); err != nil {
		log.Fatalf("Cant close the connection to nats-streaming server:%s\n", err.Error())
		return err
	}

	log.Printf("Disconnected from nats streaming server\n")

	return nil
}

func (s *StanSub) Subscribe(repo *repository.Repository) error {
	handler := func(msg *stan.Msg) {
		log.Printf("Got new message\n")
		if err := msg.Ack(); err != nil {
			log.Printf("Cant acknowledge a message:%v", err)
		}
		s.handleMsg(repo, msg.Data)
	}

	sub, err := s.connection.Subscribe("orders", handler, stan.StartWithLastReceived(), stan.SetManualAckMode())
	if err != nil {
		return err
	}

	s.subscription = sub
	return nil
}

func (s *StanSub) Unsubscribe() error {

	if err := s.subscription.Unsubscribe(); err != nil {
		return err
	}

	log.Printf("Unsubscribed from channel\n")

	return nil
}

func (s *StanSub) handleMsg(repo *repository.Repository, msg []byte) {
	var order wb_l0.Order
	err := json.Unmarshal(msg, &order)
	if err != nil {
		fmt.Printf("Message is incorrect: %v\n", err.Error())
	} else if order.OrderUid == "" {
		fmt.Printf("Message is incorrect: orderuid is empty\n")
	} else {
		id, err := repo.PushOrder(order)
		if err != nil {
			log.Printf("Couldnt push order to db: %s\n", err.Error())
		} else {
			log.Printf("Message deployed to db, id: %d\n", id)
		}
	}
}
