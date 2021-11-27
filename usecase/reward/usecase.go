package reward

import "errors"

type RewardUseCase struct {
	repoReward RewardRepoInterface
}

// NewUsecase function creates a new RewardUseCaseInterface
func NewUsecase(rewardRepo RewardRepoInterface) RewardUsecaseInterface {
	return RewardUseCase{
	rewardRepo,

	}
}

// CreateReward method to create a new Reward
func (r RewardUseCase) CreateReward(domain Domain) (Domain,error) {
	resp, err := r.repoReward.CreateReward(domain)
	if err != nil {
		return Domain{},errors.New("Error Tidak bisa Membuat Reward")
	}
	return resp,nil
}

// UpdateReward method to update reward
func (r RewardUseCase) UpdateReward(domain Domain) (Domain,error) {
	resp,err := r.repoReward.UpdateReward(domain)
	if err != nil {
		return Domain{},errors.New("Error Tidak Mengupdate Membuat Reward")
	}
	return resp,nil
}

// DeleteRewardByID methods to delete reward by id
func (r RewardUseCase) DeleteRewardByID(id int) (string,error) {
	resp,err := r.repoReward.DeleteRewardByID(id)
	if err != nil || resp == "" {
		return "Gagal Menghapus Data/Internal Error",err
	}
	return resp,nil
}

