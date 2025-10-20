<!-- METADATA
{
  "title": "Crystal HTTP Client",
  "tags": [
    "crystal",
    "http",
    "client"
  ],
  "language": "crystal"
}
-->

## HTTP Client
Making HTTP requests
```crystal
require "http/client"

# Simple GET request
response = HTTP::Client.get("https://api.example.com/data")
puts response.status_code
puts response.body

# GET with headers
headers = HTTP::Headers{
  "Authorization" => "Bearer token123",
  "Content-Type"  => "application/json"
}
response = HTTP::Client.get("https://api.example.com/data", headers: headers)

# POST request
body = %({"name": "Alice", "age": 30})
response = HTTP::Client.post(
  "https://api.example.com/users",
  headers: HTTP::Headers{"Content-Type" => "application/json"},
  body: body
)

# Using client instance
client = HTTP::Client.new("api.example.com", tls: true)
response = client.get("/data")
puts response.body
client.close
```
