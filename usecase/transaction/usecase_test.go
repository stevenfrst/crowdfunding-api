package transaction_test

import (
	"errors"
	"github.com/midtrans/midtrans-go/snap"
	config2 "github.com/stevenfrst/crowdfunding-api/app/config"
	"github.com/stevenfrst/crowdfunding-api/drivers/email"
	payment "github.com/stevenfrst/crowdfunding-api/drivers/midtrans"
	_mockPayment "github.com/stevenfrst/crowdfunding-api/drivers/midtrans/mocks"
	"github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	_mockTransaction "github.com/stevenfrst/crowdfunding-api/usecase/campaign/mocks"
	"github.com/stevenfrst/crowdfunding-api/usecase/reward"
	_mockReward "github.com/stevenfrst/crowdfunding-api/usecase/reward/mocks"
	"github.com/stevenfrst/crowdfunding-api/usecase/transaction"
	"github.com/stevenfrst/crowdfunding-api/usecase/transaction/mocks"
	_mockUser "github.com/stevenfrst/crowdfunding-api/usecase/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gomail.v2"
	"testing"
)

var paymentRepositoryMock _mockPayment.MidtransInterface
var transactionRepoMock mocks.TransactionRepoInterface
var campaignMockRepo _mockTransaction.CampaignRepoInterface
var rewardMockRepo _mockReward.RewardRepoInterface
var mockUserRepo _mockUser.UserRepoInterface
var transactionUsecase transaction.TransactionUsecaseInterface
var domainNotificationDummy transaction.DomainNotification
var dialer *gomail.Dialer
var transactionDummy transaction.Domain
var sampleSnap snap.Response
var dataCampaign  campaign.Domain


func setup() {
	config := config2.GetConfigTest()
	configPayment := payment.ConfigMidtrans{
		SERVER_KEY: config.SERVER_KEY,
	}
	gmail := email.GmailConfig{
		CONFIG_SMTP_HOST:       config.CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:       config.CONFIG_SMTP_PORT,
		CONFIG_SMTP_AUTH_EMAIL: config.CONFIG_SMTP_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD:   config.CONFIG_AUTH_PASSWORD,
		CONFIG_SENDER_NAME:     config.CONFIG_SENDER_NAME,
	}
	dialer = email.NewGmailConfig(gmail)
	configPayment.SetupGlobalMidtransConfig()
	transactionUsecase = transaction.NewUsecase(&transactionRepoMock, &campaignMockRepo, &paymentRepositoryMock, *dialer, &rewardMockRepo, &mockUserRepo)
	transactionDummy = transaction.Domain{
		ID:                1,
		CampaignID:        1,
		UserID:            1,
		PaymentLink:       "",
		Nominal:           10000,
		Status:            "",
		TransactionStatus: "",
		FraudStatus:       "",
		PaymentType:       "",
	}
	sampleSnap = snap.Response{
		Token:         "123",
		RedirectURL:   "www.tahubulat.com",
		StatusCode:    "200",
		ErrorMessages: []string{},
	}
	domainNotificationDummy = transaction.DomainNotification{
		TransactionStatus: "settlement",
		OrderID:           "1",
		PaymentType:       "bank_transfer",
		FraudStatus:       "accept",
	}
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
}

func TestCreateTransaction(t *testing.T) {
	setup()
	t.Run("Success Create Transaction", func(t *testing.T) {
		paymentRepositoryMock.On("GetLinkResponse",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&sampleSnap).Once()
		transactionRepoMock.On("CreateTransaction",
			mock.AnythingOfType("*transaction.Domain"),
			).Return(transactionDummy,nil).Once()
		resp,err := transactionUsecase.CreateTransaction(1,1,10000)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Failed Create Transaction", func(t *testing.T) {
		paymentRepositoryMock.On("GetLinkResponse",
			mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&sampleSnap).Once()
		transactionRepoMock.On("CreateTransaction",
			mock.AnythingOfType("*transaction.Domain"),
		).Return(transaction.Domain{},errors.New("Gagal Membuat Transaksi")).Once()
		resp,err := transactionUsecase.CreateTransaction(1,1,10000)
		assert.Error(t, err)
		assert.Equal(t, transaction.Domain{},resp)
	})
}

