<!-- METADATA
{
  "title": "Download Files",
  "tags": [
    "curl",
    "download",
    "files"
  ],
  "language": "bash"
}
-->

## Download Files
Save response to a file
```bash
curl -o output.json https://api.example.com/data
```

Use remote filename:
```bash
curl -O https://example.com/file.zip
```

Follow redirects and save:
```bash
curl -L -o file.tar.gz https://github.com/user/repo/releases/latest/download/file.tar.gz
```
