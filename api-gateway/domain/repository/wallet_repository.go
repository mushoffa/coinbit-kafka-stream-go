package domain

import (
	"api-gateway/domain/entity"
)

type WalletRepository interface {
	GetByWalletID(string) (entity.Wallet, error)
}