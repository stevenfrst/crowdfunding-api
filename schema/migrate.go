package schema

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	FullName string
	Email string
	Password string
	Job    string
	RoleID uint
	Campaigns []Campaign
	Transaction []Transaction
	Token string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


type Campaign struct {
	ID        uint `gorm:"primarykey"`
	UserID uint
	CampaignName string
	ShortDescription string
	LongDescription string
	Target int
	AmountNow int
	Supporters int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


type Transaction struct {
	ID        uint `gorm:"primarykey"`
	CampaignID uint
	Campaign Campaign
	UserID uint
	User User
	PaymentLink string
	Nominal int
	Status string
	TransactionStatus string
	FraudStatus string
	PaymentType string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
