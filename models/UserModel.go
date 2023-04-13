package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Avatar       string             `json:"avatar"`
	FullName     string             `json:"FullName"`
	NameUser     string             `json:"NameUser"`
	PasswordHash string             `json:"passwordHash"`
	Pais         string             `json:"Pais"`
	Ciudad       string             `json:"Ciudad"`
	Email        string             `json:"Email"`
	Instagram    string             `json:"instagram,omitempty"`
	Twitter      string             `json:"twitter,omitempty"`
	Youtube      string             `json:"youtube,omitempty"`
}
