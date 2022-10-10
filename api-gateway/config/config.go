package config

var (
	Brokers = []string{"localhost:19092"}
	BrokersDocker = []string{"broker:9092"}
	Topic = "deposits"
	BalanceGroup = "balance-group"
	ThresholdGroup = "threshold-group"
)