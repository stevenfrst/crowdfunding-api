package response

import (
	userResponse "github.com/stevenfrst/crowdfunding-api/delivery/users/response"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
)

type CampaignResponse struct {
	ID	int `json:"id"`
	UserID int `json:"user_id" validate:"required"`
	CampaignName string `json:"campaign_name" validate:"required"`
	ShortDescription string `json:"short_description" validate:"required"`
	LongDescription string `json:"long_description" validate:"required"`
	Target int `json:"target" validate:"required"`
	AmountNow int `json:"amount_now"`
	Supporters int `json:"supporters"`
}

func FromDomainCampaignUserAll(domain []campaign.Users) (response []userResponse.UserResponseWCampaign) {
	for _,x := range domain {
		newResponse := userResponse.UserResponseWCampaign{
			Id:       int(x.ID),
			FullName: x.FullName,
			Email:    x.Email,
			Password: x.Password,
			Job:      x.Job,
			RoleID: int(x.RoleID),
			Campaign: x.Campaigns,
			//Token string
		}
		response = append(response, newResponse)
	}
	return response
}

func FromDomainCampaignUser(domain campaign.UserCampaign) (response userResponse.UserResponseWCampaign) {
	return userResponse.UserResponseWCampaign{
		Id:       int(domain.ID),
		FullName: domain.FullName,
		Email:    domain.Email,
		Password: domain.Password,
		Job:      domain.Job,
		RoleID: int(domain.RoleID),
		Campaign: domain.Campaigns,
		//Token string
	}
}

func FromDomainCampaign(domain campaign.Domain) CampaignResponse {
	return CampaignResponse{
		ID:     int(domain.ID),
		UserID: int(domain.UserID),
		CampaignName: domain.CampaignName,
		ShortDescription: domain.ShortDescription,
		LongDescription: domain.LongDescription,
		Target:domain.Target,
		AmountNow: domain.AmountNow,
		Supporters:domain.Supporters,
	}
}