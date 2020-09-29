package gweb

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	libraryDB = "library"
	bookCol   = "book"
)

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

func ListBook() (books []Book) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(libraryDB).Collection(bookCol)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// small dataset
	if err = cursor.All(ctx, &books); err != nil {
		log.Fatal(err)
	}

	// large dataset
	// for cursor.Next(ctx) {
	// 	var book Book
	// cursor.Decode will unmarshal the current document into the interface{} passed in (i.e. Book)
	// 	if err = cursor.Decode(&book); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	books = append(books, book)
	// }

	return books
}

func ListBookWithFilter(query map[string]string) (books []Book) {

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(libraryDB).Collection(bookCol)

	var filter bson.D
	for k, v := range query {
		if v != "" {
			fmt.Println(v)
			filter = append(filter, bson.E{k, v})
		}
	}
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &books); err != nil {
		log.Fatal(err)
	}
	return books
}

func GetBookByID(id string) (book Book) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(libraryDB).Collection(bookCol)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	res := collection.FindOne(ctx, bson.M{"id": oid})
	res.Decode(&book)

	return book

}

func BorrowBookByID(id string) error {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(libraryDB).Collection(bookCol)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		return err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": oid},
		bson.D{
			{"$set", bson.D{{"status", 0}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}

func ReturnBookByID(id string) error {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(libraryDB).Collection(bookCol)

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		return err
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"id": oid},
		bson.D{
			{"$set", bson.D{{"status", 1}}},
		},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}
