package main

import (
	"nats-rest-proxy/config"

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
	natsSettings := configuration.NatsSettings{config.GetString("nats.connection"), config.GetString("nats.cluster")}

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	// initialize handlers
	natsProducerRepository := repository.NewNatsClient(natsSettings, config.GetString("nats.client"), config.GetString("nats.consumerGroup"))
	handler.NewRestProxyHandler(e, &natsProducerRepository)

	elastic := handler.NewElasticHandler(config.GetString("elastic.url"), config.GetString("elastic.index"))
	natsElasticRepository := repository.NewNatsClient(natsSettings, "elastic-consumer", "elastic-consumer-1")
	natsElasticRepository.Subscribe("jenkins", elastic.PushToIndex)

	// Start http server
	e.Logger.Fatal(e.Start(config.GetString("server.port")))
}
