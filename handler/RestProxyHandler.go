package handlers

import (
	model "nats-rest-proxy/model"
	repository "nats-rest-proxy/repository"
	"net/http"

	"github.com/labstack/echo"
)

type RestProxyHandler struct {
	natsClient *repository.NatsClient
}

func (rest *RestProxyHandler) GetHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (rest *RestProxyHandler) GetHome(c echo.Context) error {
	return c.JSON(http.StatusOK, "Rest Proxy API started")
}

func (rest *RestProxyHandler) PostNatsTopic(c echo.Context) error {
	var build model.Build

	topic := c.Param("topic")

	err := c.Bind(&build)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	rest.natsClient.Publish(topic, build)

	return c.JSON(http.StatusOK, build.Job)
}

func NewRestProxyHandler(e *echo.Echo, natsClient *repository.NatsClient) {
	handler := &RestProxyHandler{natsClient}

	e.GET("/healthcheck", handler.GetHealthCheck)
	e.GET("/", handler.GetHome)

	e.POST("/nats/:topic", handler.PostNatsTopic)

	return
}
