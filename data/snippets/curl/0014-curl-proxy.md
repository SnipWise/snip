<!-- METADATA
{
  "title": "Using Proxy",
  "tags": [
    "curl",
    "proxy",
    "network"
  ],
  "language": "bash"
}
-->

## Using Proxy
HTTP proxy
```bash
curl -x http://proxy.example.com:8080 https://api.example.com/data
```

Proxy with authentication:
```bash
curl -x http://user:pass@proxy.example.com:8080 https://api.example.com/data
```

SOCKS proxy:
```bash
curl --socks5 localhost:1080 https://api.example.com/data
```

Bypass proxy for specific host:
```bash
curl --noproxy localhost,127.0.0.1 https://api.example.com/data
```
