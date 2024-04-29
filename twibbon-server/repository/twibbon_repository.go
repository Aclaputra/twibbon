package repository

import (
	"context"
	"fmt"
	"twibbon-server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TwibbonRepository struct {
	coll *mongo.Collection
}

type twibbonRepository interface {
	ReadAll() (twibbons []models.Twibbon, err error)
	Create(twibbon models.Twibbon) (err error)
	Delete(id string) (err error)
}

func NewTwibbonRepository(database *mongo.Database) *TwibbonRepository {
	names, err := database.ListCollectionNames(
		context.TODO(),
		bson.D{{"options.capped", true}},
	)

	if err != nil {
		fmt.Printf("failed to get coll names: %v\n", err)
	}

	var found bool
	for _, name := range names {
		if name == "twibbon" {
			found = true
			break
		}
	}

	if !found {
		err := database.CreateCollection(context.TODO(), "twibbon")
		if err != nil {
			fmt.Println(err)
		}
	}

	return &TwibbonRepository{
		coll: database.Collection("twibbon"),
	}
}

func (r *TwibbonRepository) ReadAll() (twibbons []models.Twibbon, err error) {

	return
}

func (r *TwibbonRepository) Create(twibbon models.Twibbon) (err error) {

	return
}

func (r *TwibbonRepository) Delete(id string) (err error) {

	return
}
