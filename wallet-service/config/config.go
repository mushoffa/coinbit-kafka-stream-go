package config

var (
	// Brokers = []string{"localhost:9092", "localhost:9093", "localhost:9094"}
	Brokers = []string{"localhost:19092"}
	BrokersDocker = []string{"broker:9092"}
	Topic = "deposits"
	BalanceGroup = "balance-group"
	ThresholdGroup = "threshold-group"
)