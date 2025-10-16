<!-- METADATA
{
  "title": "Custom Headers",
  "tags": [
    "curl",
    "http",
    "headers",
    "authentication"
  ],
  "language": "bash"
}
-->

## Custom Headers
Add custom headers to your request
```bash
curl https://api.example.com/data \
  -H "Authorization: Bearer token123" \
  -H "Accept: application/json" \
  -H "User-Agent: MyApp/1.0"
```

With API key:
```bash
curl https://api.example.com/data \
  -H "X-API-Key: your-api-key-here"
```
