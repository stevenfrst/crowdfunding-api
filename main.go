package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config2 "github.com/stevenfrst/crowdfunding-api/app/config"
	_middleware "github.com/stevenfrst/crowdfunding-api/app/middleware"
	routes "github.com/stevenfrst/crowdfunding-api/app/routes"
	_campaignDelivery "github.com/stevenfrst/crowdfunding-api/delivery/campaign"
	_rewardDelivery "github.com/stevenfrst/crowdfunding-api/delivery/reward"
	_transactionDelivery "github.com/stevenfrst/crowdfunding-api/delivery/transaction"
	userDelivery "github.com/stevenfrst/crowdfunding-api/delivery/users"
	"github.com/stevenfrst/crowdfunding-api/drivers/email"
	payment "github.com/stevenfrst/crowdfunding-api/drivers/midtrans"
	"github.com/stevenfrst/crowdfunding-api/drivers/mysql"
	repoModels "github.com/stevenfrst/crowdfunding-api/drivers/repository"
	_campaignRepo "github.com/stevenfrst/crowdfunding-api/drivers/repository/campaign"
	_rewardRepo "github.com/stevenfrst/crowdfunding-api/drivers/repository/reward"
	_transactionRepo "github.com/stevenfrst/crowdfunding-api/drivers/repository/transaction"
	_userRepo "github.com/stevenfrst/crowdfunding-api/drivers/repository/users"
	_campaignUseCase "github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	_rewardUseCase "github.com/stevenfrst/crowdfunding-api/usecase/reward"
	_transactionUseCase "github.com/stevenfrst/crowdfunding-api/usecase/transaction"
	_userUsecase "github.com/stevenfrst/crowdfunding-api/usecase/users"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)



// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

type CustomValidator struct {
	Validator *validator.Validate
}


func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func dbMigrate(db *gorm.DB) {


	//err := db.AutoMigrate(&_userRepo.User{},&_campaignRepo.Campaign{},&_transactionRepo.Transaction{})
	//err := db.AutoMigrate(&_campaignRepo.Campaign{},&_userRepo.User{},&_transactionRepo.Transaction{})
	err := db.AutoMigrate(&repoModels.User{},&repoModels.Campaign{},repoModels.Transaction{},repoModels.RewardHistory{},repoModels.Reward{})
	var users = []repoModels.User{{ID:1,FullName: "admin",Email: "mail@admin.com",Password: "password",Job: "Administrator",RoleID: 1},
		{ID:2,FullName: "kafka",Email: "kafka@user.com",Password: "password",Job: "Serabutan",RoleID: 2},
		{ID:3,FullName: "ponta",Email: "ponta@user.com",Password: "password",Job: "kolee",RoleID: 2},
	}
	db.Create(users)

	var rewards = []repoModels.Reward{{ID: 1,Amount:10000 ,RewardDescription: "Kuota Zoom 6 Jam"},
		{ID: 2,Amount:50000 ,RewardDescription: "Voucher Mie Fajar"},
		{ID: 3,Amount:100000 ,RewardDescription: "Amidis 1 Galon"},
	}
	db.Create(&rewards)
	if err != nil {
		log.Println("Failed to migrate")
	}
}

func main() {
	//fmt.Println("Hello")
	config := config2.GetConfig()

	configPayment := payment.ConfigMidtrans{
		SERVER_KEY: config.SERVER_KEY,
	}

	gmail := email.GmailConfig{
		CONFIG_SMTP_HOST:       config.CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT:       config.CONFIG_SMTP_PORT,
		CONFIG_SMTP_AUTH_EMAIL: config.CONFIG_SMTP_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD:   config.CONFIG_AUTH_PASSWORD,
		CONFIG_SENDER_NAME: config.CONFIG_SENDER_NAME,
	}
	dialer := email.NewGmailConfig(gmail)

	configPayment.SetupGlobalMidtransConfig()
	payment.InitializeSnapClient()

	configdb := mysql.ConfigDB{
		DB_Username: config.DB_USERNAME,
		DB_Password: config.DB_PASSWORD,
		DB_Host:     config.DB_HOST,
		DB_Port:     config.DB_PORT,
		DB_Database: config.DB_NAME,
	}
	db := configdb.InitialDb()
	dbMigrate(db)

	jwt := _middleware.ConfigJWT{
		SecretJWT:       config.JWT_SECRET,
		ExpiresDuration: config.JWT_EXPIRED,
	}

	timeoutContext := time.Duration(config.CONTEXT_TIMEOUT) * time.Second

	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	c := jaegertracing.New(e, nil)
	defer c.Close()
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())



	// User

	userRepoInterface := _userRepo.NewUserRepository(db)
	userUseCaseInterface := _userUsecase.NewUsecase(userRepoInterface,&jwt, timeoutContext)
	userDeliveryInterface := userDelivery.NewUserDelivery(userUseCaseInterface)

	// Campaign
	CampaignRepoInterface := _campaignRepo.NewCampaignRepository(db)
	campaignUseCaseInterface := _campaignUseCase.NewCampaignUseCase(CampaignRepoInterface)
	campaignDeliveryInterface := _campaignDelivery.NewCampaignDelivery(campaignUseCaseInterface)

	RewardRepoInterface := _rewardRepo.NewRewardRepository(db)
	RewardUseCaseInterface := _rewardUseCase.NewUsecase(RewardRepoInterface)
	RewardDeliveryInterface := _rewardDelivery.NewRewardDelivery(RewardUseCaseInterface)

	TransactionRepoInterface := _transactionRepo.NewTransactionRepository(db)
	transactionUseCaseInterface := _transactionUseCase.NewUsecase(TransactionRepoInterface,CampaignRepoInterface,configPayment,*dialer,RewardRepoInterface,userRepoInterface)
	transactionDeliveryInterface := _transactionDelivery.NewTransactionDelivery(transactionUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserDelivery: *userDeliveryInterface,
		CampaignDelivery: *campaignDeliveryInterface,
		TransactionDelivery: *transactionDeliveryInterface,
		RewardDelivery: *RewardDeliveryInterface,
		JWTConfig:      jwt.Init(),
	}

	routesInit.RouteRegister(e)
	e.Logger.Fatal(e.Start(":1234"))


}