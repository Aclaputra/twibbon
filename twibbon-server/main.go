package main

import (
	"context"
	"fmt"
	"twibbon-server/database"
	"twibbon-server/repository"
)

func main() {
	mongoDB := database.NewMongoDB()
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
	twibbonRepository := repository.NewTwibbonRepository(twibbon_db)
	user, err := userRepository.ReadByName("Muhammad Acla")
	if err != nil {
		fmt.Println("not found", err)
		// panic(err)
	}
	twibbons, err := twibbonRepository.ReadAll()

	fmt.Println(user)
	fmt.Println(twibbons)
}
