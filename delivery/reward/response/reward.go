package response

import "github.com/stevenfrst/crowdfunding-api/usecase/reward"

type Reward struct {
	ID        uint `json:"id"`
	Amount int `json:"amount"`
	RewardDescription string `json:"reward_description"`
}

func FromDomainReward(reward reward.Domain) Reward {
	return Reward{
		ID: reward.ID,
		Amount: reward.Amount,
		RewardDescription: reward.RewardDescription,
	}
}