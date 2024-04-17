package apiHandler

import (
	"errors"
	"github.com/Rudro-25/Book_API_Server/dataHandler"
	"github.com/gin-gonic/gin"
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

//func Start() {
//	router := gin.Default()
//
//	router.POST("/login", login)
//	//http.HandleFunc("/home", Home)
//	//http.HandleFunc("/refresh", Refresh)
//
//	router.GET("/books", getBooks)
//	router.GET("/books/:id", bookById)
//	router.POST("/books", createBook)
//	router.PATCH("/checkout/:cnt", checkoutBook)
//	router.PATCH("/return/:cnt", returnBook)
//	router.PATCH("/books/:id", updateBookById)
//	router.DELETE("/books/:id", deleteBookById)
//	router.Run("localhost:8080")
//}

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
