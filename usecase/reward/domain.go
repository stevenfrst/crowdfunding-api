package reward

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID        uint
	Amount int
	RewardDescription string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type DomainHistory struct {
	ID        uint
	UserID uint
	RewardID uint
	TransactionID uint
}

type RewardUsecaseInterface interface {
	CreateReward(domain Domain) (Domain,error)
	UpdateReward(domain Domain) (Domain,error)
	DeleteRewardByID(id int) (string,error)
}

type RewardRepoInterface interface {
	GetRewardByAmount(amount int) (int,string,error)
	SaveRewardHistory(domain DomainHistory) error
	CreateReward(domain Domain) (Domain,error)
	UpdateReward(domain Domain) (Domain,error)
	DeleteRewardByID(id int) (string,error)
}


