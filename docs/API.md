# SNIP API Documentation

## Overview

SNIP is a streaming chat completion API with vector similarity search capabilities. It provides conversational AI with context retrieval from a vector store.

**Base URL:** `http://localhost:3500` (default)

---

## Endpoints

### 1. POST /completion

Sends a message and receives a streaming chat completion response with relevant context from similarity search.

**Request:**
```bash
curl -X POST http://localhost:3500/completion \
  -H "Content-Type: application/json" \
  -d '{
    "message": "How do I read a file in Go?"
  }'
```

**Request Body:**
```json
{
  "message": "string"
}
```

**Response:**
- Streaming text response
- Final JSON response:
```json
{
  "response": "string"
}
```

**Behavior:**
- Performs similarity search against the vector store
- Adds relevant context to the conversation
- Streams the LLM response in real-time
- Updates conversation history
- Stores similarity search results (accessible via `/similarities`)

---

### 2. GET /similarities

Retrieves the similarity search results from the most recent completion request.

**Request:**
```bash
curl -X GET http://localhost:3500/similarities
```

**Response:**
```json
{
  "status": "ok",
  "data": {
    "user_message": "How do I read a file in Go?",
    "count": 3,
    "results": [
      {
        "id": "snippet-id-123",
        "similarity": 0.85,
        "content": "Example code snippet content..."
      }
    ]
  }
}
```

**Fields:**
- `user_message`: The user's query from the last completion
- `count`: Number of similar documents found
- `results`: Array of similarity results
  - `id`: Unique identifier of the snippet
  - `similarity`: Cosine similarity score (0.0 to 1.0)
  - `content`: Text content of the similar document

---

### 3. POST /operation/validate

Validates a pending operation and allows the stream to continue.

**Request:**
```bash
curl -X POST http://localhost:3500/operation/validate \
  -H "Content-Type: application/json" \
  -d '{
    "operation_id": "op_0x14000102300"
  }'
```

**Request Body:**
```json
{
  "operation_id": "string"
}
```

**Response:**
```json
{
  "status": "ok",
  "operation_id": "op_0x14000102300",
  "message": "Operation validated and continued"
}
```

**Behavior:**
- Marks the operation as `validated`
- Allows the streaming completion to continue
- Used when you want to approve and proceed with the operation

---

### 4. POST /operation/cancel

Cancels a pending operation but allows the stream to continue.

**Request:**
```bash
curl -X POST http://localhost:3500/operation/cancel \
  -H "Content-Type: application/json" \
  -d '{
    "operation_id": "op_0x14000102300"
  }'
```

**Request Body:**
```json
{
  "operation_id": "string"
}
```

**Response:**
```json
{
  "status": "ok",
  "operation_id": "op_0x14000102300",
  "message": "Operation cancelled but stream continues"
}
```

**Behavior:**
- Marks the operation as `cancelled`
- Allows the streaming completion to continue despite cancellation
- Used when you want to reject the operation but still see the response

---

### 5. POST /operation/reset

Resets a pending operation and stops the stream completely.

**Request:**
```bash
curl -X POST http://localhost:3500/operation/reset \
  -H "Content-Type: application/json" \
  -d '{
    "operation_id": "op_0x14000102300"
  }'
```

**Request Body:**
```json
{
  "operation_id": "string"
}
```

**Response:**
```json
{
  "status": "ok",
  "operation_id": "op_0x14000102300",
  "message": "Operation reset and stopped"
}
```

**Behavior:**
- Marks the operation as `reset`
- Stops the streaming completion immediately
- Used when you want to completely abort the operation

---

### 6. POST /completion/stop

Stops all active completion requests.

**Request:**
```bash
curl -X POST http://localhost:3500/completion/stop
```

**Response:**
```json
{
  "status": "ok",
  "stopped_completions": 2,
  "message": "Stopped 2 active completions"
}
```

---

### 7. POST /memory/reset

Resets the conversation memory, keeping only the system message.

**Request:**
```bash
curl -X POST http://localhost:3500/memory/reset
```

**Response:**
```json
{
  "status": "ok",
  "message": "Memory reset successfully"
}
```

---

### 8. GET /memory/messages/list

Retrieves the current conversation history.

**Request:**
```bash
curl -X GET http://localhost:3500/memory/messages/list
```

