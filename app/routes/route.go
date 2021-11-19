package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	userDelivery "github.com/stevenfrst/crowdfunding-api/delivery/users"
)

type RouteControllerList struct {
	UserDelivery userDelivery.UserDelivery
	JWTConfig      middleware.JWTConfig
}


func (d RouteControllerList) RouteRegister(c *echo.Echo) {
	c.POST("/v1/login",d.UserDelivery.Login)
	c.POST("/v1/register",d.UserDelivery.Register)
}
