package repoModels

import (
	"github.com/stevenfrst/crowdfunding-api/usecase/reward"
)

// FromDomainHistory convert domain reward-history to reward_history repo
func FromDomainHistory(domain *reward.DomainHistory) *RewardHistory {
	return &RewardHistory {
		ID:domain.ID,
		UserID: domain.UserID,
		RewardID:domain.RewardID,
		TransactionID: domain.TransactionID,
	}
}

// ToDomain convert domain-reward repo to reward history domain
func (history RewardHistory) ToDomain() reward.DomainHistory {
	return reward.DomainHistory{
		ID: history.ID,
		UserID: history.UserID,
		RewardID: history.RewardID,
		TransactionID: history.TransactionID,
	}
}

// ToDomain convert reward repo to reward domain
func (r Reward) ToDomain() reward.Domain {
	return reward.Domain{
		ID: r.ID,
		Amount: r.Amount,
		RewardDescription: r.RewardDescription,
	}
}

// FromDomainReward convert domain reward to reward repo
func FromDomainReward(domain reward.Domain) *Reward {
	return &Reward{
		ID: domain.ID,
		Amount: domain.Amount,
		RewardDescription: domain.RewardDescription,
	}
}