<!-- METADATA
{
  "title": "PUT Request",
  "tags": [
    "curl",
    "http",
    "put",
    "update"
  ],
  "language": "bash"
}
-->

## PUT Request
Update a resource with PUT
```bash
curl -X PUT https://api.example.com/users/123 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

With authentication:
```bash
curl -X PUT https://api.example.com/users/123 \
  -H "Authorization: Bearer token123" \
  -H "Content-Type: application/json" \
  -d @update.json
```
