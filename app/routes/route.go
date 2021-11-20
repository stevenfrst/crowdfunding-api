package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	campaignDelivery "github.com/stevenfrst/crowdfunding-api/delivery/campaign"
	userDelivery "github.com/stevenfrst/crowdfunding-api/delivery/users"
)

type RouteControllerList struct {
	UserDelivery userDelivery.UserDelivery
	CampaignDelivery campaignDelivery.CampaignDelivery
	JWTConfig      middleware.JWTConfig
}


func (d RouteControllerList) RouteRegister(c *echo.Echo) {
	c.POST("/v1/login",d.UserDelivery.Login)
	c.POST("/v1/register",d.UserDelivery.Register)

	c.POST("/v1/campaign",d.CampaignDelivery.CreateCampaignHandler)
	c.GET("/v1/campaign/:id",d.CampaignDelivery.GetCampaignByID)

	c.GET("/v1/user/:id/campaign",d.CampaignDelivery.GetAllCampaignByUserID)
	c.GET("/v1/user/campaign",d.CampaignDelivery.GetAllCampaignDetail)

}
