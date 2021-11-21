package users

import (
	"context"
	"github.com/stevenfrst/crowdfunding-api/drivers/repository"
	"github.com/stevenfrst/crowdfunding-api/helper/encrypt"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
	"gorm.io/gorm"
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

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return users.Domain{},err
	}

	err = encrypt.CheckPassword(password,user.Password)
	if err != nil {
		return users.Domain{},err
	}

	return user.ToDomain(),nil
}

func (r *UserRepository) Register(user *users.Domain,ctx context.Context) (users.Domain,error) {

	userIn := repoModels.FromDomainUser(user)
	hashedPassword,err := encrypt.Hash(user.Password)
	if err != nil {
		return users.Domain{},err
	}

	userIn.Password = hashedPassword
	result := r.db.Create(userIn)
	//log.Println(reflect.TypeOf(user),result.RowsAffected)
	if result.Error != nil {
		return userIn.ToDomain(),result.Error
	}

	return userIn.ToDomain(),nil
}

func (r UserRepository) GetAllUser() ([]users.Domain,error) {
	var users []repoModels.User
	err := r.db.Find(&users).Error
	if err != nil {
		return repoModels.ConvertRepoUseCaseUserList(users),err
	}
	return repoModels.ConvertRepoUseCaseUserList(users),nil
}

