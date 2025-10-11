<!-- METADATA
{
  "title": "Swift For Loops",
  "tags": [
    "swift",
    "loops"
  ],
  "language": "swift"
}
-->

## For Loops
Different types of for loops
```swift
import Foundation

// Basic for loop
for i in 0..<3 {
    print("i = \(i)")
}

// Iterate over array with enumerated
let names = ["Alice", "Bob"]
for (index, name) in names.enumerated() {
    print("\(index): \(name)")
}
```
