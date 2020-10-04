package tracker

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID   `json:"id"`
	Title       string               `json:"title"`
	Author      string               `json:"author"`
	Status      int                  `json:"status"`
	StartTime   time.Time            `json:"startTime"`
	EndTime     time.Time            `json:"endTime"`
	Notes       []primitive.ObjectID `json:"notes"`
	Description string               `json:"description"`
}

type Note struct {
	ID      primitive.ObjectID `json:"id"`
	BookID  primitive.ObjectID `json:"bookID"`
	Content string             `json:"content"`
	ReplyTo primitive.ObjectID `json:"replyTo"`
}
