package validator

import "github.com/go-playground/validator"

type UserModelValidator struct {
	FullName  string `json:"fullName" validate:"required,max=70"`
	NameUser  string `json:"NameUser" validate:"required,max=70"`
	Password  string `json:"password" validate:"required,min=8"`
	Pais      string `json:"Pais" validate:"required"`
	Ciudad    string `json:"Ciudad" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Instagram string `json:"instagram" default:""`
	Twitter   string `json:"twitter" default:""`
	Youtube   string `json:"youtube" default:""`
}

func (u *UserModelValidator) ValidateUserFind() error {
	validate := validator.New()
	return validate.Struct(u)
}
