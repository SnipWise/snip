<!-- METADATA
{
  "title": "Query Parameters",
  "tags": [
    "curl",
    "http",
    "query",
    "parameters"
  ],
  "language": "bash"
}
-->

## Query Parameters
Pass query parameters in URL
```bash
curl "https://api.example.com/search?q=golang&limit=10"
```

Multiple parameters:
```bash
curl "https://api.example.com/users?page=2&per_page=20&sort=created"
```

URL encoded parameters:
```bash
curl -G https://api.example.com/search \
  --data-urlencode "q=hello world" \
  --data-urlencode "category=tech"
```
