package main

import (
	"github.com/Rudro-25/Book_API_Server/apiHandler"
	"github.com/gin-gonic/gin"
)

func main() {

	//http.HandleFunc("/login", authHandler.Login)
	//http.HandleFunc("/home", authHandler.Home)
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))

	router := gin.Default()

	router.GET("/books", apiHandler.GetBooks)
	router.GET("/books/:id", apiHandler.BookById)
	router.POST("/books", apiHandler.CreateBook)
	router.PATCH("/checkout/:cnt", apiHandler.CheckoutBook)
	router.PATCH("/return/:cnt", apiHandler.ReturnBook)
	router.PATCH("/books/:id", apiHandler.UpdateBookById)
	router.DELETE("/books/:id", apiHandler.DeleteBookById)
	router.Run("localhost:8080")

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
