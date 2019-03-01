package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	Id   primitive.ObjectID `json:"ID" bson:"_id"`
	Name string             `json:"Name" bson:"filename"`
	Size int                `json:"size" bson:"length"`
}
