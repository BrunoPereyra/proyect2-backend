package validator

import (
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChampionshipsValidate struct {
	Creator       primitive.ObjectID `bson:"creator" validate:"required"`
	Description   string             `bson:"description" validate:"required,min=10,max=100"`
	Name          string             `bson:"name" validate:"required,min=5,max=50"`
	Prize         string             `bson:"prize" validate:"required,min=5,max=50"`
	Entry         float64            `bson:"entry" validate:"required,gt=0"`
	Requirements  string             `bson:"requirements" validate:"required,min=5,max=50"`
	Applicants    []string           `bson:"applicants"`
	Participants  []string           `bson:"participants"`
	CurrentStatus string             `bson:"current_status"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}

func (e *ChampionshipsValidate) ChampionshipsValidate() error {
	validate := validator.New()
	return validate.Struct(e)
}
