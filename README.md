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
DB_USERNAME=
DB_PASSWORD=
DB_NAME=

JWT_SECRET=
```

### Run
```
$ go run cmd/go-rest-api-cleancode/main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/auth/home         --> github.com/Yefhem/rest-api-cleancode/controller.AuthController.Home-fm (3 handlers)       
[GIN-debug] POST   /api/v1/auth/login        --> github.com/Yefhem/rest-api-cleancode/controller.AuthController.Login-fm (3 handlers)      
[GIN-debug] POST   /api/v1/auth/register     --> github.com/Yefhem/rest-api-cleancode/controller.AuthController.Register-fm (3 handlers)   
[GIN-debug] GET    /api/v1/user/profile      --> github.com/Yefhem/rest-api-cleancode/controller.UserController.Profile-fm (4 handlers)    
[GIN-debug] PUT    /api/v1/user/profile      --> github.com/Yefhem/rest-api-cleancode/controller.UserController.Update-fm (4 handlers)     
[GIN-debug] GET    /api/v1/products/         --> github.com/Yefhem/rest-api-cleancode/controller.ProductController.All-fm (4 handlers)     
[GIN-debug] POST   /api/v1/products/         --> github.com/Yefhem/rest-api-cleancode/controller.ProductController.Insert-fm (4 handlers)  
[GIN-debug] GET    /api/v1/products/:id      --> github.com/Yefhem/rest-api-cleancode/controller.ProductController.FindByID-fm (4 handlers)
[GIN-debug] PUT    /api/v1/products/:id      --> github.com/Yefhem/rest-api-cleancode/controller.ProductController.Update-fm (4 handlers)  
[GIN-debug] DELETE /api/v1/products/:id      --> github.com/Yefhem/rest-api-cleancode/controller.ProductController.Delete-fm (4 handlers)  
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

# Directory structure
```
.
├── cmd
│   └── go-rest-api-cleancode
│       └── main.go
├── configs
│   └── db-config.go    // Setup Database Connection
├── controller
│   ├── auth.go
│   ├── product.go
│   └── user.go
├── dto                 // data transfer objects
│   ├── login.go   
│   ├── product.go 
│   ├── register.go 
│   └── user.go     
├── helper
│   └── response.go     // BuildResponse, BuildErrorResponse and EmptyObj
├── middleware
│   └── jwt-auth.go
├── models
|   ├── Product.go       
|   └── User.go   
├── repository
|   ├── product.go
|   └── user.go
├── service
|   ├── auth.go
|   ├── jwt.go
|   ├── product.go
|   └── user.go
|
...
```

## Features

- RESTful API
- Gorm
- Jwt-go
- Gin
- Cleancode