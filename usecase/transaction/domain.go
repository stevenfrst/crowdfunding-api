package transaction

type Domain struct {
	ID        uint
	CampaignID uint
	UserID uint
	PaymentLink string
	Nominal int
	Status string
	TransactionStatus string
	FraudStatus string
	PaymentType string
}

type DomainNotification struct {
	TransactionStatus string
	OrderID           string
	PaymentType       string
	FraudStatus       string
}

type TransactionUsecaseInterface interface {
	CreateTransaction(campaignID,userID,Nominal int) (Domain,error)
	GetNotificationPayment(input DomainNotification) (Domain,error)
	GetStatusByID(ID int) (Domain,error)
}

type TransactionRepoInterface interface {
	CreateTransaction(transaksi *Domain) (Domain,error)
	GetByID(ID int) (Domain,error)
	UpdateTransaction(transaction *Domain) (*Domain,error)
	GetLastTransactionID() (int, error)
}
