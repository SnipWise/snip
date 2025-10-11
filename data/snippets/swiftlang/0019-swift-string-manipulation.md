<!-- METADATA
{
  "title": "Swift String Manipulation",
  "tags": [
    "swift",
    "io",
    "strings"
  ],
  "language": "swift"
}
-->

## String Manipulation
Common string operations
```swift
import Foundation

let text = "  Hello, World!  "

print("Trimmed: '\(text.trimmingCharacters(in: .whitespaces))'")
print("Upper: \(text.uppercased())")
print("Contains 'World': \(text.contains("World"))")

let words = text.trimmingCharacters(in: .whitespaces).split(separator: " ")
print("Words: \(words)")

// String conversion
if let num = Int("123") {
    print("String to int: \(num)")
}
```
