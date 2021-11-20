package repoModels

import (
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
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

func ConvertRepoUseCaseUserCampaign(repo []User) (domain []campaign.Users) {
	for _,x := range repo {
		newDomain := campaign.Users{
			ID:x.ID,
			FullName:x.FullName,
			Email:x.Email,
			Password: x.Password,
			Job: x.Job,
			RoleID: x.RoleID,
			Campaigns: x.Campaigns,
			//Token string
		}
		domain = append(domain, newDomain)
	}
	return domain
}


func ConvertRepoUserCampaign(repo User) (domain campaign.UserCampaign) {
	return campaign.UserCampaign{
		ID:repo.ID,
		FullName:repo.FullName,
		Email:repo.Email,
		Password: repo.Password,
		Job: repo.Job,
		RoleID: repo.RoleID,
		Campaigns: repo.Campaigns,
		//Token string
	}
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

func (u User) ToDomainList() users.Domain {
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
	ID        uint `gorm:"primarykey" json:"id"`
	UserID uint `json:"user_id"`
	CampaignName string `json:"campaign_name"`
	ShortDescription string `json:"short_description"`
	LongDescription string `json:"long_description"`
	Target int `json:"target"`
	AmountNow int `json:"amount_now"`
	Supporters int `json:"supporters"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FromDomainCampaign(domain *campaign.Domain) *Campaign {
	return &Campaign {
		ID:domain.ID,
		UserID:domain.UserID,
		CampaignName:domain.CampaignName,
		ShortDescription: domain.ShortDescription,
		LongDescription: domain.LongDescription,
		Target: domain.Target,
		AmountNow: domain.AmountNow,
		Supporters:domain.Supporters,
	}
}

func (c Campaign) ToDomain() campaign.Domain {
	return campaign.Domain{
		ID:c.ID,
		UserID:c.UserID,
		CampaignName:c.CampaignName,
		ShortDescription: c.ShortDescription,
		LongDescription: c.LongDescription,
		Target: c.Target,
		AmountNow: c.AmountNow,
		Supporters:c.Supporters,
	}
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
