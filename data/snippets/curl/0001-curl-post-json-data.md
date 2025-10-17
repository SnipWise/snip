<!-- METADATA
{
  "title": "POST Request with JSON Data",
  "tags": [
    "curl",
    "http",
    "post",
    "json",
    "api"
  ],
  "language": "bash"
}
-->

## POST Request with JSON Data
Send JSON data in a POST request
```bash
curl -X POST https://api.example.com/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

From a file:
```bash
curl -X POST https://api.example.com/users \
  -H "Content-Type: application/json" \
  -d @data.json
```
