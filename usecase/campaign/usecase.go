package campaign

import (
	"errors"
)

type CampaignUseCase struct {
	CampaignRepository CampaignRepoInterface
}

func NewCampaignUseCase(campaignRepo CampaignRepoInterface) CampaignUsecaseInterface {
	return &CampaignUseCase{
		campaignRepo,
	}
}

func (u CampaignUseCase) RegisterUseCase(campaign *Domain) (string,error) {
	//log.Println(user)
	resp,err := u.CampaignRepository.CreateCampaign(campaign)
	return resp,err
}

func (u CampaignUseCase) GetByIDUseCase(id int) (Domain,error) {
	resp,err := u.CampaignRepository.FindOneCampaignByID(id)
	if err != nil {
		return Domain{},err
	}
	return resp,nil
}


func (u CampaignUseCase) GetAllCampaignDetail() ([]Users,error) {
	respDump := u.CampaignRepository.ListAllCampaignByUser()
	if len(respDump) == 0 {
		return []Users{},errors.New("Error Data Kosong")
	}
	return respDump,nil
}

func (u CampaignUseCase) ListAllCampaignByUserUseCase(id int) (UserCampaign, error) {
	resp,err := u.CampaignRepository.ListCampaignsByUserID(id)
	if err != nil {
		return resp,err
	}
	return resp, nil
}