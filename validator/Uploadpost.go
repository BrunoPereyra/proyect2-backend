package validator

import (
	"github.com/go-playground/validator"
)

type UploadPostValidate struct {
	Status string `bson:"Status" validate:"required,min=5"`
}

func (e *UploadPostValidate) UploadPostValidate() error {
	validate := validator.New()
	return validate.Struct(e)
}
