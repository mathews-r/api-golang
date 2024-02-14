package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type CategoryEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Category string             `bson:"category,omitempty"`
}
