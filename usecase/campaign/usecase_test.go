package campaign_test

import (
	"errors"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"log"
	"testing"
)

var campaignMockRepo mocks.CampaignRepoInterface
var campaignUseCase campaign.CampaignUsecaseInterface
var dataCampaign  campaign.Domain
var dataUser []campaign.Users
var singleDataUser campaign.UserCampaign

func setup() {
	campaignUseCase = campaign.NewCampaignUseCase(&campaignMockRepo)
	dataCampaign = campaign.Domain{
		ID: 1,
		UserID: 3,
		CampaignName: "tahu bulat digoreng",
		ShortDescription :"di mobil enak",
		LongDescription :"500 ratusan mauuuu",
		Target:100000,
		AmountNow :650000,
		Supporters :8,
	}
	dummyUser := campaign.Users{
		ID:1,
		FullName: "admin",
		Job: "Administrator",
		Email: "mail@admin.com",
		Password: "password",
		RoleID: 1,
	}
	dataUser =append(dataUser,dummyUser)
	singleDataUser=campaign.UserCampaign{
		ID:dummyUser.ID,
		FullName:dummyUser.FullName,
		Email:dummyUser.Email,
		Password: dummyUser.Password,
		Job: dummyUser.Job,
		RoleID: dummyUser.RoleID,
		Campaigns: dummyUser.Campaigns,
	}
}

//GetByIDUseCase(id int) (Domain,error)

func TestListAllCampaignByUserUseCase(t *testing.T) {
	setup()
	t.Run("Success Get all Data", func(t *testing.T) {
		campaignMockRepo.On("ListCampaignsByUserID",
			mock.AnythingOfType("int")).Return(singleDataUser,nil).Once()
		resp,err := campaignUseCase.ListAllCampaignByUserUseCase(1)
		assert.Nil(t, err)
		assert.Equal(t, singleDataUser.ID,resp.ID)
	})

	t.Run("Failed Get all Data", func(t *testing.T) {
		campaignMockRepo.On("ListCampaignsByUserID",
			mock.AnythingOfType("int")).Return(campaign.UserCampaign{},nil).Once()
		resp,err := campaignUseCase.ListAllCampaignByUserUseCase(1)
		log.Println(resp,err)
		assert.Error(t, err)
		assert.Equal(t, campaign.UserCampaign{},resp)
	})

}

func TestRegisterUseCase(t *testing.T) {
	setup()
	t.Run("Success Register", func(t *testing.T) {
		campaignMockRepo.On("CreateCampaign",
			mock.AnythingOfType("*campaign.Domain")).Return(dataCampaign,nil).Once()
		datamock := campaign.Domain{
			UserID:1,
			CampaignName: dataCampaign.CampaignName,
			ShortDescription:dataCampaign.ShortDescription,
			LongDescription: dataCampaign.LongDescription,
			Target: dataCampaign.Target,
			AmountNow:dataCampaign.AmountNow,
			Supporters:dataCampaign.Supporters,
		}
		resp,err := campaignUseCase.RegisterUseCase(&datamock)
		assert.Nil(t, err)
		assert.Equal(t, "tahu bulat digoreng",resp.CampaignName)
	})
}

func TestGetAllCampaignDetail(t *testing.T) {
	setup()
	t.Run("Success GetAll", func(t *testing.T) {
		campaignMockRepo.On("ListAllCampaignByUser").Return(dataUser).Once()
		resp,err := campaignUseCase.GetAllCampaignDetail()
		log.Println(resp)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Gagal Mengambil Data", func(t *testing.T) {
		campaignMockRepo.On("ListAllCampaignByUser").Return([]campaign.Users{}).Once()
		resp,err := campaignUseCase.GetAllCampaignDetail()
		assert.Error(t, err)
		assert.Equal(t, resp,[]campaign.Users{})
	})



}

func TestGetByIDUseCase(t *testing.T) {
	setup()
	t.Run("Success get data with id", func(t *testing.T) {
		campaignMockRepo.On("FindOneCampaignByID",
			mock.AnythingOfType("int")).Return(dataCampaign,nil).Once()
		resp,err := campaignUseCase.GetByIDUseCase(int(dataCampaign.ID))
		assert.Nil(t, err)
		assert.Equal(t, "tahu bulat digoreng",resp.CampaignName)
	})

	t.Run("fail to get data via id", func(t *testing.T){
		campaignMockRepo.On("FindOneCampaignByID",
			mock.AnythingOfType("int")).Return(campaign.Domain{},errors.New("Error getting campaign")).Once()
		resp,err := campaignUseCase.GetByIDUseCase(int(dataCampaign.ID))
		assert.Error(t, err)
		assert.Equal(t, resp,campaign.Domain{})
	})
}
