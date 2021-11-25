package reward

import "errors"

type RewardUseCase struct {
	repoReward RewardRepoInterface
}

func NewUsecase(rewardRepo RewardRepoInterface) RewardUsecaseInterface {
	return RewardUseCase{
	rewardRepo,

	}
}

func (r RewardUseCase) CreateReward(domain Domain) (Domain,error) {
	resp, err := r.repoReward.CreateReward(domain)
	if err != nil {
		return Domain{},errors.New("Error Tidak bisa Membuat Reward")
	}
	return resp,nil
}

func (r RewardUseCase) UpdateReward(domain Domain) (Domain,error) {
	resp,err := r.repoReward.UpdateReward(domain)
	if err != nil {
		return Domain{},errors.New("Error Tidak Mengupdate Membuat Reward")
	}
	return resp,nil
}

func (r RewardUseCase) DeleteRewardByID(id int) (string,error) {
	resp,err := r.repoReward.DeleteRewardByID(id)
	if err != nil || resp == "" {
		return "Gagal Menghapus Data/Internal Error",err
	}
	return resp,nil
}

