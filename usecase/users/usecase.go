package users

import (
	"errors"
	"fmt"
	"time"
)
import _middleware "github.com/stevenfrst/crowdfunding-api/app/middleware"

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

// NewUsecase to create a new UserUseCaseInterface
func NewUsecase(userRepo UserRepoInterface,configJWT *_middleware.ConfigJWT, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}

// LoginUseCase method handle login
func (u UserUseCase) LoginUseCase(username,password string) (Domain,error) {
	user ,err := u.repo.CheckLogin(username,password)
	if err != nil {
		return user,errors.New("internal error")
	} else if user.ID == 0  {
		return Domain{},errors.New("email/password not match")
	}
	token,err := u.jwt.GenerateTokenJWT(int(user.ID))
	user.Token = token
	return user,err
}

// RegisterUseCase method to registering user
func (u UserUseCase) RegisterUseCase(user Domain) (Domain,error) {
	resp,err := u.repo.Register(&user)
	if err != nil {
		if fmt.Sprintf("%v",err) == "failed to create record" {
			return resp,errors.New("failed to registering user")
		} else {
			return resp,errors.New("internal error")
		}
	}
	return resp,nil
}

// GetAll method to get all user
func (u UserUseCase) GetAll() ([]Domain,error) {
	resp,err := u.repo.GetAllUser()
	if err != nil {
		return []Domain{},errors.New("Error Internal / Data Tidak Ditemukan")
	}
	return resp,nil
}
// DeleteByID method to handle repo delete user
func (u UserUseCase) DeleteByID(id int) (string,error) {
	resp,err := u.repo.DeleteUserByID(id)
	if resp == 0 || err != nil {
		return "Failed",errors.New("Gagal Menghapus Data/Internal Error")
	}
	return "Success",nil
}
// UpdatePassword method to handle update password from user > repo
func (u *UserUseCase) UpdatePassword(domain DomainUpdate) (string,error) {
	resp,err := u.repo.UpdateUserPassword(domain)
	return resp,err
}
// GetUserTransactionByID method to handle get transaction via id
func (u *UserUseCase) GetUserTransactionByID(id int) (DomainTransaction,error) {
	resp,err := u.repo.GetUserTransaction(id)
	if err != nil {
		return DomainTransaction{},err
	}
	return resp,nil
}
