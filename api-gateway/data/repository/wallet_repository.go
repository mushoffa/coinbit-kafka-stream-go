package repository

import (
	"fmt"

	"api-gateway/data/datasource"
	"api-gateway/data/model"
	"api-gateway/domain/entity"
	"api-gateway/domain/repository"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/pb"
)

type wallet struct {
	walletTable datasource.KafkaViewTable
	thresholdTable datasource.KafkaViewTable
}

func NewWalletRepository(walletTable datasource.KafkaViewTable, thresholdTable datasource.KafkaViewTable) domain.WalletRepository {
	return &wallet{walletTable, thresholdTable}
}

func (r *wallet) GetByWalletID(id string) (entity.Wallet, error) {
	walletDetail := entity.Wallet{}

	walletData, err := r.walletTable.Get(id)
	if err != nil {
		return walletDetail, model.NewError(err.Error())
	}

	thresholdData, err := r.thresholdTable.Get(id)
	if err != nil {
		return walletDetail, model.NewError(err.Error())
	}

	if walletData == nil || thresholdData == nil {
		return walletDetail, model.NewError(fmt.Sprintf("Wallet with id %s is not found on the system", id))
	}

	wallet := walletData.(*pb.Wallet)
	threshold := thresholdData.(*pb.Threshold)

	walletDetail.WalletID = wallet.GetWalletId()
	walletDetail.Balance = wallet.GetBalance()
	walletDetail.Threshold = threshold.GetAboveThreshold()

	return walletDetail, nil
}