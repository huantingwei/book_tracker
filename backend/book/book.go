package book

import (
	"context"

	m "github.com/huantingwei/go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getBook(bookID primitive.ObjectID, s *Service) (book m.Book) {

	// get book
	res := s.bookCollection.FindOne(context.TODO(), bson.M{"id": bookID})
	res.Decode(&book)

	// check if the book exist
	// if no, mongodb will return primitive.NilObjectID
	// if book.ID == primitive.NilObjectID {
	// 	return book, error
	// }

	return book
}
