package book

import (
	"context"
	"log"

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
func deleteNote(bookID primitive.ObjectID, noteID primitive.ObjectID, s *Service) bool {

	// delete note from Book.notes
	if _, err := s.bookCollection.UpdateOne(context.TODO(),
		bson.M{"id": bookID},
		bson.D{{Key: "$pull", Value: bson.D{{Key: "notes", Value: noteID}}}}); err != nil {
		log.Printf("Could not delete note from book.\nError: %v", err)
		return false
	}
	return true
}

func strIdToPrimitive(id interface{}) (objID primitive.ObjectID, err error) {
	if str, ok := id.(string); ok {
		oid, err := primitive.ObjectIDFromHex(str)
		if err != nil {
			log.Println("Invalid id")
			return primitive.NilObjectID, err
		}
		return oid, nil
	} else {
		return primitive.NilObjectID, err
	}
}
