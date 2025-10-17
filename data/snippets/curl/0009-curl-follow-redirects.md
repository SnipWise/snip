<!-- METADATA
{
  "title": "Follow Redirects",
  "tags": [
    "curl",
    "http",
    "redirects"
  ],
  "language": "bash"
}
-->

## Follow Redirects
Follow HTTP redirects automatically
```bash
curl -L https://example.com/redirect
```

Limit redirect hops:
```bash
curl -L --max-redirs 5 https://example.com/redirect
```

Show redirect chain:
```bash
curl -L -v https://example.com/redirect
```
