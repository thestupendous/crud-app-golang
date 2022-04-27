package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	/*if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}*/
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	//connect
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	//coll := client.Database("sample_mflix").Collection("movies")

	//write
	//	coll := client.Database("myDB").Collection("favorite_books")
	//	docs := []interface{}{
	//		bson.D{{"title", "My Brilliant Friend"}, {"author", "Elena Ferrante"}, {"year_published", 2012}},
	//		bson.D{{"title", "Lucy"}, {"author", "Jamaica Kincaid"}, {"year_published", 2002}},
	//		bson.D{{"title", "Cat's Cradle"}, {"author", "Kurt Vonnegut Jr."}, {"year_published", 1998}},
	//	}
	//	result, err := coll.InsertMany(context.TODO(), docs)
	//	list_ids := result.InsertedIDs
	//	fmt.Printf("Documents inserted: %v\n", len(list_ids))
	//	for _, id := range list_ids {
	//		fmt.Printf("Inserted document with _id: %v\n", id)
	//	}

	//read
	var result bson.M
	coll := client.Database("myDB").Collection("favorite_books")
	title := "My Brilliant Friend"
	//	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s\n", jsonData)
}
