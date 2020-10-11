package tracker

import (
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func Server() {

	// use default
	router := gin.Default()

	// router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// router.Use(gin.Recovery())

	// session
	// router.Use(sessions.Sessions("go_lib", cookie.NewStore([]byte("secret"))))

	// authentication
	// authorized := router.Group("/user")
	// authorized.Use(AuthRequired)

	// router.LoadHTMLGlob("templates/*")
	// router.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"intro": "This is a library management system built with Gin and MongoDB!",
	// 	})
	// })

	router.Use(cors.New(cors.Config{
		// for dev
		AllowAllOrigins: true,
		// for prod
		// AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Origin"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	book := router.Group("/book")
	{
		book.GET("", ListBook)
		book.POST("", AddBook)
		book.GET("/:bookid", GetBook)
		book.DELETE("", DeleteBook)
		book.POST("/:bookid", EditBook)
	}

	note := router.Group("/note")
	{
		note.GET("", ListNoteByBook)
		note.POST("", AddNote)
		note.GET("/:noteid", GetNote)
		note.DELETE("/:noteid", DeleteNote)
		// note.POST("/:noteid", EditNote)
	}
	router.Run(":8989")
}
