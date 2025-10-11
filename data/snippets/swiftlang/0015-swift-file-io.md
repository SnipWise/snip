<!-- METADATA
{
  "title": "Swift File Io",
  "tags": [
    "swift",
    "file-io",
    "io"
  ],
  "language": "swift"
}
-->

## File I/O
Reading and writing files
```swift
import Foundation

let content = "Hello, World!\n"
let fileURL = FileManager.default.temporaryDirectory.appendingPathComponent("test.txt")

do {
    // Write file
    try content.write(to: fileURL, atomically: true, encoding: .utf8)

    // Read file
    let data = try String(contentsOf: fileURL, encoding: .utf8)
    print("File content: \(data)")

    // Clean up
    try FileManager.default.removeItem(at: fileURL)
} catch {
    print("Error: \(error)")
}
```
