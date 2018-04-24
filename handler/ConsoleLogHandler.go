package handlers

import (
	"encoding/json"
	"log"
	model "nats-rest-proxy/model"
)

type ConsoleLogHandler struct {
}

func (h *ConsoleLogHandler) WriteConsoleLog(build model.Build) {
	str, _ := json.Marshal(build)
	log.Printf("Message Received: %s\n", str)
}

func NewConsoleLogHandler() ConsoleLogHandler {
	v := ConsoleLogHandler{}
	return v
}
