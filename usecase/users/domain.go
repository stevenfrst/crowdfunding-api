package users

import (
	"context"
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

type UserUsecaseInterface interface {
	LoginUseCase(username,password string,ctx context.Context) (Domain,error)
	RegisterUseCase(user Domain,ctx context.Context) (Domain,error)
	GetAll() ([]Domain,error)
	DeleteByID(id int) (string,error)
	UpdatePassword(domain DomainUpdate,ctx context.Context) (string,error)
}

type UserRepoInterface interface {
	CheckLogin(email,password string,ctx context.Context) (Domain, error)
	Register(user *Domain,ctx context.Context) (Domain,error)
	GetAllUser() ([]Domain,error)
	DeleteUserByID(id int) (int,error)
	UpdateUserPassword(update DomainUpdate,ctx context.Context) (string, error)
}
