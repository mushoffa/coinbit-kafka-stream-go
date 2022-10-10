package model

type WalletDetailResponse struct {
	WalletID string `json:"wallet_id"`
	Balance float64 `json:"balance"`
	Threshold bool `json:"aboved_threshold"`
}