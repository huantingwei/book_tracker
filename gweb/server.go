package gweb

import (
	"github.com/gin-gonic/gin"
)

func DemoServer() {

	router := gin.Default()

	router.GET("/login", login)
	router.GET("/logout", logout)

	user := router.Group("/user")
	{
		user.GET("/profile", userProfile)
		user.GET("/setting", userSetting)
	}

	books := router.Group("/books")
	{
		books.POST("/add", addBook)
		books.GET("", listBook)
		books.GET("/get", getBookByID)
		books.POST("/borrow", borrowBookByID)
		books.POST("/return", returnBookByID)
	}
	router.Run(":8989")
}
