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

	"github.com/gorilla/mux"
)

var (
	usernames []GamerUname
)

type GamerUname struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"user_name" bson:"user_name"`
}

func getUnames(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(usernames)

}

func getUname(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range usernames {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&GamerUname{})

}

func createUname(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var username GamerUname
	_ = json.NewDecoder(request.Body).Decode(&username)
	username.ID = strconv.Itoa(rand.Intn(99999)) //TODO check for duplicates
	usernames = append(usernames, username)
	json.NewEncoder(writer).Encode(username)
}

func deleteUname(writer http.ResponseWriter, request *http.Request) {
	var found bool
	found = false
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range usernames {
		if item.ID == params["id"] {
			found = true
			usernames = append(usernames[:index], usernames[index+1:]...)
			break
		}
	}
	returnMsg := struct {
		resCode int
		message string
	}{0, ""}
	if found {
		returnMsg.resCode = 0
		returnMsg.message = "Successful"
	} else {
		returnMsg.resCode = 1
		returnMsg.message = "Unsuccessful!!"
	}
	fmt.Printf("On Deletions> type : %T : %+v\n", returnMsg, returnMsg)

	json.NewEncoder(writer).Encode(returnMsg)

}

func updateUname(writer http.ResponseWriter, request *http.Request) {

	var username GamerUname

	var found bool
	found = false
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range usernames {
		if item.ID == params["id"] {
			found = true
			usernames = append(usernames[:index], usernames[index+1:]...)

			//change values with request body
			_ = json.NewDecoder(request.Body).Decode(&username)
			username.ID = strconv.Itoa(rand.Intn(99999))
			//			fmt.Printf("\t\treceived body object: %T %+v\n", username, username)
			usernames = append(usernames, username)
			//			fmt.Println("\t\t*****NOW PRINTING ALL DATA\n", usernames)

			break
		}
	}
	returnMsg := struct {
		resCode int
		message string
	}{0, ""}
	if found {
		returnMsg.resCode = 0
		returnMsg.message = "Successful"
	} else {
		returnMsg.resCode = 1
		returnMsg.message = "Unsuccessful!!"
	}
	fmt.Printf("update> type : %T : %+v\n", returnMsg, returnMsg)
	json.NewEncoder(writer).Encode(returnMsg)

}

func main() {
	usernames = make([]GamerUname, 0)
	//init router
	router := mux.NewRouter()

	//router handlers
	router.HandleFunc("/books", getUnames).Methods("GET")
	router.HandleFunc("/books/{id}", getUnames).Methods("GET")
	router.HandleFunc("/books", createUname).Methods("POST")
	router.HandleFunc("/books/{id}", updateUname).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteUname).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
