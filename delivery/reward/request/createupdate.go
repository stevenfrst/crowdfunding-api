package request

import "github.com/stevenfrst/crowdfunding-api/usecase/reward"

type RewardRequest struct {
	ID        uint `json:"id"`
	Amount int `json:"amount"`
	RewardDescription string `json:"reward_description"`
}

func (r RewardRequest) ToDomain() reward.Domain {
	return reward.Domain{
		ID:r.ID,
		Amount: r.Amount,
		RewardDescription: r.RewardDescription,
	}
}