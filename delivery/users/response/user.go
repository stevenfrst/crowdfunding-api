package	response

import "github.com/stevenfrst/crowdfunding-api/usecase/users"

type UserResponse struct {
	Id int `json:"id"`
	FullName string `json:"fullname" validate:"required"`
	Job string `json:"job"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	RoleID int `json:"role"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse {
		Id:       int(domain.ID),
		FullName: domain.FullName,
		Job:      domain.Job,
		Email:    domain.Email,
		Password: domain.Password,
		RoleID: int(domain.RoleID),
	}
}
