package entity

type Wallet struct {
	WalletID string `json:"wallet_id"`
	Balance float64 `json:"balance"`
	Threshold bool `json:"above_threshold"`
}