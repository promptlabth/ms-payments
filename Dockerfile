# Use an official Go runtime as a parent image
FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main

CMD [ "./main" ]
