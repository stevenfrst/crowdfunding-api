package reward

import (
	repoModels "github.com/stevenfrst/crowdfunding-api/drivers/repository"
	"github.com/stevenfrst/crowdfunding-api/usecase/reward"
	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(gormDb *gorm.DB) reward.RewardRepoInterface {
	return &RewardRepository{
		db: gormDb,
	}
}

func(r RewardRepository) GetRewardByAmount(amount int) (int,string,error) {
	var reward repoModels.Reward
	err := r.db.Where("amount = ?",amount).Find(&reward).Error
	if err != nil {
		return 0,"",err
	}
	return int(reward.ID),reward.RewardDescription,nil
}

func (r RewardRepository) SaveRewardHistory(domain reward.DomainHistory) error {
	err := r.db.Save(repoModels.FromDomainHistory(&domain)).Error
	if err != nil {
		return err
	}
	return nil
}

func (r RewardRepository) CreateReward(domain reward.Domain) (reward.Domain,error) {
	rewardIn := repoModels.FromDomainReward(domain)
	err := r.db.Create(&rewardIn).Error
	if err != nil {
		return reward.Domain{},err
	}
	return rewardIn.ToDomain(),nil
}


func (r RewardRepository) UpdateReward(domain reward.Domain) (reward.Domain,error) {
	rewardIn := repoModels.FromDomainReward(domain)
	err := r.db.Where("id = ?",domain.ID).Save(&rewardIn).Error
	if err != nil {
		return reward.Domain{},err
	}
	return rewardIn.ToDomain(),nil
}

func (r RewardRepository) DeleteRewardByID(id int) (string,error) {
	var reward repoModels.Reward
	err := r.db.Where("id = ?",id).Delete(&reward).Error
	if err != nil {
		return "", err
	}
	return "Success",nil
}




