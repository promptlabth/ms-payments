FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

ENV DB_HOST=34.87.65.242
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=pR0mptL#bR3s3arch!
ENV DB_NAME=promptLab_test

EXPOSE 8080

CMD ["./main"]