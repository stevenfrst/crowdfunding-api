package users

import (
	//"github.com/stevenfrst/crowdfunding-api/drivers/repository/campaign"
	//"github.com/stevenfrst/crowdfunding-api/drivers/repository/transaction"
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID        uint
	FullName string
	Email string
	Password string
	Job    string
	RoleID uint
	Token string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DomainUpdate struct {
	Email string
	OldPassword string
	NewPassword string
}

type DomainTransaction struct {
	ID        uint
	FullName string
	Email string
	Password string
	Job    string
	RoleID uint
	Token string
	Transaction interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserUsecaseInterface interface {
	LoginUseCase(username,password string) (Domain,error)
	RegisterUseCase(user Domain) (Domain,error)
	GetAll() ([]Domain,error)
	DeleteByID(id int) (string,error)
	UpdatePassword(domain DomainUpdate) (string,error)
	GetUserTransactionByID(id int) (DomainTransaction,error)
}

type UserRepoInterface interface {
	CheckLogin(email,password string) (Domain, error)
	Register(user *Domain) (Domain,error)
	GetAllUser() ([]Domain,error)
	DeleteUserByID(id int) (int,error)
	UpdateUserPassword(update DomainUpdate) (string, error)
	GetUserTransaction(id int) (DomainTransaction,error)
	GetEmailByID(id int) (string,error)
}
