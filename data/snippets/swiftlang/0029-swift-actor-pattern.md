<!-- METADATA
{
  "title": "Swift Actor Pattern",
  "tags": [
    "swift",
    "patterns"
  ],
  "language": "swift"
}
-->

## Actor Pattern
Concurrent task processing with actors
```swift
import Foundation

struct Job {
    let id: Int
    let data: String
}

actor WorkerPool {
    private var results: [String] = []

    func processJob(_ job: Job, workerId: Int) async {
        print("Worker \(workerId) processing job \(job.id)")
        try? await Task.sleep(nanoseconds: 100_000_000)
        let result = "Job \(job.id) completed by worker \(workerId)"
        results.append(result)
    }

    func getResults() -> [String] {
        return results
    }
}

Task {
    let pool = WorkerPool()

    // Process jobs concurrently
    await withTaskGroup(of: Void.self) { group in
        for i in 1...5 {
            let job = Job(id: i, data: "task-\(i)")
            let workerId = (i % 3) + 1
            group.addTask {
                await pool.processJob(job, workerId: workerId)
            }
        }
    }

    // Get results
    let results = await pool.getResults()
    for result in results {
        print(result)
    }
}

// Keep program alive for demo
try? await Task.sleep(nanoseconds: 1_000_000_000)
```
