package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"wallet-service/config"
	"wallet-service/data/processor"
	"wallet-service/infrastructure/kafka"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/codec"

	"github.com/lovoo/goka"
	"github.com/Shopify/sarama"
	"golang.org/x/sync/errgroup"
)

func main() {
	tmc := goka.NewTopicManagerConfig()
	tmc.Table.Replication = 1
	tmc.Stream.Replication = 1

	gokaConfig := goka.DefaultConfig()
	gokaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	goka.ReplaceGlobalConfig(gokaConfig)

	tm, err := goka.NewTopicManager(config.BrokersDocker, goka.DefaultConfig(), tmc)
	if err != nil {
		log.Fatalf("Error creating topic manager: %v", err)
	}
	defer tm.Close()

	err = tm.EnsureStreamExists(config.Topic, 1)

	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	balanceProcessor := processor.NewBalanceProcessor()
	thresholdProcessor := processor.NewThresholdProcessor()

	depositMoneyCodec := new(codec.DepositMoneyCodec)
	thresholdCodec := new(codec.ThresholdCodec)
	walletCodec := new(codec.WalletCodec)

	balanceSubscriber, err := kafka.NewKafkaSubscriber(config.BrokersDocker, config.Topic, config.BalanceGroup, depositMoneyCodec, walletCodec, balanceProcessor, tmc)
	if err != nil {
		panic(err)
	}

	thresholdSubscriber, err := kafka.NewKafkaSubscriber(config.BrokersDocker, config.Topic, config.ThresholdGroup, depositMoneyCodec, thresholdCodec, thresholdProcessor, tmc)
	if err != nil {
		panic(err)
	}

	grp.Go(balanceSubscriber.Run(ctx))
	grp.Go(thresholdSubscriber.Run(ctx))

	// Wait for SIGINT/SIGTERM
	waiter := make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-waiter:
	case <-ctx.Done():
	}

	cancel()
	balanceSubscriber.Stop()
	thresholdSubscriber.Stop()

	if err := grp.Wait(); err != nil {
		log.Println(err)
	}
}