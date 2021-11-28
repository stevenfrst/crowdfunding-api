package repoModels

import "github.com/stevenfrst/crowdfunding-api/usecase/campaign"

// FromDomainCampaign convert campaign repo model to campaign domain
func FromDomainCampaign(domain *campaign.Domain) *Campaign {
	return &Campaign {
		ID:domain.ID,
		UserID:domain.UserID,
		CampaignName:domain.CampaignName,
		ShortDescription: domain.ShortDescription,
		LongDescription: domain.LongDescription,
		Target: domain.Target,
		AmountNow: domain.AmountNow,
		Supporters:domain.Supporters,
	}
}

// ToDomain convert campaign repo to campaign domain
func (c Campaign) ToDomain() campaign.Domain {
	return campaign.Domain{
		ID:c.ID,
		UserID:c.UserID,
		CampaignName:c.CampaignName,
		ShortDescription: c.ShortDescription,
		LongDescription: c.LongDescription,
		Target: c.Target,
		AmountNow: c.AmountNow,
		Supporters:c.Supporters,
	}
}
