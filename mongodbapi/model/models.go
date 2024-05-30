package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `josn:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie",omitempty`
	Watched bool               `json:"watched,omitempty"`
}
