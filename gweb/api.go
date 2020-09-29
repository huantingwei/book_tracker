package gweb

import (
	"github.com/gin-gonic/gin"
)

func listBook(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	author := c.Query("author")
	filter := map[string]string{
		"id":     id,
		"name":   name,
		"author": author,
	}
	books := ListBookWithFilter(filter)
	ResponseSuccess(c, books)
}

func getBookByID(c *gin.Context) {
	id := c.Query("id")
	book := GetBookByID(id)
	ResponseSuccess(c, book)
}

func addBook(c *gin.Context) {
	name := c.PostForm("name")
	author := c.PostForm("author")

	book := Book{Name: name, Author: author, Status: 1}

	_, err := AddOneBook(&book)
	if err != nil {
		ResponseError(c, err)
	}
	ResponseSuccess(c, book)
}

func borrowBookByID(c *gin.Context) {
	id := c.PostForm("id")
	err := BorrowBookByID(id)
	if err != nil {
		ResponseError(c, err)
	} else {
		ResponseSuccess(c, nil)
	}
}

func returnBookByID(c *gin.Context) {
	id := c.PostForm("id")
	err := ReturnBookByID(id)
	if err != nil {
		ResponseError(c, err)
	} else {
		ResponseSuccess(c, nil)
	}
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"token": 1,
	})
}

func logout(c *gin.Context) {
	ResponseSuccess(c, "1")
}

func userProfile(c *gin.Context) {
}

func userSetting(c *gin.Context) {
}
