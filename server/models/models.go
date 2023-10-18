package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//struct is our own datatype created in golang
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status, omitempty"`
}
