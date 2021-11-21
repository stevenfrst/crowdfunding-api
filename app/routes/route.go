package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	campaignDelivery "github.com/stevenfrst/crowdfunding-api/delivery/campaign"
	transactionDelivery "github.com/stevenfrst/crowdfunding-api/delivery/transaction"
	userDelivery "github.com/stevenfrst/crowdfunding-api/delivery/users"
)

type RouteControllerList struct {
	UserDelivery userDelivery.UserDelivery
	CampaignDelivery campaignDelivery.CampaignDelivery
	TransactionDelivery transactionDelivery.TransactionDelivery
	JWTConfig      middleware.JWTConfig
}


func (d RouteControllerList) RouteRegister(c *echo.Echo) {
	jwt := middleware.JWTWithConfig(d.JWTConfig)

	c.POST("/v1/login",d.UserDelivery.Login)
	c.POST("/v1/register",d.UserDelivery.Register)


	c.POST("/v1/campaign",d.CampaignDelivery.CreateCampaignHandler)
	c.GET("/v1/campaign/:id",d.CampaignDelivery.GetCampaignByID)

	c.GET("/v1/user/:id/campaign",d.CampaignDelivery.GetAllCampaignByUserID)
	c.GET("/v1/user/campaign",d.CampaignDelivery.GetAllCampaignDetail,jwt)
	c.GET("/v1/user/all",d.UserDelivery.GetAll)
	c.DELETE("/v1/user/:id",d.UserDelivery.DeletaByID)
	c.POST("/v1/user",d.UserDelivery.UpdatePassword)

	c.POST("/v1/payments/create",d.TransactionDelivery.CreateTransaction)
	c.POST("/v1/payments/notification",d.TransactionDelivery.GetNotificationPayment)
	c.GET("/v1/payments/:id",d.TransactionDelivery.GetStatusByID)


}
