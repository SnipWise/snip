<!-- METADATA
{
  "title": "Swift Date Operations",
  "tags": [
    "swift",
    "io",
    "date"
  ],
  "language": "swift"
}
-->

## Date Operations
Working with date and time
```swift
import Foundation

let now = Date()
print("Current time: \(now)")

let formatter = DateFormatter()
formatter.dateFormat = "yyyy-MM-dd HH:mm:ss"
print("Formatted: \(formatter.string(from: now))")

// Add time
if let tomorrow = Calendar.current.date(byAdding: .day, value: 1, to: now) {
    print("Tomorrow: \(tomorrow)")
}

// Parse date string
let dateStr = "2023-12-25 10:30:00"
if let parsed = formatter.date(from: dateStr) {
    print("Parsed: \(parsed)")
}
```
