package helpers

import (
	"backend/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserTMiddlExist(dataMiddleware string, db *mongo.Database) (models.UserModel, error) {
	GoMongoDBCollUsers := db.Collection("users")
	find := bson.D{
		{Key: "nameuser", Value: dataMiddleware},
	}
	var UserCreator models.UserModel
	err := GoMongoDBCollUsers.FindOne(context.TODO(), find).Decode(&UserCreator)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			return UserCreator, errors.New("ErrNoDocuments")
		} else {
			return UserCreator, errors.New("ErrServer")
		}
	}
	return UserCreator, nil
}
