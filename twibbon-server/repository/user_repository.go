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
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	// err := database.CreateCollection(context.TODO(), "user")
	// if err != nil {
	// 	fmt.Println("cannot create db")
	// 	panic(err)
	// }
	return &UserRepository{
		coll: database.Collection("user"),
	}
}

func (u *UserRepository) ReadUserByName(name string) (result models.User, err error) {
	// coll := client.Database("test").Collection("user")
	filter := bson.D{{"name", name}}
	err = u.coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("error not found", filter)
			return
		}
		panic(err)
	}

	return
}

func (u *UserRepository) CreateUser(user models.User) (result models.User, err error) {

	return
}
