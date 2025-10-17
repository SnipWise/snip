<!-- METADATA
{
  "title": "Working with Cookies",
  "tags": [
    "curl",
    "cookies",
    "session",
    "http"
  ],
  "language": "bash"
}
-->

## Working with Cookies
Send a cookie with request
```bash
curl -b "session_id=abc123" https://api.example.com/profile
```

Save cookies to file:
```bash
curl -c cookies.txt https://api.example.com/login \
  -d "username=john&password=secret"
```

Use saved cookies:
```bash
curl -b cookies.txt https://api.example.com/protected
```

Combined save and send cookies:
```bash
curl -b cookies.txt -c cookies.txt https://api.example.com/data
```
