package repoModels

import "github.com/stevenfrst/crowdfunding-api/usecase/transaction"

// FromDomainTransaction convert domain transaction to transaction repo
func FromDomainTransaction(domain *transaction.Domain) Transaction {
	return Transaction {
		ID:domain.ID,
		CampaignID:domain.CampaignID,
		UserID: domain.UserID,
		PaymentLink: domain.PaymentLink,
		Nominal: domain.Nominal,
		Status:domain.Status,
		TransactionStatus: domain.TransactionStatus,
		FraudStatus: domain.FraudStatus,
		PaymentType: domain.PaymentType,
	}
}

// ToDomain convert transaction repo to domain
func (t Transaction) ToDomain() transaction.Domain {
	return transaction.Domain{
		ID:t.ID,
		CampaignID:t.CampaignID,
		UserID: t.UserID,
		PaymentLink: t.PaymentLink,
		Nominal: t.Nominal,
		Status:t.Status,
		TransactionStatus: t.TransactionStatus,
		FraudStatus: t.FraudStatus,
		PaymentType: t.PaymentType,
	}
}


