package campaign

import (
	"errors"
)

type CampaignUseCase struct {
	CampaignRepository CampaignRepoInterface
}
// NewCampaignUseCase function creates a new CampaignUseCaseInterface
func NewCampaignUseCase(campaignRepo CampaignRepoInterface) CampaignUsecaseInterface {
	return &CampaignUseCase{
		campaignRepo,
	}
}

// RegisterUseCase method registers a user
func (u CampaignUseCase) RegisterUseCase(campaign *Domain) (Domain,error) {
	resp,err := u.CampaignRepository.CreateCampaign(campaign)
	if err != nil {
		return Domain{},errors.New("internal error")
	}
	return resp,err
}

// GetByIDUseCase method returns a domain by id
func (u *CampaignUseCase) GetByIDUseCase(id int) (Domain,error) {
	resp,err := u.CampaignRepository.FindOneCampaignByID(id)
	if err != nil {
		return Domain{},err
	}
	return resp,nil
}

// GetAllCampaignDetail methods return a list of all campaign user have
func (u CampaignUseCase) GetAllCampaignDetail() ([]Users,error) {
	respDump := u.CampaignRepository.ListAllCampaignByUser()
	if len(respDump) == 0 {
		return []Users{},errors.New("error empty data")
	}
	return respDump,nil
}
// ListAllCampaignByUserUseCase methods return a list of a user have via id
func (u CampaignUseCase) ListAllCampaignByUserUseCase(id int) (UserCampaign, error) {
	resp,err := u.CampaignRepository.ListCampaignsByUserID(id)
	if err != nil {
		return resp,errors.New("error getting data")
	} else if  resp.ID == 0 {
		return resp,errors.New("data not found")
	}
	return resp, nil
}

// EditTargetCampaign methods edit a campaign target via id and desired target
func (u CampaignUseCase) EditTargetCampaign(id,target int) (Domain,error) {
	resp,err := u.CampaignRepository.EditTargetCampaign(id,target)
	if err != nil {
		return Domain{},errors.New("Data Not Found")
	}
	return resp,nil
}