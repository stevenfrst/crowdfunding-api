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

type UserUsecaseInterface interface {
	LoginUseCase(username,password string,ctx context.Context) (bool,error)
	RegisterUseCase(user Domain,ctx context.Context) (Domain,error)
}

type UserRepoInterface interface {
	CheckLogin(email,password string,ctx context.Context) (bool, error)
	Register(user *Domain,ctx context.Context) (Domain,error)
}
