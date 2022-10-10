package datasource

import (
	"context"
)

type KafkaSubscriber interface {
	// Run(context.Context) error
	Run(context.Context) func() error
	Stop()
}