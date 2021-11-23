package repoModels

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	FullName string
	Email string `gorm:"unique"`
	Password string
	Job    string
	RoleID uint
	Campaigns []Campaign
	Transaction []Transaction
	RewardHistory []RewardHistory
	Token string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}



type Campaign struct {
	ID        uint `gorm:"primarykey" json:"id"`
	UserID uint `json:"user_id"`
	CampaignName string `json:"campaign_name"`
	ShortDescription string `json:"short_description"`
	LongDescription string `json:"long_description"`
	Target int `json:"target"`
	AmountNow int `json:"amount_now"`
	Supporters int `json:"supporters"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
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
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type TransactionUser struct {
	ID        uint `gorm:"primarykey"`
	CampaignID uint
	PaymentLink string
	Nominal int
	Status string
	TransactionStatus string
	FraudStatus string
	PaymentType string
}


type Reward struct {
	ID        uint `gorm:"primarykey"`
	Amount int
	RewardDescription string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RewardHistory struct {
	ID        uint `gorm:"primarykey"`
	UserID uint
	RewardID uint
	Reward Reward
	TransactionID uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}