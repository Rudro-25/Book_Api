# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Containerize an application : https://docs.docker.com/get-started/02_our_app/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Rudro Debnath <rudro.cse5.bu@gmail.com>"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the working directory inside the container
Copy . .

# Build the GO app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o book_api_server .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/book_api_server .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./book_api_server", "start"]