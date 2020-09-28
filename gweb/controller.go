package gweb

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const libraryDB = "library"
const bookCol = "book"

func AddOneBook(book *Book) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	book.ID = primitive.NewObjectID()

	collection := client.Database(libraryDB).Collection(bookCol)

	result, err := collection.InsertOne(ctx, book)
	if err != nil {
		log.Printf("Could not create Book: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func AddMultiBook(books *[]Book) ([]interface{}, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(libraryDB).Collection(bookCol)

	var bs []interface{}
	for _, book := range *books {
		book.ID = primitive.NewObjectID()
		// InsertMany takes in parameter type = interface{}
		// convert Book{} to interface{}
		bs = append(bs, book)
	}

	result, err := collection.InsertMany(ctx, bs)

	if err != nil {
		log.Printf("Could not create Book: %v", err)
		// return primitive.NilObjectID, err
	}

	oids := result.InsertedIDs
	return oids, nil
}

// func ListBook() (books []Book) {
func ListBook() {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(libraryDB).Collection(bookCol)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book bson.M
		if err = cursor.Decode(&book); err != nil {
			log.Fatal(err)
		}
		// bsonB, _ := bson.Marshal(book)
		// books = append(books, bson.Unmarshal(bsonB, &Book))
		fmt.Println(book)
	}
	return books

}
