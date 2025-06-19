# Go Projects Monorepo

This repository contains multiple Go projects and tools, each in its own directory. Below is a brief overview of each project.

## Projects

### 1. docker-test
A simple Go web server with Docker support.

- **Run locally:**  
  `go run main.go`
- **Build Docker image:**  
  `docker build -t docker-test .`
- **Run Docker container:**  
  `docker run -p 8081:8081 docker-test`

### 2. DSA go
A collection of Data Structures and Algorithms solutions in Go.

- **Run a solution:**  
  `go run greater_Average.go`

### 3. email-checker-tool
A CLI tool to check email domain DNS records (MX, SPF, DMARC).

- **Run:**  
  `go run main.go`
- **Usage:**  
  Enter domain names line by line in the terminal.

### 4. go-bookstore
A RESTful API for managing books using Go, Gorilla Mux, and GORM.

- **Run:**  
  `go run cmd/main/main.go`
- **API Endpoints:**  
  - `POST /book/` - Create a book  
  - `GET /book/` - List all books  
  - `GET /book/{bookId}` - Get a book  
  - `PUT /book/{bookId}` - Update a book  
  - `DELETE /book/{bookId}` - Delete a book

### 5. go-movies-crud
A simple CRUD API for movies using Go and Gorilla Mux.

- **Run:**  
  `go run main.go`
- **API Endpoints:**  
  - `GET /movies`  
  - `GET /movies/{id}`  
  - `POST /movies`  
  - `PUT /movies/{id}`  
  - `DELETE /movies/{id}`

### 6. go-server
A basic Go web server serving static files and handling forms.

- **Run:**  
  `go run main.go`
- **Access:**  
  - Static files at `/`
  - Form at `/form`
  - Hello endpoint at `/hello`

### 7. lambda-aws
A sample AWS Lambda function written in Go.

- **Handler:**  
  `main.go` with `HandleLambdaEvent`
- **Deploy:**  
  Build and upload `main.exe` or `function.zip` to AWS Lambda.

### 8. shorten-url-fiber-redis
A URL shortener API using Go, Fiber, Redis, and Docker Compose.

- **Run with Docker Compose:**  
  `docker-compose up --build`
- **API Endpoints:**  
  - `POST /api/v1` - Shorten a URL  
  - `GET /:url` - Redirect to original URL

### 9. slack-age-bot
A Slack bot that calculates age from year of birth using the Slacker framework.

- **Run:**  
  `go run main.go`
- **Configure:**  
  Set `SLACK_BOT_TOKEN` and `SLACK_APP_TOKEN` in `.env`

### 10. slack-file-bot
A Slack bot to upload files to a channel.

- **Run:**  
  `go run main.go`
- **Configure:**  
  Set `SLACK_BOT_TOKEN` and `CHANNEL_ID` in `.env`

---

## Requirements

- Go 1.22+ (some projects use Go 1.24)
- Docker (for Dockerized projects)
- Redis (for URL shortener)
- MySQL (for bookstore API)

## License

This repository is for educational and demonstration purposes.
