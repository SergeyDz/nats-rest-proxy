package handlers

import (
	"context"
	"log"
	model "nats-rest-proxy/model"

	"github.com/olivere/elastic"
)

type ElasticHandler struct {
	es    elastic.Client
	index string
}

func (h *ElasticHandler) PushToIndex(build model.Build) {

	_, err := h.es.Index().
		Index(h.index).
		Type("doc").
		BodyJson(build).
		Do(context.Background())

	if err != nil {
		log.Printf("Warning: error pushng to elastic build id %s: %v\n", build.JobID, err.Error())
	}
}

func NewElasticHandler(url string, index string) ElasticHandler {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		panic(err)
	}

	v := ElasticHandler{*client, index}
	return v
}
