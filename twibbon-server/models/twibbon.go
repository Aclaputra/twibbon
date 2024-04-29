package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Twibbon struct {
	ID       primitive.ObjectID `bson:"_id"`
	FileName string
	FileSize string
}
