package	response

import "github.com/stevenfrst/crowdfunding-api/usecase/users"

type UserResponse struct {
	Id int `json:"id"`
	FullName string `json:"fullname" validate:"required"`
	Job string `json:"job"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	RoleID int `json:"role"`
	Token string `json:"token"`
}

type UserResponseWCampaign struct {
	Id int `json:"id"`
	FullName string `json:"fullname"`
	Job string `json:"job"`
	Email string `json:"email"`
	Password string `json:"password"`
	RoleID int `json:"role"`
	Campaign interface{} `json:"campaign"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse {
		Id:       int(domain.ID),
		FullName: domain.FullName,
		Job:      domain.Job,
		Email:    domain.Email,
		Password: domain.Password,
		RoleID: int(domain.RoleID),
		Token: domain.Token,
	}
}

func FromDomainList(domain []users.Domain) (response []UserResponse) {
	for _,user := range domain {
		newUser := UserResponse{
			Id:       int(user.ID),
			FullName: user.FullName,
			Job:      user.Job,
			Email:    user.Email,
			Password: user.Password,
			RoleID: int(user.RoleID),
			Token: user.Token,
		}
		response = append(response,newUser)
	}
	return response
}
