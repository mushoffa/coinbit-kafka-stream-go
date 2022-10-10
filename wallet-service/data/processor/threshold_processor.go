package processor

import(
	"log"
	"time"

	"wallet-service/data/datasource"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/pb"

	"github.com/lovoo/goka"
)

type threshold struct {}

func NewThresholdProcessor() datasource.GokaProcessCallback {
	return &threshold{}
}

func (p *threshold) Handle(ctx goka.Context, message interface{}) {
	payload, ok := message.(*pb.DepositMoney)
	if !ok {
		return
	}

	var threshold *pb.Threshold

	currentTimestamp := time.Now().UnixNano()

	if val := ctx.Value(); val != nil {
		threshold = val.(*pb.Threshold)
	} else {
		threshold = new(pb.Threshold)
		threshold.LastDepositTimestamp = currentTimestamp
	}

	threshold.WalletId = payload.GetWalletId()
	currentDepositTimestamp := time.Unix(0, currentTimestamp)
	lastDepositTimestamp := time.Unix(0, threshold.LastDepositTimestamp)
	rollingPeriod := currentDepositTimestamp.Sub(lastDepositTimestamp)

	if rollingPeriod.Seconds() > float64(120) {
		threshold.AboveThreshold = false
		threshold.LastDepositTimestamp = currentTimestamp
		threshold.TotalDeposit = payload.GetAmount()
	} else {
		threshold.TotalDeposit += payload.GetAmount()
		if threshold.TotalDeposit > float64(10000) {
			threshold.AboveThreshold = true
		} else {
			threshold.AboveThreshold = false
		}
	}

	log.Println("[Threshold Processor]", threshold)

	ctx.SetValue(threshold)
}