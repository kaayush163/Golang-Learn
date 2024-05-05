package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //when save this automatically take github mongo in import
	Movie string             `json:"movie,omitempty"`

	Watched bool `json:"watched,omitempty"`
}
