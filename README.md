# Gin JWT Authentication/Authorization

A simple Signup & Login API built with Go using the Gin framework. This API includes authentication and authorization using JWT and password hashing with bcrypt.

## Features
- User signup and login with hashed passwords
- JWT-based authentication and authorization middleware
- PostgreSQL as the database
- Dockerized setup with Docker Compose

## Tech Stack
- **Go 1.23** (Gin framework)
- **GORM** (ORM for database interaction)
- **PostgreSQL** (Database)
- **JWT** (Authentication)
- **Bcrypt** (Password hashing)
- **Docker & Docker Compose** (Containerization)

## Middleware
### Authentication Middleware (`middlewares.ValidateAuth`)
- Verifies the JWT token from the request Cookie header
- Attaches user data to the request context

## Installation & Setup

### Prerequisites
- Go installed (>=1.23)
- Docker & Docker Compose installed
- `.env` file configured with database credentials

### Clone the Repository
```sh
$ git clone https://github.com/dawood-usman/gin-jwt-authentication.git
$ cd gin-jwt-authentication
```

### Setup Environment Variables By Copying `env.example` file:
   ```sh
$ cp env.example .env
```

### Run the Application with Docker Compose
```sh
$ docker-compose up --build
```

## API Endpoints

### Authentication Routes
#### Signup
```http
POST /signup
```
**Request Body:**
```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "securepassword",
  "subDomain": "johndoe"
}
```

#### Login
```http
POST /login
```
**Request Body:**
```json
{
  "email": "johndoe@example.com",
  "password": "securepassword"
}
```
**Response:**
```json
{
  "token": "your.jwt.token"
}
```

#### Validate Token (Protected Route)
```http
GET /validate
Headers:
Cookie: Authorization=your.jwt.token
```

**Response:**
```json
{
  "david.dawoodworld.com": "user data"
}
```

