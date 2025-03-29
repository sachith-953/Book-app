# Book API - Software Engineering Intern Assignment

## Overview
This is a robust REST API implemented in GoLang for managing a book collection. The application provides comprehensive book management capabilities with efficient search functionality and clear, structured endpoints for CRUD operations on books.

## Key Features
- CRUD operations for book management
- Concurrent search optimization
- JSON-based data persistence
- Docker and Kubernetes support
- Root endpoint for easy access to available routes and instructions

## Technologies Used
- GoLang (v1.21+)
- Docker
- Minikube

## Prerequisites
- Go (version 1.21 or higher)
- Docker (optional)
- Minikube (optional, for Kubernetes deployment)

## Project Structure
```
book-api/
│
├── cmd/
│   └── main.go
│
├── internal/
│   └── book/
│       ├── handler/
│       ├── service/
│       └── storage/
│
├── k8s/
│   ├── deployment.yaml
│   └── service.yaml
│
├── Dockerfile
└── go.mod
```

## Installation and Running

### 1. Local Development

#### Install Dependencies
```bash
go mod tidy
```

#### Run the Application
```bash
go run cmd/main.go
```
- API accessible at `http://localhost:8081`

### 2. Docker Deployment

#### Build Docker Image
```bash
docker build -t book-api .
```

#### Run Docker Container
```bash
docker run -d -p 8081:8081 --name book-api-container book-api
```

### 3. Kubernetes Deployment

#### Start Minikube
```bash
minikube start
```

#### Deploy to Kubernetes
```bash
kubectl apply -f k8s/
```

#### Access Kubernetes Service
```bash
minikube service book-api-service --url
```

## API Endpoints

### Root Endpoint
- `GET /`: Displays a welcome message and provides a list of available API endpoints.

When you visit the root URL (`http://localhost:8081/`), the server will return a message with information about all the available endpoints.

Example Response:
```
Welcome to the Book API!
Here are the available endpoints:
- GET /books         - Get a list of all books
- POST /books        - Create a new book
- GET /books/{id}    - Get a single book by ID
- PUT /books/{id}    - Update a book by ID
- DELETE /books/{id} - Delete a book by ID
- GET /search        - Search books by title/description (use query parameter ?q=your_search_term)

Example:
GET /search?q=great
```

### Book Operations
- `POST /books`: Create a new book
- `GET /books/{id}`: Retrieve a book by ID
- `PUT /books/{id}`: Update an existing book
- `DELETE /books/{id}`: Delete a book
- `GET /books/search?q=<keyword>`: Search books by title or description

## Search Optimization
The search functionality leverages:
- Goroutines for concurrent processing
- Channels for efficient result collection
- Optimized search across large datasets

### Sample Book Creation Request
```json
{
    "bookId": "bb329a31-6b1e-4daa-87ee-71631aa05866",
    "authorId": "e0d91f68-a183-477d-8aa4-1f44ccc78a70",
    "publisherId": "2f7b19e9-b268-4440-a15b-bed8177ed607",
    "title": "The Great Gatsby",
    "publicationDate": "1925-04-10",
    "isbn": "9780743273565",
    "pages": 180,
    "genre": "Novel",
    "description": "Set in the 1920s, this classic novel explores themes of wealth, love, and the American Dream.",
    "price": 15.99,
    "quantity": 5
}
```

## Testing the API

### Curl Examples

#### Create a Book
```bash
curl -X POST http://localhost:8081/books \
     -H "Content-Type: application/json" \
     -d '{
           "bookId": "unique-id",
           "title": "New Book",
           "price": 19.99,
           "quantity": 10
         }'
```

#### Search Books
```bash
curl "http://localhost:8081/books/search?q=novel"
```

#### Example for Searching Books by Keyword
To search books by a keyword (for example, "great"):

```bash
curl "http://localhost:8081/search?q=great"
```

This will return all books where the title or description contains the term "great."

## Troubleshooting
- Verify Go installation: `go version`
- Ensure port 8081 is available
- Check Docker and Kubernetes configurations
- Verify dependencies with `go mod tidy`

## Performance Considerations
- Concurrent search processing
- Efficient goroutine and channel management
- Lightweight JSON-based storage

## Notes
- Designed as a demonstration of software engineering skills
- Confidential assignment, not for public sharing

## Contact
sachithsr953@gmail.com