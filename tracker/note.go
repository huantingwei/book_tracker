package tracker

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	noteCol = "note"
)

func listNote(query map[string]string) (notes []Note, err error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(noteCol)

	listAll := true
	for _, v := range query {
		if v != "" {
			listAll = false
		}
	}
	if listAll == true {
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
			return notes, err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var note Note
			if err = cursor.Decode(&note); err != nil {
				log.Fatal(err)
				return notes, err
			} else {
				notes = append(notes, note)
			}
		}
		return notes, nil
	} else {
		var filter bson.D
		for k, v := range query {
			if v != "" {
				fmt.Println(v)
				if err != nil {
					log.Fatal(err)
				}
				filter = append(filter, bson.E{k, v})
			}
		}
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Fatal(err)
			return notes, err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var note Note
			if err = cursor.Decode(&note); err != nil {
				log.Fatal(err)
				return notes, err
			} else {
				notes = append(notes, note)
			}
		}
		return notes, nil
	}
}

func listNoteByBook(bookID primitive.ObjectID) (notes []Note, err error) {
	book, err := getBook(bookID)
	if err != nil {
		log.Fatal(err)
		return notes, err
	}

	noteIDs := book.Notes

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(noteCol)

	for _, noteID := range noteIDs {
		var note Note
		res := collection.FindOne(ctx, bson.M{"id": noteID})
		res.Decode(&note)
		notes = append(notes, note)
	}

	return notes, nil

}

func getNote(noteID primitive.ObjectID) (note Note, err error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(noteCol)

	res := collection.FindOne(ctx, bson.M{"id": noteID})
	res.Decode(&note)

	return note, nil
}

func addNote(bookID primitive.ObjectID, note *Note) (primitive.ObjectID, error) {

	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	note.ID = primitive.NewObjectID()

	collection := client.Database(db).Collection(noteCol)

	// insert a new note
	res, err := collection.InsertOne(ctx, note)
	if err != nil {
		log.Printf("Could not create Note: %v", err)
		return primitive.NilObjectID, err
	}
	noteID := res.InsertedID.(primitive.ObjectID)

	// get the book's old note array
	book, err := getBook(note.BookID)
	if err != nil {
		log.Printf("Could not create Note: %v", err)
		return primitive.NilObjectID, err
	}

	oldNotes := book.Notes
	oldNotes = append(oldNotes, noteID)

	fields := make(map[string]interface{})
	fields["id"] = bookID
	fields["notes"] = oldNotes

	// append new note id to the book's note array
	_, err = editBook(fields)
	if err != nil {
		log.Printf("Could not link the note to the Book: %v", err)
		_, _ = deleteNote(noteID)
		return primitive.NilObjectID, err
	}

	return noteID, nil

}

func deleteNote(id primitive.ObjectID) (int, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(noteCol)

	res, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
		return int(res.DeletedCount), err
	}
	return int(res.DeletedCount), nil
}
