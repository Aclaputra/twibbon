package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const uri = "mongodb://user:pass@sample.host:27017/?timeoutMS=5000"
const uri = "mongodb://root:password@admin.localhost:27017/?timeoutMS=5000"

type User struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string
}

type MongoDBConnect struct{}

func NewMongoDBConnect() *MongoDBConnect {
	return &MongoDBConnect{}
}

func (m *MongoDBConnect) getClient() (client *mongo.Client, err error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return
}

func main() {
	mongoDB := NewMongoDBConnect()
	client, err := mongoDB.getClient()
	if err != nil {
		fmt.Println("not connected to db")
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admiss").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	coll := client.Database("test").Collection("user")
	filter := bson.D{{"name", "Muhammad Acla"}}

	var userResult User
	err = coll.FindOne(context.TODO(), filter).Decode(&userResult)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("error no found", filter)
			return
		}
		panic(err)
	}

	fmt.Println(userResult)
}
