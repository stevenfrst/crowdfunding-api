package repoModels

import (
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
)

func ConvertRepoUseCaseUserList(repo []User) (domain []users.Domain) {
	for _,user := range repo {
		newdomain := users.Domain{
			ID:user.ID,
			FullName:user.FullName,
			Email:user.Email,
			Password:user.Password,
			Job:user.Job,
			RoleID:user.RoleID,
		}
		domain = append(domain,newdomain)
	}
	return domain
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

func (u *User) ToDomain() users.Domain {
	return users.Domain{
		ID:      u.ID,
		FullName: u.FullName,
		Job:      u.Job,
		Email:    u.Email,
		Password: u.Password,
		RoleID: u.RoleID,
		Token:u.Token,
	}
}

func (u *User) ToDomainUserTransaction() users.DomainTransaction {
	var transaction []TransactionUser
	for _,domain := range u.Transaction {
		newTransaction := TransactionUser {
			ID: domain.ID,
			CampaignID: domain.CampaignID,
			PaymentLink:domain.PaymentLink,
			Nominal: domain.Nominal,
			Status: domain.Status,
			TransactionStatus: domain.TransactionStatus,
			FraudStatus: domain.FraudStatus,
			PaymentType: domain.PaymentType,
		}
		transaction = append(transaction,newTransaction)

	}
	return users.DomainTransaction{
		ID:      u.ID,
		FullName: u.FullName,
		Job:      u.Job,
		Email:    u.Email,
		Password: u.Password,
		RoleID: u.RoleID,
		Transaction: transaction,
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