func TestGetStatusByID(t *testing.T) {
	setup()
	t.Run("Successfully get data", func(t *testing.T) {
		transactionRepoMock.On("GetByID",
			mock.AnythingOfType("int"),
			).Return(transactionDummy,nil).Once()
		resp,err := transactionUsecase.GetStatusByID(1)
		assert.Nil(t, err)
		assert.Equal(t, transactionDummy.ID,resp.ID)
	})

	t.Run("Failed get data | error Internal", func(t *testing.T) {
		transactionRepoMock.On("GetByID",
			mock.AnythingOfType("int"),
		).Return(transaction.Domain{},errors.New("Internal Error")).Once()
		resp,err := transactionUsecase.GetStatusByID(1)
		assert.Error(t, err)
		assert.Equal(t, transaction.Domain{},resp)
	})

	t.Run("Failed get data | data not found", func(t *testing.T) {
		transactionRepoMock.On("GetByID",
			mock.AnythingOfType("int"),
		).Return(transaction.Domain{},nil).Once()
		resp,err := transactionUsecase.GetStatusByID(1)
		assert.Error(t, err)
		assert.Equal(t, transaction.Domain{},resp)
	})
}

func TestGetNotificationPayment(t *testing.T) {
	setup()
	t.Run("Success Test 1 | Success do notification", func(t *testing.T) {
		transactionRepoMock.On("GetByID",mock.AnythingOfType("int")).
			Return(transactionDummy,nil).Once()
		transactionDummy.Status = "paid"
		transactionDummy.PaymentType = domainNotificationDummy.PaymentType
		transactionDummy.TransactionStatus = domainNotificationDummy.TransactionStatus
		transactionDummy.FraudStatus = domainNotificationDummy.FraudStatus
		transactionRepoMock.On("UpdateTransaction",mock.AnythingOfType("*transaction.Domain")).
			Return(&transactionDummy,nil).Once()
		campaignMockRepo.On("FindByID",mock.AnythingOfType("int")).
			Return(dataCampaign,nil).Once()
		dataCampaign.Supporters = dataCampaign.Supporters + 1
		dataCampaign.AmountNow = dataCampaign.AmountNow + transactionDummy.Nominal

		rewardMockRepo.On("GetRewardByAmount",mock.AnythingOfType("int")).
			Return(1,"Tahoo Bolaaat Digoreeeng",nil).Once()

		mockUserRepo.On("GetEmailByID",mock.AnythingOfType("int")).
			Return("kafka@mail.com",nil).Once()

		campaignMockRepo.On("UpdateCampaign",mock.AnythingOfType("campaign.Domain")).
			Return(dataCampaign,nil).Once()

		var domainHistory reward.DomainHistory

		domainHistory.UserID = transactionDummy.UserID
		domainHistory.TransactionID = transactionDummy.ID
		domainHistory.RewardID = uint(1)

		rewardMockRepo.On("SaveRewardHistory",mock.AnythingOfType("reward.DomainHistory")).
			Return(nil).Once()

		resp,err := transactionUsecase.GetNotificationPayment(domainNotificationDummy)
		assert.Nil(t, err)
		assert.Equal(t, transactionDummy,resp)
	})

	t.Run("Fail Test 1 | Fail Convert", func(t *testing.T) {
		domainNotificationDummy.OrderID = "Haha"
		resp,err := transactionUsecase.GetNotificationPayment(domainNotificationDummy)
		assert.Error(t, err)
		assert.Equal(t, transaction.Domain{},resp)
	})
	t.Run("Fail Test 2 | Transaction Not Found", func(t *testing.T) {
		transactionRepoMock.On("GetByID",mock.AnythingOfType("int")).
			Return(transaction.Domain{},errors.New("Error When Processing TransactionDB/ not found")).Once()
		resp,err := transactionUsecase.GetNotificationPayment(domainNotificationDummy)
		assert.Error(t, err)
		assert.Equal(t, transaction.Domain{},resp)
	})

	//t.Run("Fail Test 3 | Fail to Update Transaction", func(t *testing.T) {
	//	transactionRepoMock.On("GetByID",mock.AnythingOfType("int")).
	//		Return(transactionDummy,nil).Once()
	//	transactionDummy.Status = "paid"
	//	transactionDummy.PaymentType = domainNotificationDummy.PaymentType
	//	transactionDummy.TransactionStatus = domainNotificationDummy.TransactionStatus
	//	transactionDummy.FraudStatus = domainNotificationDummy.FraudStatus
	//	transactionRepoMock.On("UpdateTransaction",mock.AnythingOfType("*transaction.Domain")).
	//		Return(&transactionDummy,nil).Once()
	//	resp,err := transactionUsecase.GetNotificationPayment(domainNotificationDummy)
	//	assert.Error(t, err)
	//	assert.Equal(t, transaction.Domain{},resp)
	//})

}