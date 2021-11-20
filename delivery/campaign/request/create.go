package request

import "github.com/stevenfrst/crowdfunding-api/usecase/campaign"

type CampaignRequest struct {
	ID	int `json:"id"`
	UserID int `json:"user_id" validate:"required"`
	CampaignName string `json:"campaign_name" validate:"required"`
	ShortDescription string `json:"short_description" validate:"required"`
	LongDescription string `json:"long_description" validate:"required"`
	Target int `json:"target" validate:"required"`
	AmountNow int `json:"amount_now"`
	Supporters int `json:"supporters"`
}

func (c CampaignRequest) ToDomain() *campaign.Domain {
	return &campaign.Domain{
		ID:               uint(c.ID),
		UserID: uint(c.UserID),
		CampaignName:     c.CampaignName,
		ShortDescription: c.ShortDescription,
		LongDescription:  c.LongDescription,
		Target:           c.Target,
		AmountNow:        c.AmountNow,
		Supporters:       c.Supporters,
	}
}