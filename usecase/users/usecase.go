package users

import (
	"context"
	"time"
)
import _middleware "github.com/stevenfrst/crowdfunding-api/app/middleware"

type UserUseCase struct {
	repo UserRepoInterface
	ctx  time.Duration
	jwt  *_middleware.ConfigJWT
}

func NewUsecase(userRepo UserRepoInterface, contextTimeout time.Duration) UserUsecaseInterface {
	return &UserUseCase{
		repo: userRepo,
		ctx:  contextTimeout,
		//jwt:  configJWT,
	}
}


func (u UserUseCase) LoginUseCase(username,password string,ctx context.Context) (bool,error) {
	return u.repo.CheckLogin(username,password,ctx)
}


func (u UserUseCase) RegisterUseCase(user Domain,ctx context.Context) (Domain,error) {
	//log.Println(user)
	resp,err := u.repo.Register(&user,ctx)
	//log.Println("HITUSECASE")
	return resp,err
}

