package model

type DepositRequest struct {
	WalletID string `json:"wallet_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}