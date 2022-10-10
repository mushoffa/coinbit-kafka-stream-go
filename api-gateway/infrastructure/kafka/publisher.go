package kafka

import (
	"api-gateway/data/datasource"

	"github.com/lovoo/goka"
)

type publisher struct {
	*goka.Emitter
}

func NewKafkaPublisher(brokers []string, topic string, codec datasource.GokaMessageCodec) (datasource.KafkaPublisher, error) {
	emitter, err := goka.NewEmitter(brokers, goka.Stream(topic), codec)
	if err != nil {
		return nil, err
	}

	return &publisher{emitter}, nil
}

func (p *publisher) Run() {
	
}

func (p *publisher) Publish(key string, data interface{}) error{
	if err := p.EmitSync(key, data); err != nil {
		return err
	}

	return nil
}

func (p *publisher) Close() {
	p.Emitter.Finish()
}