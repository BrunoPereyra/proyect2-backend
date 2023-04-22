package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Status    string               `json:"status"`
	PostImage string               `json:"PostImage,omitempty"`
	TimeStamp time.Time            `json:"TimeStamp"`
	UserID    primitive.ObjectID   `json:"UserID"`
	Likes     []primitive.ObjectID `json:"Likes"`
}
