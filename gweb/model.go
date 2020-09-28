package gweb

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `json:"id"`
	Name   string             `json:"title"`
	Author string             `json:"author"`
	Status int                `json:"status"`
}

type User struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
}
