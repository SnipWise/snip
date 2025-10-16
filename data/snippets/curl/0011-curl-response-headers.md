<!-- METADATA
{
  "title": "View Response Headers",
  "tags": [
    "curl",
    "http",
    "headers",
    "debugging"
  ],
  "language": "bash"
}
-->

## View Response Headers
Show only response headers
```bash
curl -I https://api.example.com/data
```

Include headers in output:
```bash
curl -i https://api.example.com/data
```

Save headers to file:
```bash
curl -D headers.txt https://api.example.com/data
```

Show HTTP status code:
```bash
curl -s -o /dev/null -w "%{http_code}" https://api.example.com/data
```
