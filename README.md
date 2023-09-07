# MS-Payments Promptlab AI

The MS-Payments Service is a microservice designed to handle payment transactions. Written in Go, this service utilizes the Gin framework and adheres to the principles of clean architecture.

## Features

- Accept and process payment transactions
- Maintain payment history and records
- Support multiple payment methods

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/promptlabth/ms-payments.git
cd ms-payments
```


### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Build
```bash
go build -o ms-payments
```

## Running Locally

### 1. Start the Service
```bash
./ms-payments
```

### 2. Health Check
To verify the service is running smoothly, navigate to the /health endpoint.

## Testing
To run unit tests, execute:
```bash
go test ./...
```




## in this code will split to 5 part of the code
1. `entities` it will be create a entity and struct for the code
2. `repositories` it will be for manage a database of the struct => select, insert, update, delete etc.
3. `usecases` it will contain a business usecase of the repository => before select should do something ? before insert should before do some thing ?
4. `controllers` it will be for validate a data and prepare a data before send a data to business usecase
5. `routes` is for define a routing url 