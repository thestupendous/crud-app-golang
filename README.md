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

- endpoints
    - getall records `/usermapper `
    - get by id ` /usermapper/<id> `
    - create a record `/usermapper `
    - update a record `/usermapper/<id> `
    - delete a record `/usermapper/<id> `


