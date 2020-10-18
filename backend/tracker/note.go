package tracker

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			log.Printf("Could not get all notes.\nError: %v", err)
			return notes, err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var note Note
			if err = cursor.Decode(&note); err != nil {
				log.Printf("Could not decode this note.\nError: %v", err)
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
					log.Printf("Could not what?.\nError: %v", err)
				}
				filter = append(filter, bson.E{Key: k, Value: v})
			}
		}
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			log.Printf("Could not get notes with filter %v.\nError: %v", filter, err)
			return notes, err
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var note Note
			if err = cursor.Decode(&note); err != nil {
				log.Printf("Could not decode this note.\nError: %v", err)
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
		log.Printf("Could not get book %v.\nError: %v", bookID, err)
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

	noteCollection := client.Database(db).Collection(noteCol)
	bookCollection := client.Database(db).Collection(bookCol)

	note.ID = primitive.NewObjectID()

	// append the noteID to the Book's notes
	// also checking if bookID exists
	if _, err := bookCollection.UpdateOne(ctx,
		bson.M{"id": bookID},
		bson.D{{Key: "$push", Value: bson.D{{Key: "notes", Value: note.ID}}}}); err != nil {
		log.Printf("Could not append note to book: %v", err)
		return primitive.NilObjectID, err
	}

	// insert note
	if _, err := noteCollection.InsertOne(ctx, note); err != nil {
		log.Printf("Could not create Note: %v", err)
		return primitive.NilObjectID, err
	}
	return note.ID, nil
}

func deleteNote(noteID primitive.ObjectID) (int, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	noteCollection := client.Database(db).Collection(noteCol)
	bookCollection := client.Database(db).Collection(bookCol)

	var note Note
	// get note, to get bookID
	res := noteCollection.FindOne(ctx, bson.M{"id": noteID})
	res.Decode(&note)
	bookID := note.BookID

	// delete note from Book.notes
	if _, err := bookCollection.UpdateOne(ctx,
		bson.M{"id": bookID},
		bson.D{{Key: "$pull", Value: bson.D{{Key: "notes", Value: noteID}}}}); err != nil {
		log.Printf("Could not delete note from book.\nError: %v", err)
		return -1, err
	}

	// delete note
	delNoteRes, err := noteCollection.DeleteOne(ctx, bson.M{"id": noteID})
	if err != nil {
		log.Printf("Could not delete Note.\nError: %v", err)
		return -1, err
	}
	return int(delNoteRes.DeletedCount), nil

}

func editNote(fields map[string]interface{}) (int, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	collection := client.Database(db).Collection(noteCol)

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
		log.Printf("Could not edit note %v.\nError: %v", fields["id"], err)
		return 0, err
	}
	return int(result.ModifiedCount), nil
}

// WithTransaction
/*
func addNote(bookID primitive.ObjectID, note *Note) (interface{}, error) {

	// check if bookID exists

	// all or zero (transaction):
	// - Book.notes append noteID
	// - insert Note
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	// Prereq: Create collections.
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)

	noteCollection := client.Database(db).Collection(noteCol, wcMajorityCollectionOpts)
	bookCollection := client.Database(db).Collection(bookCol, wcMajorityCollectionOpts)

	// Step 1: Define the callback that specifies the sequence of operations to perform inside the transaction.
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.

		// insert note
		note.ID = primitive.NewObjectID()
		if _, err := noteCollection.InsertOne(sessCtx, note); err != nil {
			log.Printf("Could not create Note: %v", err)
			return primitive.NilObjectID, err
		}
		// append the noteID to the Book's notes
		if _, err := bookCollection.UpdateOne(sessCtx,
			bson.M{"id": bookID},
			bson.D{{Key: "$push", Value: bson.D{{Key: "notes", Value: note.ID}}}}); err != nil {
			log.Printf("Could not append note to book: %v", err)
			return primitive.NilObjectID, err
		}
		return note.ID, nil

	}

	// Step 2: Start a session and run the callback using WithTransaction.
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	noteID, err := session.WithTransaction(ctx, callback)
	if err != nil {
		panic(err)
	}
	fmt.Printf("note ID: %v\n", noteID)

	return noteID, nil

}
*/
/*
func deleteNote(noteID primitive.ObjectID) (int, error) {
	client, ctx, cancel := getConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	// Prereq: Create collections.
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	wcMajorityCollectionOpts := options.Collection().SetWriteConcern(wcMajority)

	noteCollection := client.Database(db).Collection(noteCol, wcMajorityCollectionOpts)
	bookCollection := client.Database(db).Collection(bookCol, wcMajorityCollectionOpts)

	// Step 1: Define the callback that specifies the sequence of operations to perform inside the transaction.
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Important: You must pass sessCtx as the Context parameter to the operations for them to be executed in the
		// transaction.

		var note Note
		// get note, to get bookID
		res := noteCollection.FindOne(ctx, bson.M{"id": noteID})
		res.Decode(&note)
		bookID := note.BookID

		// delete note from Book.notes
		if _, err := bookCollection.UpdateOne(sessCtx,
			bson.M{"id": bookID},
			bson.D{{Key: "$pop", Value: bson.D{{Key: "notes", Value: noteID}}}}); err != nil {
			log.Printf("Could not delete note from book. Error: %v", err)
			return -1, err
		}

		// delete note
		delNoteRes, err := noteCollection.DeleteOne(ctx, bson.M{"id": noteID})
		if err != nil {
			log.Printf("Could not delete Note. Error: %v", err)
			return -1, err
		}

		return delNoteRes.DeletedCount, nil
	}

	// Step 2: Start a session and run the callback using WithTransaction.
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(ctx)

	delCount, err := session.WithTransaction(ctx, callback)
	if err != nil {
		panic(err)
	}
	fmt.Printf("deleted %v note(s)\n", delCount)

	if dc, ok := delCount.(int); ok {
		return dc, nil
	} else {
		return -1, nil
	}

}
*/
