package main

import (
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/huantingwei/go/book"
	"github.com/huantingwei/go/util"
)

func main() {

	mongoURI := "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"

	db, ctx := util.NewDatabase(mongoURI)
	defer db.Client.Disconnect(ctx)

	router := gin.Default()

	// avoid CORS
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

	// APIs

	v1 := router.Group("/api/v1")
	{
		book.NewService(v1, db)
	}

	router.Run(":8989")

}
