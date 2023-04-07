package validator

import "github.com/go-playground/validator"

type LoginValidatorStruct struct {
	NameUser string `json:"NameUser" validate:"required,max=70"`
	Password string `json:"password" validate:"required,min=8"`
}

func (L *LoginValidatorStruct) LoginValidator() error {
	validate := validator.New()
	return validate.Struct(L)
}
