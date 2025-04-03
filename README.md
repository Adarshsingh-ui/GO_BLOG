


## Features

- **CRUD Operations**: Create, Read, Update, and Delete blog posts
- **Rich Blog Model**: 
  - Title and content
  - Author information
  - Tags for categorization
  - Read time estimation
  - Creation and update timestamps
  - Draft/Published state
- **API Versioning**: Using `/api/v1` prefix
- **Input Validation**: Built-in validation for all fields
- **Error Handling**: Consistent error responses
- **Logging**: Built-in request logging
- **UUID-based IDs**: Secure and unique identifiers

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Adarshsingh-ui/GO_BLOG.git
cd blog-application
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:3060`

## API Endpoints

### Blogs

#### Get All Blogs
```http
GET /api/v1/blogs
```

#### Get Single Blog
```http
GET /api/v1/blogs/:id
```

#### Create Blog
```http
POST /api/v1/blogs
Content-Type: application/json

{
    "title": "My New Blog",
    "content": "This is the content of my blog post...",
    "author": "Your Name",
    "tags": ["golang", "web", "api"]
}
```

#### Update Blog
```http
PUT /api/v1/blogs/:id
Content-Type: application/json

{
    "title": "Updated Title",
    "content": "Updated content...",
    "author": "Your Name",
    "tags": ["golang", "web", "api", "updated"]
}
```

#### Publish Blog
```http
PATCH /api/v1/blogs/:id/publish
```

#### Delete Blog
```http
DELETE /api/v1/blogs/:id
```

## Response Format

All responses follow a consistent format:

```json
{
    "success": true,
    "message": "Optional message",
    "data": {
        // Blog data here
    }
}
```

## Blog Model

```go
type Blog struct {
    ID          string    `json:"id" binding:"required"`
    Title       string    `json:"title" binding:"required,min=3,max=100"`
    Content     string    `json:"content" binding:"required,min=10"`
    Author      string    `json:"author" binding:"required"`
    Tags        []string  `json:"tags"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    ReadTime    int       `json:"read_time"`
    IsPublished bool      `json:"is_published"`
}
```

## Error Handling

The API returns appropriate HTTP status codes:
- 200: Success
- 201: Created
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [Go Programming Language](https://golang.org/)
