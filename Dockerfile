# Use an official Go runtime as a parent image
FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go get

RUN go build -o main

EXPOSE 8000

CMD [ "./main" ]
