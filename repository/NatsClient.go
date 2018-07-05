package repository

import (
	"encoding/json"
	"fmt"
	"log"
	config "nats-rest-proxy/config"
	model "nats-rest-proxy/model"

	"github.com/nats-io/go-nats-streaming"
	"github.com/nats-io/go-nats-streaming/pb"
)

type NatsClient struct {
	connection    stan.Conn
	consumerGroup string
}

func (cl *NatsClient) Publish(topic string, build model.Build) {

	ackHandler := func(ackedNuid string, err error) {
		if err != nil {
			log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
		}
	}

	payload, _ := json.Marshal(build)
	cl.connection.PublishAsync(topic, []byte(payload), ackHandler)
}

func (cl *NatsClient) Subscribe(topic string, handler model.BuildHandler) {

	mcb := func(msg *stan.Msg) {
		build := model.Build{}
		json.Unmarshal(msg.Data, &build)

		if handler != nil {
			handler(build)
		}

		msg.Ack()
	}

	startOpt := stan.StartAt(pb.StartPosition_LastReceived)

	_, err := cl.connection.QueueSubscribe(topic, cl.consumerGroup, mcb, startOpt, stan.DurableName(cl.consumerGroup))

	if err != nil {
		cl.connection.Close()
		log.Fatal(err)
	}
}

func NewNatsClient(settings config.NatsSettings, clientId string, group string) NatsClient {
	sc, err := stan.Connect(settings.ClusterId, clientId, stan.NatsURL(settings.Url))

	if err != nil {
		fmt.Println(err)
	}
	client := NatsClient{sc, group}
	return client
}
