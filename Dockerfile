FROM golang:1.21-alpine

WORKDIR /app

COPY . /app

RUN go mod tidy

RUN go build -o ms-payments

EXPOSE 8080

CMD ./ms-payments
 