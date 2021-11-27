package users

import (
	"errors"
	"github.com/stevenfrst/crowdfunding-api/drivers/repository"
	"github.com/stevenfrst/crowdfunding-api/helper/encrypt"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository function to create a new UserRepository
func NewUserRepository(gormDb *gorm.DB) users.UserRepoInterface {
	return &UserRepository{
		db: gormDb,
	}
}

// CheckLogin methods to check if user ok or not
func (r *UserRepository) CheckLogin(email,password string) (users.Domain, error) {
	var user repoModels.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return users.Domain{},errors.New("error db not working")
	}
	err = encrypt.CheckPassword(password,user.Password)
	if err != nil {
		return users.Domain{},nil
	}

	return user.ToDomain(),nil
}

// Register methods to registering a user
func (r *UserRepository) Register(user *users.Domain) (users.Domain,error) {

	userIn := repoModels.FromDomainUser(user)
	hashedPassword,err := encrypt.Hash(user.Password)
	if err != nil {
		return users.Domain{},err
	}

	userIn.Password = hashedPassword
	result := r.db.Create(userIn)
	//log.Println(reflect.TypeOf(user),result.RowsAffected)
	if result.Error != nil {
		return userIn.ToDomain(),errors.New("failed to create record")
	}

	return userIn.ToDomain(),nil
}

// GetAllUser methods to get all user
func (r UserRepository) GetAllUser() ([]users.Domain,error) {
	var users []repoModels.User
	err := r.db.Find(&users).Error
	if err != nil {
		return repoModels.ConvertRepoUseCaseUserList(users),err
	}
	return repoModels.ConvertRepoUseCaseUserList(users),nil
}

// DeleteUserByID methods to delete a user by id
func (r UserRepository) DeleteUserByID(id int) (int,error) {
	var user repoModels.User
	result := r.db.Where("id = ? AND role_id = ?",id,2).Delete(&user)
	if result.Error != nil {
		return int(result.RowsAffected),result.Error
	}
	return int(result.RowsAffected),nil
}

// UpdateUserPassword methods to update a user
func (r UserRepository) UpdateUserPassword(update users.DomainUpdate) (string, error) {
	user,err := r.CheckLogin(update.Email,update.OldPassword)
	if err != nil {
		return "failed",err
	}

	hashedPassword,err := encrypt.Hash(update.NewPassword)
	log.Println("errdisni1")
	if err != nil {
		return "failed",err
	}
	user.Password = hashedPassword
	log.Println("errdisni2")

	err = r.db.Save(repoModels.FromDomainUser(&user)).Error
	if err != nil {
		return "failed",err
	}
	return "Success",err
}

// GetUserTransaction methods to get a user transaction
func(r UserRepository) GetUserTransaction(id int) (users.DomainTransaction,error) {
	var user repoModels.User
	err := r.db.Preload("Transaction").Where("id = ?",id).Find(&user).Error
	if err != nil {
		return users.DomainTransaction{},nil
	}
	return user.ToDomainUserTransaction(),nil
}

// GetEmailByID methods to get email by id
func (r UserRepository) GetEmailByID(id int) (string,error) {
	var user repoModels.User
	err := r.db.Where("id = ?",id).Find(&user).Error
	if err != nil {
		return "ID not found",err
	}
	return user.Email,nil
}