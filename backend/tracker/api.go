package tracker

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const layoutISO = "2006-01-02 15:04:05"

func ListBook(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	author := c.Query("author")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	filter := map[string]string{
		"id":        id,
		"title":     title,
		"author":    author,
		"startTime": startTime,
		"endTime":   endTime,
	}
	books, err := listBook(filter)
	if err != nil {
		ResponseBadRequest(c, err)
	} else {
		ResponseSuccess(c, books)
	}
}

func GetBook(c *gin.Context) {
	id := c.Param("bookid")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		ResponseFailure(c, err, 504)
	}
	book, err := getBook(oid)
	if err != nil {
		ResponseBadRequest(c, err)
	} else {
		ResponseSuccess(c, book)
	}
}

func AddBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	status, _ := strconv.Atoi(c.PostForm("status"))
	startTime, _ := time.Parse(layoutISO, c.PostForm("startTime"))
	endTime, _ := time.Parse(layoutISO, c.PostForm("endTime"))
	description := c.PostForm("description")
	notes := []primitive.ObjectID{}
	book := Book{
		Title:       title,
		Author:      author,
		Status:      status,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: description,
		Notes:       notes,
	}
	oid, err := addBook(&book)
	if err != nil {
		ResponseBadRequest(c, err)
	}
	ResponseSuccess(c, oid)
}

func DeleteBook(c *gin.Context) {
	id := c.PostForm("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		ResponseFailure(c, err, 504)
		return
	}
	deleteCount, err := deleteBook(oid)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	} else {
		ResponseSuccess(c, deleteCount)
		return
	}
}

func EditBook(c *gin.Context) {
	fields := make(map[string]interface{})
	fields["id"], _ = primitive.ObjectIDFromHex(c.Param("bookid"))
	fields["title"] = c.PostForm("title")
	fields["author"] = c.PostForm("author")
	fields["status"], _ = strconv.Atoi(c.PostForm("status"))
	fields["startTime"], _ = time.Parse(layoutISO, c.PostForm("startTime"))
	fields["endTime"], _ = time.Parse(layoutISO, c.PostForm("endTime"))
	fields["description"] = c.PostForm("description")

	editCount, err := editBook(fields)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	} else {
		ResponseSuccess(c, editCount)
		return
	}
}

// Note
func ListNoteByBook(c *gin.Context) {
	id := c.Query("bookid")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		ResponseFailure(c, err, 504)
		return
	}
	notes, err := listNoteByBook(oid)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	} else {
		ResponseSuccess(c, notes)
		return
	}
}

func GetNote(c *gin.Context) {
	id := c.Param("noteid")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		ResponseFailure(c, err, 504)
		return
	}
	note, err := getNote(oid)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	} else {
		ResponseSuccess(c, note)
		return
	}
}

func AddNote(c *gin.Context) {
	id := c.PostForm("bookID")
	content := c.PostForm("content")
	title := c.PostForm("title")
	// createTime := c.PostForm("createTime")

	bookID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		ResponseFailure(c, err, 504)
		return
	}
	note := Note{
		BookID:  bookID,
		Content: content,
		Title:   title,
	}

	oid, err := addNote(bookID, &note)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	}
	ResponseSuccess(c, oid)
	return
}

func DeleteNote(c *gin.Context) {
	id := c.PostForm("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
		ResponseFailure(c, err, 504)
		return
	}
	deleteCount, err := deleteNote(oid)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	} else {
		ResponseSuccess(c, deleteCount)
		return
	}
}

func EditNote(c *gin.Context) {
	fields := make(map[string]interface{})
	fields["id"], _ = primitive.ObjectIDFromHex(c.Param("noteid"))
	fields["title"] = c.PostForm("title")
	fields["content"] = c.PostForm("content")

	editCount, err := editNote(fields)
	if err != nil {
		ResponseBadRequest(c, err)
		return
	} else {
		ResponseSuccess(c, editCount)
		return
	}
}
