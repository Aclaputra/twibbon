package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:password@admin.localhost:27017/?timeoutMS=5000"

type User struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string
}

type MongoDBConnect struct{}

type mongoDBConnect interface {
	GetClient() (client *mongo.Client, err error)
	Ping(client *mongo.Client)
}

func NewMongoDBConnect() *MongoDBConnect {
	return &MongoDBConnect{}
}

func (m *MongoDBConnect) GetClient() (client *mongo.Client, err error) {

	defer fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return
}

func (m *MongoDBConnect) Ping(client *mongo.Client) {
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admiss").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
}

func ReadUsers(client *mongo.Client) {
	coll := client.Database("test").Collection("user")
	filter := bson.D{{"name", "Muhammad Acla"}}

	var userResult User
	err := coll.FindOne(context.TODO(), filter).Decode(&userResult)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("error no found", filter)
			return
		}
		panic(err)
	}

	fmt.Println(userResult)

}

func main() {
	mongoDB := NewMongoDBConnect()
	client, err := mongoDB.GetClient()
	if err != nil {
		fmt.Println("not connected to db")
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mongoDB.Ping(client)
	ReadUsers(client)
}
