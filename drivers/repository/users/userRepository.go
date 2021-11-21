package users

import (
	"context"
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

func (r *UserRepository) CheckLogin(email,password string,ctx context.Context) (users.Domain, error) {
	var user repoModels.User

	err := r.db.Where("email = ?", email,"password = ?",password).First(&user).Error
	log.Println(err)
	if err != nil {
		return users.Domain{},err
	}
	return user.ToDomain(),nil
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

func (r UserRepository) GetAllUser() ([]users.Domain,error) {
	var users []repoModels.User
	err := r.db.Find(&users).Error
	if err != nil {
		return repoModels.ConvertRepoUseCaseUserList(users),err
	}
	return repoModels.ConvertRepoUseCaseUserList(users),nil
}

