package request

import "github.com/stevenfrst/crowdfunding-api/usecase/transaction"

type TransactionNotification struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

func (notification TransactionNotification) ToDomainNotification() transaction.DomainNotification {
	return transaction.DomainNotification{
		TransactionStatus:notification.TransactionStatus,
		OrderID:notification.OrderID,
		PaymentType:notification.PaymentType,
		FraudStatus:notification.FraudStatus,
	}
}