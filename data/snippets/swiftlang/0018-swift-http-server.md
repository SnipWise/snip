<!-- METADATA
{
  "title": "Swift Http Server",
  "tags": [
    "swift",
    "http",
    "networking"
  ],
  "language": "swift"
}
-->

## HTTP Server
Creating a simple HTTP server (using Vapor framework)
```swift
import Vapor

let app = Application()
defer { app.shutdown() }

// Home route
app.get { req in
    return "Hello, World!"
}

// JSON route
app.get("json") { req in
    return ["message": "Hello, JSON!"]
}

print("Server starting on :8080")
try app.run()
```
