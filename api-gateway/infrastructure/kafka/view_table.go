package kafka

import (
	"context"
	"log"

	"api-gateway/data/datasource"

	"github.com/lovoo/goka"
)

type viewtable struct {
	view *goka.View
}

func NewKafkaViewTable(brokers []string, tableGroup string, codec datasource.GokaMessageCodec) (datasource.KafkaViewTable, error) {
	view, err := goka.NewView(
		brokers,
		goka.GroupTable(goka.Group(tableGroup)),
		codec,
	)
	if err != nil {
		return nil, err
	}


	return &viewtable{view}, nil
}

func (v *viewtable) Run() {
	go func(ctx context.Context) {
		if err := v.view.Run(ctx); err != nil {
			log.Println("Error running view table: ", err)
		}
	}(context.Background())
}

func (v *viewtable) Get(key string) (interface{}, error) {
	return v.view.Get(key)
}

func (v *viewtable) GetInstance() *goka.View {
	return v.view
}