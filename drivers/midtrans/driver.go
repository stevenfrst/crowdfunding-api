package payment

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransInterface interface {
	SetupGlobalMidtransConfig()
	GenerateSnapReq(id,nominal int) *snap.Request
	GetLinkResponse(id,nominal int) *snap.Response
}

type ConfigMidtrans struct {
	SERVER_KEY string
}

var s snap.Client

func InitializeSnapClient() {
	s.New(midtrans.ServerKey, midtrans.Sandbox)
}

func (config ConfigMidtrans) SetupGlobalMidtransConfig() {
	midtrans.ServerKey = config.SERVER_KEY
	midtrans.Environment = midtrans.Sandbox
}


func (config ConfigMidtrans) GenerateSnapReq(id,nominal int) *snap.Request {
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("%v",id),
			GrossAmt: int64(nominal),
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}
	return snapReq
}


func (config ConfigMidtrans) GetLinkResponse(id,nominal int) *snap.Response {
	resp, err := s.CreateTransaction(config.GenerateSnapReq(id,nominal))
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	//fmt.Println("Response : ", resp)
	return resp
}
