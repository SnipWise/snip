<!-- METADATA
{
  "title": "Swift Task Groups",
  "tags": [
    "swift",
    "concurrency"
  ],
  "language": "swift"
}
-->

## Task Groups
Concurrent task management with task groups
```swift
import Foundation

func producer(_ value: Int) async -> Int {
    print("Sent: \(value)")
    return value
}

Task {
    await withTaskGroup(of: Int.self) { group in
        for i in 1...3 {
            group.addTask {
                await producer(i)
            }
        }

        for await value in group {
            print("Received: \(value)")
        }
    }
}

// Keep program alive for demo
try? await Task.sleep(nanoseconds: 500_000_000)
```
