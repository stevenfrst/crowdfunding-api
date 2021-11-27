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

// InitializeSnapClient use for initializing servery key and midtrans mode
func InitializeSnapClient() {
	s.New(midtrans.ServerKey, midtrans.Sandbox)
}

// SetupGlobalMidtransConfig use for setup midtrans config to global
func (config ConfigMidtrans) SetupGlobalMidtransConfig() {
	midtrans.ServerKey = config.SERVER_KEY
	midtrans.Environment = midtrans.Sandbox
}

// GenerateSnapReq used for generate snap request to server
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

// GetLinkResponse request midtrans response like link etc.
func (config ConfigMidtrans) GetLinkResponse(id,nominal int) *snap.Response {
	resp, err := s.CreateTransaction(config.GenerateSnapReq(id,nominal))
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	//fmt.Println("Response : ", resp)
	return resp
}
