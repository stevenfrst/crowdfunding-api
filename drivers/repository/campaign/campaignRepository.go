package campaign

import (
	"errors"
	repoModels "github.com/stevenfrst/crowdfunding-api/drivers/repository"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(gormDb *gorm.DB) campaign.CampaignRepoInterface {
	return &CampaignRepository{
		db: gormDb,
	}
}

func (c CampaignRepository) CreateCampaign(campaign *campaign.Domain) (string,error) {
	result := c.db.Create(repoModels.FromDomainCampaign(campaign))
	//log.Println(reflect.TypeOf(user),result.RowsAffected)
	if result.Error != nil {
		return strconv.Itoa(int(result.RowsAffected)),result.Error
	}
	return "Success",nil
}

func (c CampaignRepository) FindOneCampaignByID(id int) (campaign.Domain,error) {
	var campaign repoModels.Campaign
	err := c.db.Where("id = ?",id).Find(&campaign).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return campaign.ToDomain(),errors.New("Campaign not found")
		}
		return campaign.ToDomain(),errors.New("Error getting campaign")
	}
	return campaign.ToDomain(),nil
}


func (c CampaignRepository) FindByID(ID int) (campaign.Domain,error) {
	var campaignQuery repoModels.Campaign
	err := c.db.First(&campaignQuery,ID).Error
	if err != nil {
		return campaignQuery.ToDomain(),err
	}
	return campaignQuery.ToDomain(), nil
}

//
//
func (c CampaignRepository) ListCampaignsByUserID(id int) (campaign.UserCampaign,error) {
	var usersQuery repoModels.User
	err := c.db.Preload("Campaigns").Where("id = ?",id).Find(&usersQuery).Error
	//log.Println(usersQuery)
	if err != nil {
		return repoModels.ConvertRepoUserCampaign(usersQuery),err
	}
	return repoModels.ConvertRepoUserCampaign(usersQuery),nil
}


func (c CampaignRepository) UpdateCampaign(campaign campaign.Domain) (campaign.Domain,error) {
	err := c.db.Save(&campaign).Error
	if err != nil {
		return campaign,err
	}
	return campaign,nil
}

func (c CampaignRepository) ListAllCampaignByUser() []campaign.Users {
	var users []repoModels.User
	c.db.Preload("Campaigns").Find(&users)
	log.Println(users)
	return repoModels.ConvertRepoUseCaseUserCampaign(users)
}



