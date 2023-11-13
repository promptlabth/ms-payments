# Use an official Go runtime as a parent image
FROM golang:1.21-alpine

# Set the working directory to /app
WORKDIR /app

# Copy file
COPY go.mod go.sum ./

# Download package the Go application
RUN go mod download

# Copy the rest
COPY . /app

# Install the Go package using go get air
RUN go install github.com/cosmtrek/air@latest

# Expose the port your Go application will listen on
EXPOSE 8080

CMD ["air"]
