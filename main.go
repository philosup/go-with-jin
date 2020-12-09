package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConn() (client *mongo.Client) {
	// credential := options.Credential{
	//    Username: "<USER_NAME>",
	//    Password: "<PASSWORD>",
	// }
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") //.SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	return client
}

type Post struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1",
		})
	})

	r.POST("/post", func(c *gin.Context) {

		var post Post

		c.BindJSON(&post)

		client := mongoConn()

		// post := Post{title, body}
		collection := client.Database("testdb").Collection("post")
		insertResult, err := collection.InsertOne(context.TODO(), post)

		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, insertResult)
	})

	r.GET("/post", func(c *gin.Context) {

		client := mongoConn()

		collection := client.Database("testdb").Collection("post")

		filter := bson.D{}
		var post Post
		err := collection.FindOne(context.TODO(), filter).Decode(&post)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, post)
	})

	r.GET("/post/all", func(c *gin.Context) {

		client := mongoConn()

		collection := client.Database("testdb").Collection("post")

		filter := bson.M{}
		cursor, err := collection.Find(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		var posts []Post
		if err = cursor.All(context.TODO(), &posts); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, posts)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
