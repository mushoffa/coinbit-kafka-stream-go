package main

import (
	"log"
	"os"
	"time"

	"api-gateway/config"
	"api-gateway/data/controller"
	"api-gateway/data/repository"
	"api-gateway/infrastructure/kafka"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/codec"
	"github.com/mushoffa/go-library/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("Error loading time location: %v", err)
		}
	}

	httpHandler := gin.Default()

	publisher, err := kafka.NewKafkaPublisher(config.BrokersDocker, config.Topic, new(codec.DepositMoneyCodec))
	if err != nil {
		panic(err)
	}
	defer publisher.Close()

	walletViewTable, err := kafka.NewKafkaViewTable(config.BrokersDocker, config.BalanceGroup, new(codec.WalletCodec))
	if err != nil {
		panic(err)
	}

	thresholdViewTable, err := kafka.NewKafkaViewTable(config.BrokersDocker, config.ThresholdGroup, new(codec.ThresholdCodec))
	if err != nil {
		panic(err)
	}

	walletViewTable.Run()
	thresholdViewTable.Run()

	repository := repository.NewWalletRepository(walletViewTable, thresholdViewTable)
	controller := controller.NewWalletController(publisher, repository)
	controller.Router(httpHandler)

	server := http.NewHttpServer(9091, httpHandler)
	server.Run()
}