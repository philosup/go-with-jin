package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
func initMongoDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
}

type Post struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

func InsertPost(title string, body string) {

	client := mongoConn()

	post := Post{title, body}
	collection := client.Database("testdb").Collection("post")
	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted post with ID:", insertResult.InsertedID)
}

// func GetPost(id bson.ObjectId) {
// 	collection := client.Database("testdb").Collection("post")
// 	filter := bson.D
// 	var post Post

// 	err := collection.FindOne(context.TODO(), filter).Decode(&post)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Found post with title ", post.Title)
// }

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	InsertPost("test1", "body1")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
