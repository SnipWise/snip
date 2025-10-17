<!-- METADATA
{
  "title": "Silent Mode and Progress",
  "tags": [
    "curl",
    "output",
    "progress"
  ],
  "language": "bash"
}
-->

## Silent Mode and Progress
Silent mode (no progress bar)
```bash
curl -s https://api.example.com/data
```

Show only errors:
```bash
curl -sS https://api.example.com/data
```

Progress bar:
```bash
curl --progress-bar -o file.zip https://example.com/large-file.zip
```

Custom output format:
```bash
curl -w "\nTime: %{time_total}s\nSize: %{size_download} bytes\n" \
  https://api.example.com/data
```
