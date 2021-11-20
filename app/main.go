package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	config2 "github.com/stevenfrst/crowdfunding-api/app/config"
	_middleware "github.com/stevenfrst/crowdfunding-api/app/middleware"
	routes "github.com/stevenfrst/crowdfunding-api/app/routes"
	_campaignDelivery "github.com/stevenfrst/crowdfunding-api/delivery/campaign"
	userDelivery "github.com/stevenfrst/crowdfunding-api/delivery/users"
	"github.com/stevenfrst/crowdfunding-api/drivers/mysql"
	repoModels "github.com/stevenfrst/crowdfunding-api/drivers/repository"
	_campaignRepo "github.com/stevenfrst/crowdfunding-api/drivers/repository/campaign"
	_userRepo "github.com/stevenfrst/crowdfunding-api/drivers/repository/users"
	_campaignUseCase "github.com/stevenfrst/crowdfunding-api/usecase/campaign"
	_userUsecase "github.com/stevenfrst/crowdfunding-api/usecase/users"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

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
	err := db.AutoMigrate(&repoModels.User{},&repoModels.Campaign{},repoModels.Transaction{})
	var users = []repoModels.User{{ID:1,FullName: "admin",Email: "mail@admin.com",Password: "password",Job: "Administrator",RoleID: 1},
		{ID:2,FullName: "kafka",Email: "kafka@user.com",Password: "password",Job: "Serabutan",RoleID: 2},
		{ID:3,FullName: "ponta",Email: "ponta@user.com",Password: "password",Job: "kolee",RoleID: 2},
	}
	db.Create(users)
	if err != nil {
		log.Println("Failed to migrate")
	}
}

func main() {
	//fmt.Println("Hello")
	config := config2.GetConfig()
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

	// User

	userRepoInterface := _userRepo.NewUserRepository(db)
	userUseCaseInterface := _userUsecase.NewUsecase(userRepoInterface, timeoutContext)
	userDeliveryInterface := userDelivery.NewUserDelivery(userUseCaseInterface)

	// Campaign
	CampaignRepoInterface := _campaignRepo.NewCampaignRepository(db)
	campaignUseCaseInterface := _campaignUseCase.NewCampaignUseCase(CampaignRepoInterface)
	campaignDeliveryInterface := _campaignDelivery.NewCampaignDelivery(campaignUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserDelivery: *userDeliveryInterface,
		CampaignDelivery: *campaignDeliveryInterface,
		JWTConfig:      jwt.Init(),
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(":1234"))

}