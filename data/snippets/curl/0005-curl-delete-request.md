<!-- METADATA
{
  "title": "DELETE Request",
  "tags": [
    "curl",
    "http",
    "delete"
  ],
  "language": "bash"
}
-->

## DELETE Request
Delete a resource
```bash
curl -X DELETE https://api.example.com/users/123
```

With authentication:
```bash
curl -X DELETE https://api.example.com/users/123 \
  -H "Authorization: Bearer token123"
```

With confirmation response:
```bash
curl -X DELETE https://api.example.com/users/123 \
  -H "Authorization: Bearer token123" \
  -w "\nHTTP Status: %{http_code}\n"
```
