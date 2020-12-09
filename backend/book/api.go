package book

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	m "github.com/huantingwei/go/models"
	"github.com/huantingwei/go/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const layoutISO = "2006-01-02 15:04:05"

type Service struct {
	bookCollection *mongo.Collection
	noteCollection *mongo.Collection
}

func NewService(r *gin.RouterGroup, db util.Database) {
	s := &Service{
		bookCollection: db.Handle.Collection("book"),
	}

	r = r.Group("/book")

	r.GET("", s.listBook)
	r.GET("/:bookid", s.getBook)
	r.POST("", s.addBook)
	r.DELETE("", s.deleteBook)
	r.POST("/:bookid", s.editBook)

}

// listBook enumerate all books
// request: GET "/api/v1/book"
// response: [ {...BOOK_1}, {...BOOK_2}]
func (s *Service) listBook(c *gin.Context) {
	query := map[string]string{
		// "id":        c.Query("id"),
		"title":  c.Query("title"),
		"author": c.Query("author"),
		"status": c.Query("status"),
		// "startTime": c.Query("startTime"),
		// "endTime":   c.Query("endTime"),
	}

	// construct filter
	// bson.D{{"name", "hello"}} => find data with name == hello
	// check if any filter value exists with `listAll`
	var f bson.D
	listAll := true
	for k, v := range query {
		if v != "" {
			f = append(f, bson.E{Key: k, Value: v})
			listAll = false
		}
	}
	// convert type of filter
	var filter interface{}
	if listAll == true {
		// if no filter value, use bson.M{}
		filter = bson.M{}
	} else {
		// otherwise, use bson.D{bson.E{}}
		// e.g. bson.D{{"name", "hello"}}
		filter = f
	}

	fmt.Printf("filter: %s\n", filter)

	cursor, err := s.bookCollection.Find(context.TODO(), filter)
	if err != nil {
		log.Printf("Could not get books with filter %v.\nError: %v", filter, err)
		util.ResponseError(c, err)
		return
	}

	var books []m.Book
	if err = cursor.All(context.TODO(), &books); err != nil {
		log.Printf("Could not decode books.\nError: %v", err)
		util.ResponseError(c, err)
		return
	}

	util.ResponseSuccess(c, books)
}

// getBook retrieves the book with the given bookid
// return the information of the book
// request: GET "/api/v1/book/:bookid"
// response: {...BOOK}
func (s *Service) getBook(c *gin.Context) {

	id := c.Param("bookid")

	// convert to primitiv.ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		util.ResponseError(c, err)
		return
	}

	// get book
	var book m.Book
	res := s.bookCollection.FindOne(context.TODO(), bson.M{"id": oid})
	res.Decode(&book)

	// check if the book exist
	// if no, mongodb will return primitive.NilObjectID
	if book.ID == primitive.NilObjectID {
		util.ResponseError(c, fmt.Errorf("Book %v does not exist\n", id))
		return
	}

	util.ResponseSuccess(c, book)

}

// addBook receives all information of a book and insert one in db
// returns the id of the newly created book
// request: POST "/api/v1/book" form-data: {...BOOK}
// response: `string(primitive.ObjectID)` BOOK_ID
func (s *Service) addBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	status, _ := strconv.Atoi(c.PostForm("status"))
	startTime, _ := time.Parse(layoutISO, c.PostForm("startTime"))
	endTime, _ := time.Parse(layoutISO, c.PostForm("endTime"))
	description := c.PostForm("description")
	notes := []primitive.ObjectID{}
	book := m.Book{
		Title:       title,
		Author:      author,
		Status:      status,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
		Notes:       notes,
	}

	// self generated id field
	book.ID = primitive.NewObjectID()

	_, err := s.bookCollection.InsertOne(context.TODO(), book)
	if err != nil {
		log.Printf("Could not create Book: %v", err)
		util.ResponseError(c, err)
	}

	util.ResponseSuccess(c, book.ID)

}

// deleteBook delete the book with the given id and all its notes
// request: DELETE "/api/v1/book" form-data: {id: ID}
func (s *Service) deleteBook(c *gin.Context) {
	id := c.PostForm("id")
	// convert to primitiv.ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		util.ResponseError(c, err)
		return
	}

	// delete notes where note.bookID = bookID
	// TODO: panicked
	if _, err := s.noteCollection.DeleteMany(context.TODO(), bson.D{{Key: "bookid", Value: oid}}); err != nil {
		log.Printf("Could not delete book %v's notes.\nError: %v", oid, err)
		util.ResponseError(c, err)
		return
	}

	// delete Book
	res, err := s.bookCollection.DeleteOne(context.TODO(), bson.M{"id": oid})
	if err != nil {
		// TODO: restore all deleted notes
		log.Printf("Could not delete book %v.\nError: %v", oid, err)
		util.ResponseError(c, err)
		return
	}

	util.ResponseSuccess(c, int(res.DeletedCount))

}

// editBook edit the book with the given id
// request: POST "/api/v1/book/:bookid" form-data: {...FIELD(s)}
// response: {...EDITED_BOOK}
func (s *Service) editBook(c *gin.Context) {
	fields := make(map[string]interface{})
	fields["id"], _ = primitive.ObjectIDFromHex(c.Param("bookid"))
	fields["title"] = c.PostForm("title")
	fields["author"] = c.PostForm("author")
	fields["status"], _ = strconv.Atoi(c.PostForm("status"))
	fields["startTime"], _ = time.Parse(layoutISO, c.PostForm("startTime"))
	fields["endTime"], _ = time.Parse(layoutISO, c.PostForm("endTime"))
	fields["description"] = c.PostForm("description")

	var updateFields bson.D
	for k, v := range fields {
		if v != "" {
			updateFields = append(updateFields, bson.E{Key: k, Value: v})
		}
	}
	var updatedDocument bson.M
	err := s.bookCollection.FindOneAndUpdate(
		context.TODO(),
		bson.D{{Key: "id", Value: fields["id"]}},
		bson.D{
			{Key: "$set", Value: updateFields},
		},
	).Decode(&updatedDocument)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			util.ResponseError(c, fmt.Errorf("id %v does not match any book\n", fields["id"]))
			return
		}
		log.Printf("Could not edit book %v.\nError: %v", fields["id"], err)
		util.ResponseError(c, err)
		return
	}
	// TODO: updatedDocument is not updated
	util.ResponseSuccess(c, updatedDocument)
}
