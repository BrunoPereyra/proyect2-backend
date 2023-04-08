package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID            string               `bson:"_id,omitempty"`
	Creator       primitive.ObjectID   `bson:"creator"`
	Description   string               `bson:"description"`
	Name          string               `bson:"name"`
	Prize         string               `bson:"prize"`
	Entry         float64              `bson:"entry"`
	Requirements  string               `bson:"requirements"`
	Applicants    []primitive.ObjectID `bson:"applicants"`
	Participants  []primitive.ObjectID `bson:"participants"`
	CurrentStatus string               `bson:"current_status"`
	CreatedAt     time.Time            `bson:"created_at"`
	UpdatedAt     time.Time            `bson:"updated_at"`
}
