package config

import (
	"github.com/pkg/errors"
	"github.com/jaegertracing/jaeger/pkg/dashbase"
)

// Configuration describes the configuration properties needed to connect to an ElasticSearch cluster

type Configuration struct {
	Server     string
	KafkaHost  []string
	KafkaTopic string
}

// ClientBuilder creates new es.Client
type ClientBuilder interface {
	NewKafkaClient() (dashbase.KafkaClient, error)
	GetKafkaTopic() string
	GetService() string
}

// NewClient creates a new ElasticSearch client
func (c *Configuration) NewKafkaClient() (dashbase.KafkaClient, error) {
	kafkaClient := dashbase.KafkaClient{Hosts: c.KafkaHost}
	if len(c.KafkaHost) < 1 {
		return kafkaClient, errors.New("No servers specified")
	}
	err := kafkaClient.Open()
	if err != nil {
		return kafkaClient, err
	}
	return kafkaClient, nil
}

// GetNumShards returns number of shards from Configuration
func (c *Configuration) GetKafkaTopic() string {
	return c.KafkaTopic
}

func (c *Configuration) GetService() string {
	return c.Server
}
