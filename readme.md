# Go Chat API

A RESTful API service that provides chat functionality using OpenAI's GPT models.

## Features

- Chat endpoint for message processing with OpenAI integration
- Conversation management with history tracking
- Built-in request validation and comprehensive error handling
- Cross-Origin Resource Sharing (CORS) support for web clients
- Detailed request logging with timing information
- Health check endpoint with timestamp
- Thread-safe conversation storage using mutex locks
- Maximum message history limit to optimize context
- Clear conversation functionality

## API Endpoints

### POST /api/v1/chat
Process a chat message and get an AI-generated response.

**Request Body:**

```json
{
"user_id": "user123",
"message": "What is the capital of France?",
"context": {
"system_prompt": "You are a helpful assistant that specializes in geography.",
"metadata": {
"language": "en",
"topic": "geography"
},
"temperature": 0.7
}
}
```

**Response:**

```json
{
"success": true,
"response": "The capital of France is Paris."
}
```

### DELETE /api/v1/chat/{userID}
Clear the conversation history for a specific user.

**Response:**

```json
{
"success": true,
"response": "Conversation cleared successfully"
}
```

### GET /health
Check if the service is running.

**Response:**

```json
{
"status": "ok",
"time": "2024-03-21T10:00:00Z"
}
```

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd go-api
```

2. Copy the environment example file and configure your settings:
```bash
cp .env.example .env
# Edit .env with your OpenAI API key
```

3. Install dependencies:
```bash
go mod tidy
```

4. Run the application:
```bash
# Development with hot reload
air

# Or standard run
go run main.go
```

## Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| PORT | Server port | No | 8080 |
| OPENAI_API_KEY | OpenAI API key | Yes | - |

## Project Structure

```
.
├── config/         # Configuration management
├── controllers/    # Request handlers
├── models/         # Data structures
├── routes/         # Route definitions
├── services/       # Business logic
├── .env           # Environment configuration
├── .env.example   # Environment example
├── main.go        # Application entry point
└── README.md      # Documentation
```

## API Request Examples

### Send a Chat Message
```bash
curl -X POST http://localhost:8080/api/v1/chat \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "message": "What is the capital of France?",
    "context": {
        "system_prompt": "You are a helpful assistant.",
        "temperature": 0.7
    }
  }'
```

### Clear Chat History
```bash
curl -X DELETE http://localhost:8080/api/v1/chat/user123
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:

json
{
    "success": false,
    "error": "Error message here"
}

Common status codes:
- 400: Bad Request (invalid input)
- 401: Unauthorized (invalid API key)
- 500: Internal Server Error

## Development

### Hot Reload
The project uses Air for hot reloading during development:

1. Install Air:
```bash
go install github.com/cosmtrek/air@latest
```

2. Run with hot reload:
```bash
air
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [OpenAI API](https://platform.openai.com/docs/api-reference)

