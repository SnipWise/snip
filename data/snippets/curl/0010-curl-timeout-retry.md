<!-- METADATA
{
  "title": "Timeout and Retry",
  "tags": [
    "curl",
    "timeout",
    "retry",
    "reliability"
  ],
  "language": "bash"
}
-->

## Timeout and Retry
Set connection timeout (seconds)
```bash
curl --connect-timeout 10 https://api.example.com/data
```

Maximum time for the operation:
```bash
curl --max-time 30 https://api.example.com/slow-endpoint
```

Retry on failure:
```bash
curl --retry 3 --retry-delay 2 https://api.example.com/data
```

Combined timeout and retry:
```bash
curl --connect-timeout 5 \
     --max-time 30 \
     --retry 3 \
     --retry-delay 1 \
     https://api.example.com/data
```
