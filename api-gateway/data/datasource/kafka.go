package datasource

import (
	// "context"

	"github.com/lovoo/goka"
)

type KafkaPublisher interface {
	Run()
	Publish(string, interface{}) error
	Close()
}

type KafkaViewTable interface {
	Run()
	Get(string) (interface{}, error)
	GetInstance() *goka.View
}