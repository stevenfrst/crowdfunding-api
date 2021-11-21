package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/stevenfrst/crowdfunding-api/delivery"
	"github.com/stevenfrst/crowdfunding-api/delivery/users/request"
	"github.com/stevenfrst/crowdfunding-api/delivery/users/response"
	"github.com/stevenfrst/crowdfunding-api/usecase/users"
	"net/http"
)

type UserDelivery struct {
	usecase users.UserUsecaseInterface
}

func NewUserDelivery(uc users.UserUsecaseInterface) *UserDelivery {
	return &UserDelivery{
		usecase: uc,
	}
}

func (d *UserDelivery) Register(c echo.Context) (err error) {
	ctx := c.Request().Context()
	var user request.UserRegister
	if err = c.Bind(&user);err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(&user); err != nil {
		return err
	}

	out,err := d.usecase.RegisterUseCase(user.ToDomain(),ctx)
	//log.Println(out)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"error",err)
	}


	return delivery.SuccessResponse(c,out)
}

func (d *UserDelivery) Login(c echo.Context) error {
	ctx := c.Request().Context()

	email := c.FormValue("email")
	password := c.FormValue("password")

	res,err := d.usecase.LoginUseCase(email,password,ctx)
	if err != nil {
		return delivery.ErrorResponse(c, http.StatusInternalServerError, "error", err)
	}

	return delivery.SuccessResponse(c,response.FromDomain(res))
}

func (d *UserDelivery) GetAll(c echo.Context) error {
	res,err := d.usecase.GetAll()
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,"Failed",err)
	}
	return delivery.SuccessResponse(c,response.FromDomainList(res))
}