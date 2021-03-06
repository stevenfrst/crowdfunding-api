package campaign

import (
	"errors"
	repoModels "github.com/stevenfrst/crowdfunding-api/drivers/repository"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	"gorm.io/gorm"
	"log"
)

type CampaignRepository struct {
	db *gorm.DB
}
// NewCampaignRepository handle all campaign repositories methods
func NewCampaignRepository(gormDb *gorm.DB) campaign.CampaignRepoInterface {
	return &CampaignRepository{
		db: gormDb,
	}
}

// GetLast methods to getting last id
func (c CampaignRepository) GetLast() int {
	var campaign repoModels.Campaign
	c.db.Last(&campaign)
	return int(campaign.ID)
}

// CreateCampaign methods to create new campaign
func (c CampaignRepository) CreateCampaign(campaignIn *campaign.Domain) (campaign.Domain,error) {
	result := c.db.Create(repoModels.FromDomainCampaign(campaignIn))
	if result.Error != nil {
		return campaign.Domain{},result.Error
	}

	campaignIn.ID = uint(c.GetLast())
	return *campaignIn,nil
}

// FindOneCampaignByID methods to get 1 campaign by id
func (c CampaignRepository) FindOneCampaignByID(id int) (campaign.Domain,error) {
	var campaign repoModels.Campaign
	err := c.db.Where("id = ?",id).Find(&campaign).Error

	if err != nil || campaign.ID == 0{
		if err == gorm.ErrRecordNotFound {
			return campaign.ToDomain(),errors.New("Campaign not found")
		}
		return campaign.ToDomain(),errors.New("Error getting campaign")
	}
	return campaign.ToDomain(),nil
}

// FindByID method to find id via id return campaign domain to use case
func (c CampaignRepository) FindByID(ID int) (campaign.Domain,error) {
	var campaignQuery repoModels.Campaign
	err := c.db.First(&campaignQuery,ID).Error
	if err != nil {
		return campaignQuery.ToDomain(),err
	}
	return campaignQuery.ToDomain(), nil
}


func (c CampaignRepository) ListCampaignsByUserID(id int) (campaign.UserCampaign,error) {
	var usersQuery repoModels.User
	err := c.db.Preload("Campaigns").Where("id = ?",id).Find(&usersQuery).Error
	log.Println(usersQuery)
	if err != nil || usersQuery.ID == 0 {
		return repoModels.ConvertRepoUserCampaign(usersQuery),errors.New("Data Tidak Ditemukan/Error")
	}
	return repoModels.ConvertRepoUserCampaign(usersQuery),nil
}


func (c CampaignRepository) UpdateCampaign(campaignInput campaign.Domain) (campaign.Domain,error) {
	campaign := repoModels.FromDomainCampaign(&campaignInput)
	err := c.db.Save(&campaign).Error
	if err != nil {
		return campaign.ToDomain(),err
	}
	return campaign.ToDomain(),nil
}

func (c CampaignRepository) ListAllCampaignByUser() []campaign.Users {
	var users []repoModels.User
	c.db.Preload("Campaigns").Find(&users)
	log.Println(users)
	return repoModels.ConvertRepoUseCaseUserCampaign(users)
}

func (c CampaignRepository) EditTargetCampaign(id,target int) (campaign.Domain,error) {
	campaignIn, err := c.FindOneCampaignByID(id)
	if err != nil {
		return campaign.Domain{},err
	}
	campaignIn.Target = target
	campaignRepo := repoModels.FromDomainCampaign(&campaignIn)
	err = c.db.Save(&campaignRepo).Error
	if err != nil {
		return campaign.Domain{},err
	}
	return campaignRepo.ToDomain(),nil
}