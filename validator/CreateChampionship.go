package validator

import (
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChampionshipsValidate struct {
	Creator       primitive.ObjectID   `json:"creator" validate:"required"`
	Description   string               `json:"description" validate:"required,min=10,max=100"`
	Name          string               `json:"name" validate:"required,min=5,max=50"`
	Prize         string               `json:"prize" validate:"required,min=5,max=50"`
	Entry         float64              `json:"entry" validate:"required,gt=0"`
	Requirements  string               `json:"requirements" validate:"required,min=5,max=50"`
	Applicants    []primitive.ObjectID `json:"applicants"`
	Participants  []primitive.ObjectID `json:"participants"`
	CurrentStatus string               `json:"current_status"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

func (e *ChampionshipsValidate) ChampionshipsValidate() error {
	validate := validator.New()
	return validate.Struct(e)
}
