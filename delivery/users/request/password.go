package request

import "github.com/stevenfrst/crowdfunding-api/usecase/users"

type PasswordUpdate struct {
	Email string `json:"email" validate:"required,email"`
	OldPassword string `json:"old_password" validate:"min=8"`
	NewPassword string `json:"new_password" validate:"min=8"`
}

func (req PasswordUpdate) ToDomain() users.DomainUpdate {
	return users.DomainUpdate{
		Email:req.Email,
		OldPassword:req.OldPassword,
		NewPassword:req.NewPassword,
	}
}