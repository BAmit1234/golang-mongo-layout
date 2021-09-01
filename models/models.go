package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Book   *bookName          `json:"title" bson:"title,omitempty"`
	Author string             `json:"author" bson:"author,omitempty"`
}

type bookName struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Cost string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
