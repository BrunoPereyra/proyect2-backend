package helpers

import (
	"backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserTMiddlExist(dataMiddleware string, db *mongo.Database, UserCreator chan models.User, errChanelUserTMiddlExist chan error) {
	GoMongoDBCollUsers := db.Collection("users")

	find := bson.D{
		{Key: "nameuser", Value: dataMiddleware},
	}
	var UserCreatorMddl models.User
	err := GoMongoDBCollUsers.FindOne(context.Background(), find).Decode(&UserCreatorMddl)

	if err != nil {
		errChanelUserTMiddlExist <- err
	} else {
		UserCreator <- UserCreatorMddl
	}
}
