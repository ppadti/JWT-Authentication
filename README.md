# Golang Jwt Authentication Project
This project provides a simple authentication system implemented in Go (Golang), utilizing various libraries and technologies such as Fiber, GORM, CORS, MySQL, bcrypt, and JWT-Go.

## Features
  - User registration
  - User login
  - User logout
  - Retrieving a list of users


## Technologies Used
 **[Fiber](https://docs.gofiber.io/)**: A web framework for Go, known for its performance and ease of use.

**[GORM](https://gorm.io/docs/)**: An ORM library for Go, used for interacting with the MySQL database.

**[CORS](https://docs.gofiber.io/api/middleware/cors/)**: Cross-Origin Resource Sharing middleware to allow secure cross-origin requests.

**[MySQL](https://docs.fedoraproject.org/en-US/quick-docs/installing-mysql-mariadb/#_installing_mysql_on_fedora)**: A popular relational database management system used for storing user data.

**[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)**: A library for hashing passwords securely.

**[JWT-Go](https://pkg.go.dev/github.com/dgrijalva/jwt-go)**: A JSON Web Token implementation for Go, used for authentication purposes.

## Steps to install:
**1. Clone the repository:**
   
`git clone https://github.com/ppadti/JWT-Authentication.git`

`cd JWT-Authentication`

**2. Install dependencies:**

`go mod tidy`

Set up the MySQL database. Make sure to update the database configuration in `database/connection.go`

**3. Run the project:**

 `go run main.go`

## Endpoints
**POST** `/api/register`: Register a new user. Requires a JSON body with username, email and password fields.

**POST** `/api/login` : Login with username and password. Returns a JWT token upon successful authentication.

**POST** `/api/logout`: Logout the user. Requires a valid JWT token.

**GET** `/api/users` : Retrieve a list of users. Requires a valid JWT token.
 

