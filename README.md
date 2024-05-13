# book-api-server
 - Simple CRUD API Server with Golang

# Technogy-used
- Golang (Gin Framework)
- JWT Auth
- Cobra Cli

# Running the server
- `git clone https://github.com/Rudro-25/Book_Api.git`   

From `book-api-server` directory:
- ` go mod tidy & go mod vector`
- `go run . start or go run . -p <port>`

# Running the server from docker image
- `docker pull rudro25/book-api-server`
- `docker run -dp <port>8080 rudro25/book-api-server`

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



# Deployment and Service

- `kubectl port-forward pod/<pod_name> 8080:8080`
- `kubectl port-forward svc/<service_name> 4025:3200`
 
Check: 

- `http://localhost:8080/books`
- `http://localhost:4025/books`