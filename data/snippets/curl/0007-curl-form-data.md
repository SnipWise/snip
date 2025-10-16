<!-- METADATA
{
  "title": "Form Data Submission",
  "tags": [
    "curl",
    "http",
    "form",
    "post"
  ],
  "language": "bash"
}
-->

## Form Data Submission
Submit form data (application/x-www-form-urlencoded)
```bash
curl -X POST https://api.example.com/login \
  -d "username=john" \
  -d "password=secret123"
```

Multipart form data with file upload:
```bash
curl -X POST https://api.example.com/upload \
  -F "file=@document.pdf" \
  -F "description=Important document"
```

Multiple files:
```bash
curl -X POST https://api.example.com/upload \
  -F "files=@image1.jpg" \
  -F "files=@image2.jpg" \
  -F "user_id=123"
```
