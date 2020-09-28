package gweb

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func listBook(c *gin.Context) {
	ListBook()
	ResponseSuccess(c, 1)
}

func borrowBook(c *gin.Context) {
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

func getBook(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	author := c.Query("author")
	fmt.Println(id, name, author)
}
