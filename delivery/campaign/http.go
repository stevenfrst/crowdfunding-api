package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/stevenfrst/crowdfunding-api/delivery"
	"github.com/stevenfrst/crowdfunding-api/delivery/campaign/request"
	"github.com/stevenfrst/crowdfunding-api/delivery/campaign/response"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type CampaignDelivery struct {
	usecase campaign.CampaignUsecaseInterface
}
// NewCampaignDelivery function to manage all handlers in campaign
func NewCampaignDelivery(cc campaign.CampaignUsecaseInterface) *CampaignDelivery {
	return &CampaignDelivery{
		usecase: cc,
	}
}

// CreateCampaignHandler handler to handle creating campaigns with json
func (u CampaignDelivery) CreateCampaignHandler(c echo.Context) (err error) {
	var campaignIn request.CampaignRequest
	if err = c.Bind(&campaignIn);err != nil {
		//return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return delivery.ErrorResponse(c,http.StatusBadRequest,err.Error(),err)
	}
	if err = c.Validate(&campaignIn); err != nil {
		return err
	}

	out,err := u.usecase.RegisterUseCase(campaignIn.ToDomain())
	//log.Println(out)
	if err != nil {
		//return echo.NewHTTPError(http.StatusInternalServerError,err)
		return delivery.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)
	}


	return delivery.SuccessResponse(c,response.FromDomainCampaign(out))
}

// GetCampaignByID handler to take care of creating a campaign via id using params
func (u CampaignDelivery) GetCampaignByID(c echo.Context) (err error) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,err.Error(),err)
	}
	resp, err := u.usecase.GetByIDUseCase(id)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,err.Error(),err)
	}

	return delivery.SuccessResponse(c,response.FromDomainCampaign(resp))
}

// GetAllCampaignDetail handler to take care of getting all campaign details
func (u CampaignDelivery) GetAllCampaignDetail(c echo.Context) (err error) {

	if err != nil {
		//return echo.NewHTTPError(http.StatusInternalServerError,err)
		return delivery.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)
	}
	resp, err := u.usecase.GetAllCampaignDetail()
	if err != nil {
		//return echo.NewHTTPError(http.StatusInternalServerError,err)
		//log.Println(err)
		return delivery.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)
	}
	log.Println(reflect.TypeOf(resp))
	return delivery.SuccessResponse(c,response.FromDomainCampaignUserAll(resp))
}

// GetAllCampaignByUserID handler to take care of getting all campaigns owned by the user via id
func (u CampaignDelivery) GetAllCampaignByUserID(c echo.Context) (err error) {
	idParam := c.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil {
		//return echo.NewHTTPError(http.StatusInternalServerError,err)
		return delivery.ErrorResponse(c,http.StatusBadRequest,err.Error(),err)

	}
	resp, err := u.usecase.ListAllCampaignByUserUseCase(id)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusInternalServerError,err.Error(),err)

	}

	return delivery.SuccessResponse(c,response.FromDomainCampaignUser(resp))
}

// EditCampaignTargetByID handler to edit target campaign via id
func (u CampaignDelivery) EditCampaignTargetByID(c echo.Context) (err error) {
	id,_ := strconv.Atoi(c.FormValue("id"))
	target,_ := strconv.Atoi(c.FormValue("target"))

	resp,err := u.usecase.EditTargetCampaign(id,target)
	if err != nil {
		return delivery.ErrorResponse(c,http.StatusBadRequest,"Failed",err)
	}
	return delivery.SuccessResponse(c,response.FromDomainCampaign(resp))
}