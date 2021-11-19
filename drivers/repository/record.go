package repoModels

import (
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
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

func FromDomainUser(domain *users.Domain) *User {
	return &User {
		ID:      domain.ID,
		FullName: domain.FullName,
		Job:      domain.Job,
		Email:    domain.Email,
		Password: domain.Password,
		RoleID: domain.RoleID,
	}
}

func (u User) ToDomain() users.Domain {
	return users.Domain{
		ID:      u.ID,
		FullName: u.FullName,
		Job:      u.Job,
		Email:    u.Email,
		Password: u.Password,
		RoleID: u.RoleID,
	}
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
