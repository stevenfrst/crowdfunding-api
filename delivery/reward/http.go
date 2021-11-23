package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/stevenfrst/crowdfunding-api/delivery"
	"github.com/stevenfrst/crowdfunding-api/delivery/reward/request"
	"github.com/stevenfrst/crowdfunding-api/delivery/reward/response"
	"github.com/stevenfrst/crowdfunding-api/usecase/reward"
	"net/http"
	"strconv"
)

type RewardDelivery struct {
	usecase reward.RewardUsecaseInterface
}

func NewRewardDelivery(cc reward.RewardUsecaseInterface) *RewardDelivery {
	return &RewardDelivery{
		usecase: cc,
	}
}


func (r RewardDelivery) CreateReward(c echo.Context) error {
	var rewardRequest request.RewardRequest
	if err := c.Bind(&rewardRequest); err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,err.Error(),err)
	}
	resp, err := r.usecase.CreateReward(rewardRequest.ToDomain())
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)
	}
	return delivery.SuccessResponse(c,response.FromDomainReward(resp))
}


func (r RewardDelivery) UpdateReward(c echo.Context) error {
	var rewardRequest request.RewardRequest
	if err := c.Bind(&rewardRequest); err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,err.Error(),err)
	}
	resp, err := r.usecase.UpdateReward(rewardRequest.ToDomain())
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)
	}
	return delivery.SuccessResponse(c,response.FromDomainReward(resp))
}

func (r RewardDelivery) DeleteRewardByID(c echo.Context) error {
	idParam,_ := strconv.Atoi(c.Param("id"))
	resp,err := r.usecase.DeleteRewardByID(idParam)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,"Failed",err)
	}
	return delivery.SuccessResponse(c,resp)

}