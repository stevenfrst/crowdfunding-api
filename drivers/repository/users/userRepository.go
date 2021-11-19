package users

import (
	"context"
	"errors"
	"github.com/stevenfrst/crowdfunding-api/drivers/repository"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

func (r *UserRepository) CheckLogin(email,password string,ctx context.Context) (bool, error) {
	var user repoModels.User

	result := r.db.Where("email = ?", email,"password = ?",password).First(&user)
	if result.Error != nil {
		return false,result.Error
	}
	if result.RowsAffected == 1 {
		return true,nil
	}
	//log.Println(user,reflect.TypeOf(user))

	return false,errors.New("user not found")
}

func (r *UserRepository) Register(user *users.Domain,ctx context.Context) (users.Domain,error) {
	//log.Println("HIT")
	result := r.db.Create(repoModels.FromDomainUser(user))
	log.Println(reflect.TypeOf(user),result.RowsAffected)
	if result.Error != nil {
		return *user,result.Error
	}
	//var out repoModels.User
	return *user,nil
}

