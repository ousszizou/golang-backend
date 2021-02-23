// https://github.com/mongodb/mongo-go-driver
// https://github.com/kamva/mgm
// https://github.com/qiniu/qmgo

package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post Type
type Post struct {
	Title string `json:"title"`

	Body string `json:"body"`
}

// Used to create a singleton object of MongoDB client.
var clientInstance *mongo.Client

// Used during creation of singleton client object
var clientInstanceError error

// Used to execute client creation procedure only once.
var mongoOnce sync.Once

// database config's.
const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "blog"
	POSTS            = "posts"
)

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func insertPost(post Post) error {
	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(DB).Collection(POSTS)

	_, err = collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func getPosts() ([]Post, error) {
	posts := []Post{}

	client, err := GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(DB).Collection(POSTS)

	cursor, FindError := collection.Find(context.TODO(), bson.D{})

	if FindError != nil {
		return posts, FindError
	}

	for cursor.Next(context.TODO()) {
		post := Post{}
		err := cursor.Decode(&post)

		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	cursor.Close(context.TODO())

	if len(posts) == 0 {
		return posts, mongo.ErrNoDocuments
	}

	return posts, nil
}

func main() {

	// Get ALL
	posts, _ := getPosts()
	fmt.Println(posts)

	// INSERT
	post := Post{Title: "Learn Go", Body: "Golang tutorial by algortihm"}
	err := insertPost(post)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("new post was added to DB")
}
