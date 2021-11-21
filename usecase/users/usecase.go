package users

import (
	"context"
	"log"
	"reflect"
	"time"
)
import _middleware "github.com/stevenfrst/crowdfunding-api/app/middleware"

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUsecase(userRepo UserRepoInterface,configJWT *_middleware.ConfigJWT, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
		jwt:  configJWT,
	}
}


func (u UserUseCase) LoginUseCase(username,password string,ctx context.Context) (Domain,error) {

	user ,err := u.repo.CheckLogin(username,password,ctx)
	log.Println(reflect.TypeOf(user),user.Token)
	if err != nil {
		return user,err
	}
	token,err := u.jwt.GenerateTokenJWT(int(user.ID))
	user.Token = token
	return user,err
}


func (u UserUseCase) RegisterUseCase(user Domain,ctx context.Context) (Domain,error) {
	//log.Println(user)
	resp,err := u.repo.Register(&user,ctx)
	//log.Println("HITUSECASE")
	return resp,err
}


func (u UserUseCase) GetAll() ([]Domain,error) {
	resp,err := u.repo.GetAllUser()
	if err != nil {
		return []Domain{},err
	}
	return resp,nil
}

func (u UserUseCase) DeleteByID(id int) (string,error) {
	resp,err := u.repo.DeleteUserByID(id)
	if resp == 0 {
		return "Failed",err
	}
	return "Success",err
}
