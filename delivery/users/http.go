package delivery

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	middlewares "github.com/stevenfrst/crowdfunding-api/app/middleware"
	"github.com/stevenfrst/crowdfunding-api/delivery"
	"github.com/stevenfrst/crowdfunding-api/delivery/users/request"
	"github.com/stevenfrst/crowdfunding-api/delivery/users/response"
	"github.com/stevenfrst/crowdfunding-api/helper/errorHelper"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
	"net/http"
	"strconv"
)

type UserDelivery struct {
	usecase users.UserUsecaseInterface
}

// NewUserDelivery function used to assign all handlers to usecase
func NewUserDelivery(uc users.UserUsecaseInterface) *UserDelivery {
	return &UserDelivery{
		usecase: uc,
	}
}
// Register takes care of the register function of the user and is forwarded to the use case register
func (d *UserDelivery) Register(c echo.Context) (err error) {
	var user request.UserRegister
	if err = c.Bind(&user);err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&user); err != nil {
		return err
	}

	out,err := d.usecase.RegisterUseCase(user.ToDomain())
	//log.Println(out)
	if err != nil {
		//return delivery.ErrorResponse(c,http.StatusInternalServerError,errorHelper.ERROR_USER_REGISTER,err)
		if fmt.Sprintf("%v",err) == "failed to registering user"{
			return delivery.ErrorResponse(c,http.StatusBadRequest,errorHelper.ERROR_USER_REGISTER,err)
		} else {
			return delivery.ErrorResponse(c,http.StatusInternalServerError,errorHelper.ERROR_INTERNAL,err)
		}
	}

	return delivery.SuccessResponse(c,out)
}

// Login handler to handle login from user
func (d *UserDelivery) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	var user request.UserRegister
	user.Email = email
	user.Password = password
	if err := c.Validate(&user); err != nil {
		return delivery.ErrorResponse(c, http.StatusBadRequest, "Format Email/Password Salah", err)
	}
	res,err := d.usecase.LoginUseCase(email,password)
	if err != nil {
		if fmt.Sprintf("%v",err) == "email/password not match" {
			return delivery.ErrorResponse(c, http.StatusUnauthorized, "error", err)
		} else {
			return delivery.ErrorResponse(c,http.StatusInternalServerError,"Internal Error",err)
		}
	}

	return delivery.SuccessResponse(c,response.FromDomain(res))
}

// GetAll handler to take care of retrieving all data from the user
func (d *UserDelivery) GetAll(c echo.Context) error {
	res,err := d.usecase.GetAll()
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"Failed",err)
	}
	return delivery.SuccessResponse(c,response.FromDomainList(res))
}

// DeleteByID handler to take care of deleting data via id
func (d *UserDelivery) DeleteByID(c echo.Context) error {
	idParam := c.Param("id")
	id,_ := strconv.Atoi(idParam)
	res,err := d.usecase.DeleteByID(id)
	if res == "Failed" {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,res,err)
	}
	return delivery.SuccessResponse(c,res)
}

// GetUserTransaction handler to take care of getting all user transactions via id
func (d *UserDelivery) GetUserTransaction(c echo.Context) error {
	idParam,_ := strconv.Atoi(c.Param("id"))
	//var user response.UserResponseWTransaction
	resp,err := d.usecase.GetUserTransactionByID(idParam)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,"Failed",err)
	}
	return delivery.SuccessResponse(c,response.FromDomainUserTransaction(resp))
}

// GetUserJWT handler to take care of getting user jwt (unused)
func(d *UserDelivery) GetUserJWT(c echo.Context) error {
	user := c.Get("UserId").(*jwt.Token)
	claims := user.Claims.(*middlewares.JwtCustomClaims)
	name := strconv.Itoa(claims.UserId)

	return c.String(http.StatusOK,"Welcome "+name )
}

// UpdatePassword handler to handle password update by entering json
func (d *UserDelivery) UpdatePassword(c echo.Context) error {
	var user request.PasswordUpdate
	err := c.Bind(&user)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"Failed to Bind Data",err)
	}
	err = c.Validate(&user)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,"Failed, Wrong Input",err)
	}
	resp,err := d.usecase.UpdatePassword(user.ToDomain())
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,resp,err)
	}
	return delivery.SuccessResponse(c,resp)
}
