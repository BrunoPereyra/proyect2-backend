package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Championships struct {
	ID            string               `bson:"_id,omitempty"`
	Creator       primitive.ObjectID   `json:"creator"`
	Description   string               `json:"description"`
	Name          string               `json:"name"`
	Prize         string               `json:"prize"`
	Entry         float64              `json:"entry"`
	Requirements  string               `json:"requirements"`
	Applicants    []primitive.ObjectID `json:"applicants"`
	Participants  []primitive.ObjectID `json:"participants"`
	CurrentStatus string               `json:"current_status"`
	CreatedAt     time.Time            `json:"created_at"`
	UpdatedAt     time.Time            `json:"updated_at"`
}
