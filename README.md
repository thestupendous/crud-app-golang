# CRUD Application 
A crud app which uses Mongodb for storing and retrieving the data.\
Gorilla mux is used along with http library for setting up http\
server. This app's code can be used for creating a generic crud app\
in golang.\
The program listens on **8080** port on localhost.


## Setting up Env and Running the App
Things that need to be set up before the program can run, and how to \
actually run the program.
- install golang, mongodb (on localhost)
- commands to run 
    - in the main directory \
    ` go mod tidy ` \
    ` go mod vendor ` \
    ` go run crud-app.go`

## Api doc
data structure
```go 
type GamerUname struct {
        ID       string `json:"id"          bson:"_id"`
        Name     string `json:"name"        bson:"name"`
        Username string `json:"user_name"   bson:"user_name"`
}
 ```
- endpoints
    - getall records `/usermapper `
        - request
            method - GET
        - response
            - header - `"Content-Type: application/json"`
            - body - array of json objects
    - get by id ` /usermapper/<id> `
        - request
            - method - GET
        - response
            - header - `"Content-Type: application/json"`
            - body - single json object
    - create a record `/usermapper `
        - request
            - method - POST
            - body - json object of GamerUname, without the ID field,\
            - ex - 
            ```bash
            curl -X POST localhost:8080/usermapper -H 'Content-Type: application/json' -d '{"name":"Sumita Singh","user_name":"gamerS"}'
            ```
        - response
            - header - `"Content-Type: application/json"`
            - body - Status Message of the form -
            ```go
                struct {
                ResCode int
                Message string
                }{0, "Successull, added a new record"}
            ```
    - update a record `/usermapper/<id> `
        - request
             - method - PUT
             - body - json object without id in body
             - ex - 
            ```bash
            curl -X PUT localhost:8080/usermapper/11066 -H 'Content-Type: application/json' -d '{"name":"Sumita Sigh","user_name":"GamerSumi"}'
            ```
        - response
            - header - `"Content-Type: application/json"`
            - body - Status Message of the form -
            ```go
                struct {
                ResCode int
                Message string
                }{0, "Successful, updated one record"}
            ```
    - delete a record `/usermapper/<id> `
        - request
             - method - DELETE
             - ex - 
            ```bash
            curl -X DELETE localhost:8080/usermapper/11066
            ```
        - response
            - header - `"Content-Type: application/json"`
            - body - Status Message of the form -
            ```go
                struct {
                ResCode int
                Message string
                }{0, "Successful, deleted one record"}
            ```


