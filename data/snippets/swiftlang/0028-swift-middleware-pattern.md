<!-- METADATA
{
  "title": "Swift Middleware Pattern",
  "tags": [
    "swift",
    "patterns"
  ],
  "language": "swift"
}
-->

## Middleware Pattern
HTTP middleware implementation (using Vapor)
```swift
import Vapor

struct LoggingMiddleware: AsyncMiddleware {
    func respond(to request: Request, chainingTo next: AsyncResponder) async throws -> Response {
        let start = Date()
        request.logger.info("Started \(request.method) \(request.url.path)")

        let response = try await next.respond(to: request)

        let duration = Date().timeIntervalSince(start)
        request.logger.info("Completed in \(duration)s")

        return response
    }
}

let app = Application()
defer { app.shutdown() }

// Register middleware
app.middleware.use(LoggingMiddleware())

app.get { req in
    return "Hello, World!"
}

try app.run()
```
