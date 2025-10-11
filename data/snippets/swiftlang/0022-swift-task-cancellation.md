<!-- METADATA
{
  "title": "Swift Task Cancellation",
  "tags": [
    "swift",
    "concurrency"
  ],
  "language": "swift"
}
-->

## Task Cancellation
Using task cancellation for timeouts
```swift
import Foundation

func longTask() async throws {
    for i in 0..<5 {
        if Task.isCancelled {
            print("Task cancelled")
            throw CancellationError()
        }
        print("Working... \(i)")
        try await Task.sleep(nanoseconds: 500_000_000)
    }
    print("Task completed")
}

Task {
    let task = Task {
        try await longTask()
    }

    // Cancel after 1 second
    try? await Task.sleep(nanoseconds: 1_000_000_000)
    task.cancel()
}

// Keep program alive for demo
try? await Task.sleep(nanoseconds: 3_000_000_000)
```
