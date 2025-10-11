<!-- METADATA
{
  "title": "Swift Optionals",
  "tags": [
    "swift",
    "optionals"
  ],
  "language": "swift"
}
-->

## Optionals
Working with optionals
```swift
import Foundation

var value: Int? = 42

// Optional binding
if let unwrapped = value {
    print("Value: \(unwrapped)")
}

// Nil coalescing
let defaultValue = value ?? 0
print("Value with default: \(defaultValue)")

// Force unwrapping (use with caution)
if value != nil {
    print("Force unwrapped: \(value!)")
}
```
