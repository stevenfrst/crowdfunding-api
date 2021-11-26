package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	campaignDelivery "github.com/stevenfrst/crowdfunding-api/delivery/campaign"
	rewardDelivery "github.com/stevenfrst/crowdfunding-api/delivery/reward"
	transactionDelivery "github.com/stevenfrst/crowdfunding-api/delivery/transaction"
	userDelivery "github.com/stevenfrst/crowdfunding-api/delivery/users"
)

type RouteControllerList struct {
	UserDelivery userDelivery.UserDelivery
	CampaignDelivery campaignDelivery.CampaignDelivery
	TransactionDelivery transactionDelivery.TransactionDelivery
	RewardDelivery rewardDelivery.RewardDelivery
	JWTConfig      middleware.JWTConfig
}


func (d RouteControllerList) RouteRegister(c *echo.Echo) {
	jwt := middleware.JWTWithConfig(d.JWTConfig)

	c.POST("/v1/login",d.UserDelivery.Login)
	c.POST("/v1/register",d.UserDelivery.Register)
	//c.GET("/swagger/*", echoSwagger.WrapHandler)


	c.POST("/v1/campaign",d.CampaignDelivery.CreateCampaignHandler,jwt)
	c.GET("/v1/campaign/:id",d.CampaignDelivery.GetCampaignByID,jwt)
	c.POST("/v1/campaign/edit/target",d.CampaignDelivery.EditCampaignTargetByID,jwt)

	c.GET("/v1/user/:id/campaign",d.CampaignDelivery.GetAllCampaignByUserID,jwt)
	c.GET("/v1/user/:id/transaction",d.UserDelivery.GetUserTransaction,jwt)
	c.GET("/v1/user/campaign",d.CampaignDelivery.GetAllCampaignDetail,jwt)
	c.GET("/v1/user/all",d.UserDelivery.GetAll,jwt)
	c.DELETE("/v1/user/:id",d.UserDelivery.DeleteByID,jwt)
	c.POST("/v1/user/edit/password",d.UserDelivery.UpdatePassword,jwt)

	c.POST("/v1/payments/create",d.TransactionDelivery.CreateTransaction,jwt)
	c.POST("/v1/payments/notification",d.TransactionDelivery.GetNotificationPayment)
	c.GET("/v1/payments/:id",d.TransactionDelivery.GetStatusByID,jwt)

	c.POST("/v1/reward/create",d.RewardDelivery.CreateReward,jwt)
	c.POST("/v1/reward/update",d.RewardDelivery.UpdateReward,jwt)
	c.DELETE("/v1/reward/delete/:id",d.RewardDelivery.DeleteRewardByID,jwt)

}
