package main

import (
	"github.com/Rudro-25/Book_API_Server/apiHandler"
	"github.com/Rudro-25/Book_API_Server/authHandler"
	"github.com/gin-gonic/gin"
)

func main() {

	//http.HandleFunc("/login", authHandler.Login)
	//http.HandleFunc("/home", authHandler.Home)
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))

	router := gin.Default()

	router.POST("/login", authHandler.LoginHandler)
	router.POST("/adduser", apiHandler.AddUser)

	router.GET("/books", apiHandler.GetBooks)
	router.GET("/books/:id", apiHandler.BookById)
	router.POST("/books", apiHandler.CreateBook)
	router.PATCH("/checkout/:cnt", apiHandler.CheckoutBook)
	router.PATCH("/return/:cnt", apiHandler.ReturnBook)
	router.PATCH("/books/:id", apiHandler.UpdateBookById)
	router.DELETE("/books/:id", apiHandler.DeleteBookById)
	router.Run("localhost:8080")

}