**Response:**
```json
{
  "status": "ok",
  "messages": [
    {
      "role": "system",
      "content": "You are a helpful AI assistant."
    },
    {
      "role": "user",
      "content": "Hello"
    }
  ],
  "count": 2
}
```

---

### 9. GET /memory/messages/tokens

Calculates and returns the total token count of the conversation history.

**Request:**
```bash
curl -X GET http://localhost:3500/memory/messages/tokens
```

**Response:**
```json
{
  "status": "ok",
  "tokens": 1523,
  "count": 12,
  "limit": 3000
}
```

**Fields:**
- `tokens`: Total token count of all messages
- `count`: Number of messages in memory
- `limit`: Context size limit (from `CONTEXT_SIZE_LIMIT` env var)

---

### 10. GET /models

Returns information about the configured models.

**Request:**
```bash
curl -X GET http://localhost:3500/models
```

**Response:**
```json
{
  "status": "ok",
  "chat_model": "hf.co/menlo/jan-nano-gguf:q4_k_m",
  "embeddings_model": "ai/mxbai-embed-large:latest",
  "tools_model": "hf.co/menlo/jan-nano-gguf:q4_k_m"
}
```

---

### 11. GET /health

Health check endpoint.

**Request:**
```bash
curl -X GET http://localhost:3500/health
```

**Response:**
```json
{
  "status": "ok"
}
```

---

## Configuration

The API can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `HTTP_PORT` | `3500` | HTTP server port |
| `MODEL_RUNNER_BASE_URL` | `http://localhost:12434/engines/llama.cpp/v1` | LLM API base URL |
| `SNIP_MODEL` | `hf.co/menlo/jan-nano-gguf:q4_k_m` | Chat completion model |
| `TOOLS_MODEL` | `hf.co/menlo/jan-nano-gguf:q4_k_m` | Model for tool calls detection and execution |
| `EMBEDDING_MODEL` | `ai/mxbai-embed-large:latest` | Embeddings model for vector search |
| `SIMILARITY_THRESHOLD` | `0.5` | Minimum similarity score (0.0-1.0) |
| `SIMILARITY_MAX_RESULTS` | `3` | Maximum number of similar documents to retrieve |
| `CONTEXT_SIZE_LIMIT` | `3000` | Maximum tokens in conversation context |
| `SYSTEM_INSTRUCTION` | `You are a helpful AI assistant.` | System prompt |
| `MCP_SERVER_BASE_URL` | `http://localhost:9011` | MCP Gateway server base URL for tool integration |

---

## Operation Control Workflow

When a completion request is initiated, it enters a pending state and waits for user confirmation. You have three options:

### Option 1: Validate (approve and continue)
```bash
# Extract operation_id from the streaming response
./validate.sh op_0x14000102300
```

### Option 2: Cancel (reject but continue anyway)
```bash
# The stream continues but marked as cancelled
./cancel.sh op_0x14000102300
```

### Option 3: Reset (abort completely)
```bash
# Stops the stream immediately
./reset.sh op_0x14000102300
```

---

## Complete Workflow Example

```bash
# 1. Send a completion request (will pause and wait for confirmation)
curl -X POST http://localhost:3500/completion \
  -H "Content-Type: application/json" \
  -H "Accept: text/event-stream" \
  -d '{"message": "Show me file I/O examples"}'

# Response will include:
# data: {"message": "tool detected", "status": "pending", "operation_id": "op_0x14000102300"}

# 2. Validate the operation to continue
curl -X POST http://localhost:3500/operation/validate \
  -H "Content-Type: application/json" \
  -d '{"operation_id": "op_0x14000102300"}'

# 3. Get the similarity results
curl -X GET http://localhost:3500/similarities | jq '.'

# 4. Check conversation history
curl -X GET http://localhost:3500/memory/messages/list | jq '.'

# 5. Check token usage
curl -X GET http://localhost:3500/memory/messages/tokens | jq '.'

# 6. Reset memory when needed
curl -X POST http://localhost:3500/memory/reset
```

---

## Error Handling

All endpoints return appropriate HTTP status codes:
- `200 OK`: Successful request
- `400 Bad Request`: Invalid request format
- `500 Internal Server Error`: Server error

Error responses follow this format:
```json
{
  "error": "Error message description"
}
```
