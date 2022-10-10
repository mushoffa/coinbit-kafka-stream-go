package processor

import (
	"log"

	"wallet-service/data/datasource"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/pb"

	"github.com/lovoo/goka"
)

type balance struct {}

func NewBalanceProcessor() datasource.GokaProcessCallback {
	return &balance{}
}

func (p *balance) Handle(ctx goka.Context, message interface{}) {

	payload, ok := message.(*pb.DepositMoney)
	if !ok {
		return
	}

	var wallet *pb.Wallet
	if val := ctx.Value(); val != nil {
		wallet = val.(*pb.Wallet)
	} else {
		wallet = new(pb.Wallet)
	}

	wallet.WalletId = payload.GetWalletId()
	wallet.Balance += payload.GetAmount()

	log.Println("[Balance Processor]", wallet)

	ctx.SetValue(wallet)
}