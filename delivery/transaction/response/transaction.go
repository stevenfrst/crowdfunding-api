package response

import "github.com/stevenfrst/crowdfunding-api/usecase/transaction"

type TransactionResponse struct {
	ID        uint `json:"id"`
	CampaignID uint `json:"campaign_id"`
	UserID uint `json:"user_id"`
	PaymentLink string `json:"payment_link"`
	Nominal int `json:"nominal"`
	Status string
	TransactionStatus string `json:"transaction_status"`
	FraudStatus string `json:"fraud_status"`
	PaymentType string `json:"payment_type"`
}

func FromDomain(domain transaction.Domain) TransactionResponse {
	return TransactionResponse{
		ID:domain.ID,
		CampaignID:domain.CampaignID,
		UserID:domain.UserID,
		PaymentLink:domain.PaymentLink,
		Nominal:domain.Nominal,
		Status:domain.Status,
		TransactionStatus:domain.TransactionStatus,
		FraudStatus:domain.FraudStatus,
		PaymentType:domain.PaymentType,
	}
}