package reward_test

import (
	"errors"
	"github.com/stevenfrst/crowdfunding-api/usecase/reward"
	"github.com/stevenfrst/crowdfunding-api/usecase/reward/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var rewardMockRepo mocks.RewardRepoInterface
var rewardUsecase reward.RewardUsecaseInterface
var dataDummy reward.Domain
var dataHistoryDummy reward.DomainHistory

func setup() {
	rewardUsecase = reward.NewUsecase(&rewardMockRepo)
	dataDummy = reward.Domain{
		ID:1,
		Amount: 20000,
		RewardDescription: "Hahahihi",
	}
}


func TestUpdateReward(t *testing.T) {
	setup()
	t.Run("Success Update Reward", func(t *testing.T) {
		rewardMockRepo.On("UpdateReward",mock.AnythingOfType("reward.Domain")).
			Return(dataDummy,nil).Once()
		resp,err := rewardUsecase.UpdateReward(dataDummy)
		assert.Nil(t, err)
		assert.Equal(t, dataDummy.ID,resp.ID)
	})
	t.Run("Failed Update Reward", func(t *testing.T) {
		rewardMockRepo.On("UpdateReward",mock.AnythingOfType("reward.Domain")).
			Return(reward.Domain{},errors.New("Data Tidak Ditemukan/Gagal Mengupdate/Error Internal")).Once()
		resp,err := rewardUsecase.UpdateReward(dataDummy)
		assert.Error(t, err)
		assert.Equal(t, reward.Domain{},resp)
	})
}


func TestDeleteRewardByID(t *testing.T) {
	setup()
	t.Run("Success Delete ID", func(t *testing.T) {
		rewardMockRepo.On("DeleteRewardByID",mock.AnythingOfType("int")).
			Return("Success",nil).Once()
		resp,err := rewardUsecase.DeleteRewardByID(1)
		assert.Nil(t, err)
		assert.Equal(t, "Success",resp)
	})
	t.Run("Failed Delete ID", func(t *testing.T) {
		rewardMockRepo.On("DeleteRewardByID",mock.AnythingOfType("int")).
			Return("",nil).Once()
		_,err := rewardUsecase.DeleteRewardByID(1)
		assert.Nil(t, err)
		assert.Equal(t, err,err)
	})
}

func TestCreateReward(t *testing.T) {
	setup()
	t.Run("Success Create Reward", func(t *testing.T) {
		rewardMockRepo.On("CreateReward",mock.AnythingOfType("reward.Domain")).
			Return(dataDummy,nil).Once()
		resp,err := rewardUsecase.CreateReward(dataDummy)
		assert.Nil(t, err)
		assert.Equal(t, dataDummy.ID,resp.ID)
	})

	t.Run("Success Create Reward", func(t *testing.T) {
		rewardMockRepo.On("CreateReward",mock.AnythingOfType("reward.Domain")).
			Return(dataDummy,nil).Once()
		resp,err := rewardUsecase.CreateReward(dataDummy)
		assert.Nil(t, err)
		assert.Equal(t, dataDummy.ID,resp.ID)
	})

	t.Run("Failed Create Reward", func(t *testing.T) {
		rewardMockRepo.On("CreateReward",mock.AnythingOfType("reward.Domain")).
			Return(reward.Domain{},errors.New("")).Once()
		resp,err := rewardUsecase.CreateReward(dataDummy)
		assert.Error(t, err)
		assert.Equal(t, reward.Domain{},resp)
	})
}
