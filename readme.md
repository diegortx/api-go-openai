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

