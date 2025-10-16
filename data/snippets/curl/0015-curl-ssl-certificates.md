<!-- METADATA
{
  "title": "SSL Certificates",
  "tags": [
    "curl",
    "ssl",
    "tls",
    "security"
  ],
  "language": "bash"
}
-->

## SSL Certificates
Ignore SSL certificate validation (insecure)
```bash
curl -k https://self-signed.example.com/api
```

Use custom CA certificate:
```bash
curl --cacert /path/to/ca-cert.pem https://api.example.com/data
```

Use client certificate:
```bash
curl --cert client-cert.pem --key client-key.pem https://api.example.com/data
```

Show SSL certificate info:
```bash
curl -v --insecure https://example.com 2>&1 | grep -A 10 "Server certificate"
```
