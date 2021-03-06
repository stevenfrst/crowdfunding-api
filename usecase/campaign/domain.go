package campaign

import (
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	ID        uint `gorm:"primarykey"`
	UserID uint
	CampaignName string
	ShortDescription string
	LongDescription string
	Target int
	AmountNow int
	Supporters int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

type Users struct {
	ID        uint `gorm:"primarykey"`
	FullName string
	Email string
	Password string
	Job    string
	RoleID uint
	Campaigns interface{}
	Token string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserCampaign struct {
	ID        uint `gorm:"primarykey"`
	FullName string
	Email string
	Password string
	Job    string
	RoleID uint
	Campaigns interface{}
	Token string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}



type CampaignUsecaseInterface interface {
	RegisterUseCase(campaign *Domain) (Domain,error)
	GetByIDUseCase(id int) (Domain,error)
	GetAllCampaignDetail() ([]Users,error)
	ListAllCampaignByUserUseCase(id int) (UserCampaign, error)
	EditTargetCampaign(id,target int) (Domain,error)
}

type CampaignRepoInterface interface {
	CreateCampaign(campaignIn *Domain) (Domain,error)
	//GetLast() int
	FindOneCampaignByID(id int) (Domain,error)
	FindByID(ID int) (Domain,error)
	ListCampaignsByUserID(id int) (UserCampaign,error)
	UpdateCampaign(campaign Domain) (Domain,error)
	ListAllCampaignByUser() (users []Users)
	EditTargetCampaign(id,target int) (Domain,error)
}