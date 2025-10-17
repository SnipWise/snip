<!-- METADATA
{
  "title": "Basic Authentication",
  "tags": [
    "curl",
    "authentication",
    "security",
    "http"
  ],
  "language": "bash"
}
-->

## Basic Authentication
HTTP Basic Auth with username and password
```bash
curl -u username:password https://api.example.com/protected
```

Interactive password prompt:
```bash
curl -u username https://api.example.com/protected
```

With explicit Authorization header:
```bash
curl https://api.example.com/protected \
  -H "Authorization: Basic dXNlcm5hbWU6cGFzc3dvcmQ="
```
