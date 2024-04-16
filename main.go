package main

import api "github.com/Rudro-25/Book_API_Server/apiHandler"

func main() {
	api.Start()
}

//go run main.go
//Get all book list
//curl localhost:8080/books
//Get Book by Id
//curl localhost:8080/books/1
//Add a book:
//curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
//Delete some copy of book
//curl localhost:8080/checkout/2?id=1 --request "PATCH"
//Return some copy of book
//curl localhost:8080/return/3?id=1 --request "PATCH"
//full delete a data by id
//curl -X DELETE localhost:8080/books/1
//Update data of a specific book by id [off-field will be unchanged]
//curl -X PATCH -H "Content-Type: application/json" -d '{"title": "New Title", "author": "Rudro", "quantity": 10}' localhost:8080/books/1
