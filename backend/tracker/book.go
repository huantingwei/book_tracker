package tracker

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book
func listBook(query map[string]string) (books []Book, err error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(bookCol)

	listAll := true
	for _, v := range query {
		if v != "" {
			listAll = false
		}
	}
	if listAll == true {
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Printf("Could not get all books.\nError: %v", err)
			return books, err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var book Book
			if err = cursor.Decode(&book); err != nil {
				log.Printf("Could not decode this book.\nError: %v", err)
				return books, err
			} else {
				books = append(books, book)
			}
		}
		return books, nil
	} else {
		var filter bson.D
		for k, v := range query {
			if v != "" {
				fmt.Println(v)
				if err != nil {
					log.Printf("Could not what?.\nError: %v", err)
				}
				filter = append(filter, bson.E{Key: k, Value: v})
			}
		}
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			fmt.Println("error 1")
			log.Printf("Could not get books with filter %v.\nError: %v", filter, err)
			return books, err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var book Book
			if err = cursor.Decode(&book); err != nil {
				log.Printf("Could not decode this book.\nError: %v", err)
				return books, err
			} else {
				books = append(books, book)
			}
		}
		return books, nil
	}
}

func getBook(bookID primitive.ObjectID) (book Book, err error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(bookCol)

	res := collection.FindOne(ctx, bson.M{"id": bookID})
	res.Decode(&book)

	return book, nil
}

func addBook(book *Book) (primitive.ObjectID, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	book.ID = primitive.NewObjectID()

	collection := client.Database(db).Collection(bookCol)

	_, err := collection.InsertOne(ctx, book)
	if err != nil {
		log.Printf("Could not create Book: %v", err)
		return primitive.NilObjectID, err
	}

	return book.ID, nil
}

func deleteBook(bookID primitive.ObjectID) (int, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	noteCollection := client.Database(db).Collection(noteCol)
	bookCollection := client.Database(db).Collection(bookCol)

	// delete notes where note.bookID = bookID
	if _, err := noteCollection.DeleteMany(ctx, bson.D{{Key: "bookid", Value: bookID}}); err != nil {
		log.Printf("Could not remove book %v's notes.\nError: %v", bookID, err)
		return -1, err
	}

	// delete Book
	res, err := bookCollection.DeleteOne(ctx, bson.M{"id": bookID})
	if err != nil {
		log.Printf("Could not delete book %v.\nError: %v", bookID, err)
		return int(res.DeletedCount), err
	}

	return int(res.DeletedCount), nil
}

func editBook(fields map[string]interface{}) (int, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(bookCol)

	var updateFields bson.D
	for k, v := range fields {
		if v != "" {
			updateFields = append(updateFields, bson.E{Key: k, Value: v})
		}
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": fields["id"]},
		bson.D{
			{Key: "$set", Value: updateFields},
		},
	)
	if err != nil {
		log.Printf("Could not edit book %v.\nError: %v", fields["id"], err)
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

// Note
