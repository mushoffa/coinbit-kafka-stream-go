package kafka

import (
	"context"

	"wallet-service/data/datasource"

	"github.com/lovoo/goka"
)

type subscriber struct {
	processor *goka.Processor
}

func NewKafkaSubscriber(brokers []string, topic, groupID string, inputCodec, tableCodec datasource.GokaMessageCodec,callback datasource.GokaProcessCallback, tmc *goka.TopicManagerConfig) (datasource.KafkaSubscriber, error) {
	group := goka.DefineGroup(
		goka.Group(groupID),
		goka.Input(goka.Stream(topic), inputCodec, callback.Handle),
		goka.Persist(tableCodec),
	)

	processor, err := goka.NewProcessor(
		brokers,
		group,
		goka.WithTopicManagerBuilder(goka.TopicManagerBuilderWithTopicManagerConfig(tmc)),
		goka.WithConsumerGroupBuilder(goka.DefaultConsumerGroupBuilder),
	)

	if err != nil {
		return nil, err
	}

	return &subscriber{processor}, nil
}

func (s *subscriber) Run(ctx context.Context) func() error {
	return func() error {
		return s.processor.Run(ctx)
	}
}

func (s *subscriber) Stop() {
	s.processor.Stop()
}