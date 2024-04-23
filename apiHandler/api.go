package apiHandler

import (
	"errors"
	"fmt"
	"github.com/Rudro-25/book_api_server/authHandler"
	"github.com/Rudro-25/book_api_server/dataHandler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dataHandler.Books)
}

func GetBookById(id string) (*dataHandler.Book, error) {
	for i, b := range dataHandler.Books {
		if b.ID == id {
			return &dataHandler.Books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func CheckoutBook(c *gin.Context) {
	cnt, err := strconv.Atoi(c.Param("cnt"))

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found."})
		return
	}

	if book.Quantity < 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book Not available."})
		return
	}

	book.Quantity -= cnt
	c.IndentedJSON(http.StatusOK, book)
}

func ReturnBook(c *gin.Context) {
	cnt, err := strconv.Atoi(c.Param("cnt"))

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found."})
		return
	}

	book.Quantity += cnt
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook dataHandler.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	dataHandler.Books = append(dataHandler.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func UpdateBookById(c *gin.Context) {
	id := c.Param("id")

	var ind int = -1

	for i, b := range dataHandler.Books {
		if b.ID == id {
			ind = i
			break
		}
	}

	if ind == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found."})
		return
	}

	var updatedBook dataHandler.Book
	err := c.BindJSON(&updatedBook)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid request body."})
		return
	}

	if updatedBook.Title != "" {
		dataHandler.Books[ind].Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		dataHandler.Books[ind].Author = updatedBook.Author
	}
	if updatedBook.Quantity != 0 {
		dataHandler.Books[ind].Quantity = updatedBook.Quantity
	}

	c.IndentedJSON(http.StatusOK, dataHandler.Books[ind])
}

func DeleteBookById(c *gin.Context) {
	id := c.Param("id")

	var ind int = -1

	for i, b := range dataHandler.Books {
		if b.ID == id {
			ind = i
			break
		}
	}

	if ind == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found."})
		return
	}

	dataHandler.Books = append(dataHandler.Books[:ind], dataHandler.Books[ind+1:]...)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Deleted successfully."})
}

func AddUser(c *gin.Context) {
	var newUser dataHandler.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	dataHandler.Users = append(dataHandler.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func Start(Port int) {
	router := gin.Default()

	router.POST("/login", authHandler.LoginHandler)
	router.POST("/adduser", AddUser)

	router.GET("/books", GetBooks)
	router.GET("/books/:id", BookById)
	router.POST("/books", CreateBook)
	router.PATCH("/checkout/:cnt", CheckoutBook)
	router.PATCH("/return/:cnt", ReturnBook)
	router.PATCH("/books/:id", UpdateBookById)
	router.DELETE("/books/:id", DeleteBookById)
	err := router.Run(fmt.Sprintf(":%v", Port))
	if err != nil {
		log.Fatal(err)
	}
}
