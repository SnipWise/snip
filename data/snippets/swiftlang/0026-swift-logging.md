<!-- METADATA
{
  "title": "Swift Logging",
  "tags": [
    "swift",
    "logging"
  ],
  "language": "swift"
}
-->

## Logging
Different logging approaches
```swift
import Foundation
import os.log

// Basic print logging
print("Standard log message")
print("User \("Alice") logged in")

// OSLog (unified logging system)
let logger = Logger(subsystem: "com.example.app", category: "general")

logger.info("User action")
logger.debug("Debug information: userId=\(123)")
logger.warning("Warning: high memory usage")
logger.error("Error: connection timeout")

// Custom log with metadata
logger.log(level: .info, "User \("Alice", privacy: .public) performed action")
```
