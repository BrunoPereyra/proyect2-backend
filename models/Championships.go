package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Championships struct {
	ID           string             `json:"id" bson:"_id,omitempty"`
	Creator      primitive.ObjectID `json:"creator"`
	Description  string             `json:"description"`
	Name         string             `json:"name"`
	Prize        string             `json:"prize"`
	Entry        float64            `json:"entry"`
	Requirements string             `json:"requirements"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	CurrentStatus string `json:"current_status"`

	Applicants         []primitive.ObjectID `json:"applicants"`
	AcceptedApplicants []primitive.ObjectID `json:"acceptedApplicants"`

	ParticipantsWhoPaidTheEntrance []primitive.ObjectID `json:"ParticipantsWhoPaidTheEntrance"`

	Votesoftheparticipants map[primitive.ObjectID][]primitive.ObjectID `json:"votesOfTheParticipants"`
	Voters                 []primitive.ObjectID                        `json:"voters"`
}
