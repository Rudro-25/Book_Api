#api-server
 - Simple CRUD API Server with Golang

#Technogy-used
- Golang (Gin Framework)

#Running the server
- git clone https://github.com/Rudro-25/Book_Api.git


- go run main.go


#Api-Calls

| # | Method | url                                   | Actions                 | Example |
|---|--------|---------------------------------------|-------------------------|----|
01 | POST   | http://localhost:8080/login           | login Authentication    | |
02 | POST   | http://localhost:8080/adduser         | User Data Add           | |
03 | GET    | http://localhost:8080/books           | Get All Book List       | |
04 | GET    | http://localhost:8080/books/1         | Find Book By Id         | |
05 | POST   | http://localhost:8080/books           | Add New Book Data       | |
06 | PATCH  | http://localhost:8080/checkout/2?id=1 | Checkout a copy of Book | |
07 | PATCH  | http://localhost:8080/return/3?id=1   | Return a copy of book   | |
08 | DELETE | http://localhost:8080/books/1           | Delete a book data      | |
09 | PATCH  | http://localhost:8080/books/1           | Update a book data      | |
10 |  | http://localhost:8080/        |                         | |



