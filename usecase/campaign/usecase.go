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

func (u CampaignUseCase) RegisterUseCase(campaign *Domain) (Domain,error) {
	//log.Println(user)
	resp,err := u.CampaignRepository.CreateCampaign(campaign)
	if err != nil {
		return Domain{},errors.New("internal error")
	}
	return resp,err
}

func (u *CampaignUseCase) GetByIDUseCase(id int) (Domain,error) {
	resp,err := u.CampaignRepository.FindOneCampaignByID(id)
	if err != nil {
		return Domain{},err
	}
	return resp,nil
}


func (u CampaignUseCase) GetAllCampaignDetail() ([]Users,error) {
	respDump := u.CampaignRepository.ListAllCampaignByUser()
	if len(respDump) == 0 {
		return []Users{},errors.New("error empty data")
	}
	return respDump,nil
}
// refactor
func (u CampaignUseCase) ListAllCampaignByUserUseCase(id int) (UserCampaign, error) {
	resp,err := u.CampaignRepository.ListCampaignsByUserID(id)
	if err != nil {
		return resp,errors.New("error getting data")
	} else if  resp.ID == 0 {
		return resp,errors.New("data not found")
	}
	return resp, nil
}

func (u CampaignUseCase) EditTargetCampaign(id,target int) (Domain,error) {
	resp,err := u.CampaignRepository.EditTargetCampaign(id,target)
	if err != nil {
		return Domain{},errors.New("Data Not Found")
	}
	return resp,nil
}