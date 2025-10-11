<!-- METADATA
{
  "title": "Swift Async Await",
  "tags": [
    "swift",
    "concurrency"
  ],
  "language": "swift"
}
-->

## Async/Await
Concurrent programming with async/await
```swift
import Foundation

func worker(_ id: Int) async {
    print("Worker \(id) starting")
    try? await Task.sleep(nanoseconds: 100_000_000)
    print("Worker \(id) done")
}

Task {
    await withTaskGroup(of: Void.self) { group in
        for i in 1...3 {
            group.addTask {
                await worker(i)
            }
        }
    }
}

// Keep program alive for demo
try? await Task.sleep(nanoseconds: 500_000_000)
```
