<!-- METADATA
{
  "title": "Swift If Else Statements",
  "tags": [
    "swift",
    "control-flow"
  ],
  "language": "swift"
}
-->

## If-Else Statements
Conditional statements
```swift
import Foundation

let age = 25
if age >= 18 {
    print("Adult")
} else {
    print("Minor")
}

// If with optional binding
let num: Int? = 42
if let num = num, num > 40 {
    print("Number \(num) > 40")
}
```
