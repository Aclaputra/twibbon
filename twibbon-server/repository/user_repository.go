package repository

import (
	"context"
	"fmt"
	"twibbon-server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	coll *mongo.Collection
}

type userRepository interface {
	ReadUserByName(name string) (result models.User, err error)
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	names, err := database.ListCollectionNames(
		context.TODO(),
		bson.D{{"options.capped", true}},
	)
	if err != nil {
		fmt.Printf("failed to get coll names: %v\n", err)
	}

	var found bool
	for _, name := range names {
		if name == "user" {
			found = true
			break
		}
	}

	if !found {
		err := database.CreateCollection(context.TODO(), "user")
		if err != nil {
			fmt.Println(err)
		}
	}

	return &UserRepository{
		coll: database.Collection("user"),
	}
}

func (u *UserRepository) ReadUserByName(name string) (result models.User, err error) {
	filter := bson.D{{"name", name}}
	err = u.coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("error not found", filter)
			return
		}
		fmt.Println(err)
	}

	return
}

// func (u *UserRepository) CreateUser(user models.User) (result models.User, err error) {
//
// 	return
// }
