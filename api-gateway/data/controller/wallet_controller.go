package controller

import (
	"fmt"
	"net/http"
	"time"

	"api-gateway/data/model"
	"api-gateway/data/datasource"
	"api-gateway/domain/repository"

	"github.com/mushoffa/coinbit-kafka-stream-go-proto/pb"

	"github.com/gin-gonic/gin"
)

type WalletController interface {
	Router(*gin.Engine)
	Deposit(*gin.Context)
	Details(*gin.Context)
}

type wallet struct {
	publisher datasource.KafkaPublisher
	repository domain.WalletRepository
}

func NewWalletController(publisher datasource.KafkaPublisher, repository domain.WalletRepository) WalletController {
	return &wallet{publisher, repository}
}

func (c *wallet) Router(r *gin.Engine) {
	r.POST("/wallet/deposit", c.Deposit)
	r.GET("/wallet/details/:walletID", c.Details)
}

func (c *wallet) Deposit(ctx *gin.Context) {
	request := model.DepositRequest{}
	response := model.BaseResponse{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		response.Error = model.NewError("Invalid payload request")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if!(request.Amount > 0) {
		response.Error = model.NewError("The value of amount mus be greater than 0")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deposit := &pb.DepositMoney {
		WalletId: request.WalletID,
		Amount: request.Amount,
	}

	if err := c.publisher.Publish(request.WalletID, deposit); err != nil {
		response.Error = model.NewError(err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	now := time.Now().Local()
	timestamp := fmt.Sprintf("%d/%d/%d %02d:%02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second())

	response.Data = model.SuccessResponse {
		Timestamp: timestamp,
		Message: "Success",
	}

	ctx.JSON(http.StatusOK, response)
	return
}

func (c *wallet) Details(ctx *gin.Context) {
	response := model.BaseResponse{}

	walletID := ctx.Param("walletID")

	wallet, err := c.repository.GetByWalletID(walletID)
	if err != nil {
		response.Error = err
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = wallet
	ctx.JSON(http.StatusOK, response)
	return
}