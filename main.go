package main

import (
	"github.com/labstack/echo"

	config "nats-rest-proxy/config"
	handler "nats-rest-proxy/handler"
	middleware "nats-rest-proxy/middleware"
	repository "nats-rest-proxy/repository"
)

func main() {
	// Echo instance
	e := echo.New()
	config := config.NewViperConfig()

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	// initialize handlers
	natsProducerRepository := repository.NewNatsClient(config.GetString("nats.connection"), config.GetString("nats.cluster"), config.GetString("nats.client"), config.GetString("nats.consumerGroup"))
	handler.NewRestProxyHandler(e, &natsProducerRepository)

	console := handler.NewConsoleLogHandler()
	natsProducerRepository.Subscribe("jenkins", console.WriteConsoleLog)

	// Start http server
	e.Logger.Fatal(e.Start(config.GetString("server.port")))
}
