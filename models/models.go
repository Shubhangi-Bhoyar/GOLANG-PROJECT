package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type List struct {
	Index      primitive.ObjectID `json:"_id.omitempty"`
	Task       string             `json:"task"`
	Completion bool               `json:"completion"`
}
