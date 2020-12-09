package main

import "github.com/huantingwei/go/tracker"

// "context"
// "fmt"
// "log"
// "time"

// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"
// "go.mongodb.org/mongo-driver/mongo/readpref"

func main_bak() {
	tracker.Server()
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)

	// r := gin.Default()
	// store := cookie.NewStore([]byte("secret"))
	// r.Use(sessions.Sessions("mysession", store))

	// r.GET("/hello", func(c *gin.Context) {
	// 	session := sessions.Default(c)

	// 	if session.Get("hello") != "world" {
	// 		session.Set("hello", "world")
	// 		session.Save()
	// 	}

	// 	c.JSON(200, gin.H{"hello": session.Get("hello")})
	// })
	// r.Run(":8000")
}
