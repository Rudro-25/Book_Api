package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context) {
	cnt, err := strconv.Atoi(c.Param("cnt"))

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

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

func returnBook(c *gin.Context) {
	cnt, err := strconv.Atoi(c.Param("cnt"))

	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found."})
		return
	}

	book.Quantity += cnt
	c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func updateBookById(c *gin.Context) {
	id := c.Param("id")

	var ind int = -1

	for i, b := range books {
		if b.ID == id {
			ind = i
			break
		}
	}

	if ind == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found."})
		return
	}

	var updatedBook book
	err := c.BindJSON(&updatedBook)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid request body."})
		return
	}

	if updatedBook.Title != "" {
		books[ind].Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		books[ind].Author = updatedBook.Author
	}
	if updatedBook.Quantity != 0 {
		books[ind].Quantity = updatedBook.Quantity
	}

	c.IndentedJSON(http.StatusOK, books[ind])
}

func deleteBookById(c *gin.Context) {
	id := c.Param("id")

	var ind int = -1

	for i, b := range books {
		if b.ID == id {
			ind = i
			break
		}
	}

	if ind == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not Found."})
		return
	}

	books = append(books[:ind], books[ind+1:]...)

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Deleted successfully."})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout/:cnt", checkoutBook)
	router.PATCH("/return/:cnt", returnBook)
	router.PATCH("/books/:id", updateBookById)
	router.DELETE("/books/:id", deleteBookById)
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
