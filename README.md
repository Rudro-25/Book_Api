#api-server
 - Simple CRUD API Server with Golang

#Technogy-used
- Golang (Gin Framework)

#Running the server
- git clone https://github.com/Rudro-25/Book_Api.git


- go run main.go


#Api-Calls

| # | Method | url                                   | Actions                 | Example |
|---|--------|---------------------------------------|-------------------------|---------|
01 | POST   | http://localhost:8080/login           | login Authentication    | x       |
02 | POST   | http://localhost:8080/adduser         | User Data Add           | x       |
03 | GET    | http://localhost:8080/books           | Get All Book List       | x       |
04 | GET    | http://localhost:8080/books/1         | Find Book By Id         | x       |
05 | POST   | http://localhost:8080/books           | Add New Book Data       | x       |
06 | PATCH  | http://localhost:8080/checkout/2?id=1 | Checkout a copy of Book | x       |
07 | PATCH  | http://localhost:8080/return/3?id=1   | Return a copy of book   | x       |
08 | DELETE | http://localhost:8080/books/1           | Delete a book data      | x       |
09 | PATCH  | http://localhost:8080/books/1           | Update a book data      | x       |
10 | x      | http://localhost:8080/        | x                       | x       |



