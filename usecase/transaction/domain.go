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

type Email struct {
	Sender string
	ToEmail string
	Subject string
	Reward string
	Nominal string
}

type BodyEmail struct {
	TransactionStatus string
	OrderID string
	Reward string
}

type TransactionUsecaseInterface interface {
	CreateTransaction(campaignID,userID,Nominal int) (Domain,error)
	GetNotificationPayment(input DomainNotification) (Domain,error)
	GetStatusByID(ID int) (Domain,error)
	GetRewardByAmount(amount int) (int,string,error)
}

type TransactionRepoInterface interface {
	CreateTransaction(transaksi *Domain) (Domain,error)
	GetByID(ID int) (Domain,error)
	UpdateTransaction(transaction *Domain) (*Domain,error)
	GetLastTransactionID() (int, error)
}
