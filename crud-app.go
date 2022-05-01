package main

import (
	//	"encoding/json"
	//	"log"
	//	"net/http"
	//	"math/rand"
	//	"strconv"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	//logging related headers
	"io"
	"log"
	"os"

	"github.com/gorilla/mux"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx  = context.TODO()
	coll *mongo.Collection
	//usernames []GamerUname
)

type GamerUname struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"user_name" bson:"user_name"`
}

func getUnames(writer http.ResponseWriter, request *http.Request) {
	log.Println("in getAll:")
	writer.Header().Set("Content-Type", "application/json")
	resultStructs := getUnamesDB()
	//	log.Println("result from db getall function : ", resultStructs)
	json.NewEncoder(writer).Encode(resultStructs)

}

func getUname(writer http.ResponseWriter, request *http.Request) {
	log.Println("in getByID:")
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	requestedID := params["id"]

	resultUnameStruct := getUnameDB(requestedID)

	json.NewEncoder(writer).Encode(resultUnameStruct)

}

func createUname(writer http.ResponseWriter, request *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	log.Println("in Create:")
	writer.Header().Set("Content-Type", "application/json")
	var username GamerUname
	_ = json.NewDecoder(request.Body).Decode(&username)
	username.ID = strconv.Itoa(rand.Intn(99999)) //TODO check for duplicates
	out, err := createUnameDB(username)
	if err != nil {
		log.Println("error in object creation in mongo : ", err)
	}
	_ = out //for now ignoring mongodb output value
	//	usernames = append(usernames, username)
	returnMsg := struct {
		ResCode int
		Message string
	}{0, "Successull, added a new record"}
	if err != nil {
		returnMsg.ResCode = 1
		returnMsg.Message = "Error in db record creation!"
	} else {
		returnMsg.ResCode = 0
		returnMsg.Message = "Successfully added record to db"
	}

	log.Printf("create> type : %T : %+v\n", returnMsg, returnMsg)
	json.NewEncoder(writer).Encode(returnMsg)
}

func deleteUname(writer http.ResponseWriter, request *http.Request) {
	log.Println("in delete:")
	//	var found bool
	//	found = false
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	requestedID := params["id"]
	//	for index, item := range usernames {
	//		if item.ID == params["id"] {
	//			found = true
	//			usernames = append(usernames[:index], usernames[index+1:]...)
	//			break
	//		}
	//	}
	result, err := deleteUnameDB(requestedID)
	_ = result
	returnMsg := struct {
		ResCode int
		Message string
	}{0, ""}
	if err == nil {
		returnMsg.ResCode = 0
		returnMsg.Message = "Successful, deleted one record"
	} else {
		returnMsg.ResCode = 1
		returnMsg.Message = "Unsuccessful!! could not find the ID provided"
	}
	log.Printf("On Deletions> type : %T : %+v\n", returnMsg, returnMsg)

	json.NewEncoder(writer).Encode(returnMsg)

}

func updateUname(writer http.ResponseWriter, request *http.Request) {
	log.Println("in update:")

	var username GamerUname

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	requestedID := params["id"]
	err := json.NewDecoder(request.Body).Decode(&username)
	if err != nil {
		log.Println("error in reading request body : ", err)
	}

	res, err := updateUnameDB(requestedID, username)
	if err != nil {
		log.Println("error in update method from mongo : ", err)
	}
	_ = res

	returnMsg := struct {
		ResCode int
		Message string
	}{0, ""}
	if err == nil {
		returnMsg.ResCode = 0
		returnMsg.Message = "Successful, updated one record"
	} else {

		returnMsg.Message = "Unsuccessful!! could not find the provide ID in records!"
	}
	log.Printf("update> type : %T : %+v\n", returnMsg, returnMsg)
	json.NewEncoder(writer).Encode(returnMsg)

}

func main() {
	//setting up logging
	logFile, _ := os.OpenFile("output-of-crud-app.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	//usernames = make([]GamerUname, 0)
	//init router
	router := mux.NewRouter()

	initMongoConnection()

	//router handlers
	router.HandleFunc("/usermapper", getUnames).Methods("GET")
	router.HandleFunc("/usermapper/{id}", getUname).Methods("GET")
	router.HandleFunc("/usermapper", createUname).Methods("POST")
	router.HandleFunc("/usermapper/{id}", updateUname).Methods("PUT")
	router.HandleFunc("/usermapper/{id}", deleteUname).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}

func createUnameDB(username GamerUname) (string, error) {
	log.Println("in createDB:")
	result, err := coll.InsertOne(ctx, username)
	if err != nil {
		log.Println("error in inserting to mongo : ", err)
		return "0", err
	}
	return fmt.Sprintf("%v", result.InsertedID), err
}

func getUnamesDB() interface{} {
	log.Println("in getAll DB: ")
	var results []bson.D
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println("error in find() function: ", err)

	}

	//	var resultStruct GamerUname
	for cursor.Next(ctx) {
		var result bson.D
		if err := cursor.Decode(&result); err != nil {
			log.Println("error in cursor decoding : ", err)
		}
		//convert bson result field to struct
		//		bytes, err := json.Marshal(result)
		results = append(results, result)
	}

	//converting []bson.D data to string
	bytes, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("error in marshaling to string : ", err)
	}
	// log.Println("printing all bytes : ", bytes)

	//converting string to list of structs
	var resultStructs interface{}
	err = json.Unmarshal(bytes, &resultStructs)
	if err != nil {
		fmt.Println("error in unmarshalling to struct slice : ", err)
	}
	// log.Println(resultStructs)

	//	resultJson, err = json.Marshal(results)
	return resultStructs

}

func initMongoConnection() {
	host := "localhost"
	port := "27017"
	connectionURI := "mongodb://" + host + ":" + port + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("error in mongo connection : ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("error in pinging the connection : ", err)
	} else {
		log.Println("NO error in pinging the connection to mongodb")
	}

	db := client.Database("golang_project")
	coll = db.Collection("username_mapper")
}

func getUnameDB(requestedID string) GamerUname {
	log.Println("in getByID DB: ")
	var username GamerUname
	var result bson.M
	filter := bson.D{{"_id", requestedID}}
	err := coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println("error in findONe method on mongodb : ", err)
	}
	bytes, err := bson.Marshal(result)

	err = bson.Unmarshal(bytes, &username)
	if err != nil {
		log.Println("error in unmarshal : ", err)
	}

	return username
}

func deleteUnameDB(requestedID string) (string, error) {
	log.Println("in deleteDB :")
	var deletedDoc bson.D
	filter := bson.D{{"_id", requestedID}}
	err := coll.FindOneAndDelete(ctx, filter).Decode(&deletedDoc)
	if err != nil {
		log.Println("error in findAndDelete method from mongo : ", err)
	}
	resString := fmt.Sprintf("%v", deletedDoc)
	return resString, err

}

// updatUnameDB function updates a document by replacing its complete
// body, and applying a filter by id
func updateUnameDB(requestedID string, username GamerUname) (string, error) {
	log.Println("in updateDB :")
	filter := bson.D{{"_id", requestedID}}
	update := username

	updatedDoc, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("error in update call to mongodb : ", err)
	}

	resString := fmt.Sprintf("%v", updatedDoc)
	return resString, err

}
