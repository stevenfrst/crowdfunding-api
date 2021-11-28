package users_test

import (
	"errors"
	"fmt"
	config2 "github.com/stevenfrst/crowdfunding-api/app/config"
	middlewares "github.com/stevenfrst/crowdfunding-api/app/middleware"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
	_mockUserRepository "github.com/stevenfrst/crowdfunding-api/usecase/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

var userRepositoryMock _mockUserRepository.UserRepoInterface
var userUsecase users.UserUsecaseInterface
var dataUser  users.Domain
var dataUsers []users.Domain
const emailTo = "stevenhumam69@gmail.com"
const pointerUser = "*users.Domain"

func setup() {
	config := config2.GetConfigTest()
	configJWT := middlewares.ConfigJWT{
		SecretJWT: config.JWT_SECRET,
		ExpiresDuration: config.JWT_EXPIRED,
	}
	userUsecase = users.NewUsecase(&userRepositoryMock,&configJWT,time.Hour*1)
	dataUser = users.Domain{
		ID:5,
		FullName:"kafka",
		Email:emailTo,
		Password:"$2a$04$8T1.pXn3uEXzS3how4kokOhHMuPGl8aQhaT4qtsG8U.qmqAH5OEjy",
		Job:"serabutan",
		RoleID:2,
		Token:"",
	}
	dataUsers = append(dataUsers,dataUser)
}

func TestRegisterUser(t *testing.T) {
	setup()
	t.Run("Success Registering User", func(t *testing.T) {
		userRepositoryMock.On("Register",
			mock.AnythingOfType(pointerUser),
			).Return(dataUser,nil).Once()
		resp,err := userUsecase.RegisterUseCase(dataUser)
		assert.Nil(t, err)
		assert.Equal(t, "kafka",resp.FullName)
	})

	t.Run("Fail Registering User", func(t *testing.T) {
		userRepositoryMock.On("Register",
			mock.AnythingOfType(pointerUser),
		).Return(users.Domain{},errors.New("failed to create record")).Once()
		resp,err := userUsecase.RegisterUseCase(dataUser)
		assert.Error(t, err)
		assert.Equal(t, fmt.Errorf("failed to registering user"),err)
		assert.Equal(t, users.Domain{},resp)
	})

	t.Run("Fail Registering User case db error", func(t *testing.T) {
		userRepositoryMock.On("Register",
			mock.AnythingOfType(pointerUser),
		).Return(users.Domain{},errors.New("db err")).Once()
		resp,err := userUsecase.RegisterUseCase(dataUser)
		assert.Error(t, err)
		assert.Equal(t, users.Domain{},resp)
	})


}

func TestLoginUseCase(t *testing.T) {
	setup()
	t.Run("valid login", func(t *testing.T) {
		userRepositoryMock.On("CheckLogin",
				mock.AnythingOfType("string"),
				mock.AnythingOfType("string"),
		).Return(dataUser,nil).Once()
	user, err := userUsecase.LoginUseCase(emailTo,"johnlennon")
	assert.Nil(t, err)
	assert.Equal(t, "kafka",user.FullName)
	})

	t.Run("invalid login", func(t *testing.T) {
		userRepositoryMock.On("CheckLogin",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(users.Domain{},nil).Once()
		user, err := userUsecase.LoginUseCase(emailTo,"johnlennon")
		assert.Error(t, err)
		assert.Equal(t, users.Domain{},user)
	})

	t.Run("error internal", func(t *testing.T) {
		userRepositoryMock.On("CheckLogin",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(users.Domain{},errors.New("internal error")).Once()
		user, err := userUsecase.LoginUseCase(emailTo,"johnlennon")
		assert.Error(t, err)
		assert.Equal(t, users.Domain{},user)
	})

}

func TestGetAll(t *testing.T) {
	setup()
	t.Run("Success Get all Data", func(t *testing.T) {
		userRepositoryMock.On("GetAllUser").Return(dataUsers,nil).Once()
		resp,err := userUsecase.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Fail Get all Data", func(t *testing.T) {
		userRepositoryMock.On("GetAllUser").Return([]users.Domain{},errors.New("error data tidak ditemukan")).Once()
		resp,err := userUsecase.GetAll()
		assert.Error(t, err)
		assert.Equal(t, []users.Domain{},resp)
	})
}

func TestDeleteByID(t *testing.T) {
	setup()
	t.Run("Success Delete", func(t *testing.T) {
		userRepositoryMock.On("DeleteUserByID",
			mock.AnythingOfType("int"),
			).Return(1,nil).Once()
		resp,err := userUsecase.DeleteByID(1)
		assert.Nil(t, err)
		assert.Equal(t, resp,"Success")
	})

	t.Run("Data Not Found", func(t *testing.T) {
		userRepositoryMock.On("DeleteUserByID",
			mock.AnythingOfType("int"),
		).Return(0,nil).Once()
		resp,err := userUsecase.DeleteByID(1)
		assert.Error(t, err)
		assert.Equal(t, resp,"Failed")
	})

	t.Run("Error Internal DB", func(t *testing.T) {
		userRepositoryMock.On("DeleteUserByID",
			mock.AnythingOfType("int"),
		).Return(0,errors.New("Connection error")).Once()
		resp,err := userUsecase.DeleteByID(1)
		assert.Error(t, err)
		assert.Equal(t, resp,"Failed")
	})


}