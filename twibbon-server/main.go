package main

import (
	"context"
	"fmt"
	"twibbon-server/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:password@admin.localhost:27017/?timeoutMS=5000"

type MongoDB struct{}

type mongoDB interface {
	GetClient() (client *mongo.Client, err error)
	Ping(client *mongo.Client)
}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (m *MongoDB) GetClient() (client *mongo.Client, err error) {

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

func (m *MongoDB) ConnectTwibbon(ctx context.Context, client *mongo.Client) (db *mongo.Database, err error) {
	return client.Database("twibbon_db"), nil
}

func (m *MongoDB) Ping(client *mongo.Client) {
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admiss").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
}

func main() {
	mongoDB := NewMongoDB()
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

	twibbon_db, err := mongoDB.ConnectTwibbon(context.TODO(), client)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(twibbon_db)
	user, err := userRepository.ReadUserByName("Muhammad Acla")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
