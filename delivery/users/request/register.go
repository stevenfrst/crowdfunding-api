package request

import "github.com/stevenfrst/crowdfunding-api/usecase/users"

type UserRegister struct {
	Id int `json:"id"`
	FullName string `json:"fullname" validate:"required"`
	Job string `json:"job"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	RoleID int `json:"role"`
}

type UserLogin struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		ID:       uint(u.Id),
		FullName: u.FullName,
		Email:    u.Email,
		Password: u.Password,
		Job:      u.Job,
		RoleID: uint(u.RoleID),
	}
}