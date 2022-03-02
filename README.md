# GO REST API CLEANCODE

Rest api based on cleancode architecture

## Installation

```
$ go get github.com/Yefhem/go-rest-api-cleancode
```

## How to run

### Required

- Postgresql

### Conf

You should modify `.env`

```
# database
DB_DRIVER=postgres
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=postgres

JWT_SECRET=postgres
```

### Run
```
$ go run cmd/go-rest-api-cleancode/main.go 
```