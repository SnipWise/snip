<!-- METADATA
{
  "title": "Swift Dictionaries",
  "tags": [
    "swift",
    "dictionaries"
  ],
  "language": "swift"
}
-->

## Dictionaries
Creating and manipulating dictionaries
```swift
import Foundation

var fruits = ["orange": 10, "grape": 15]
fruits["apple"] = 5

if let value = fruits["apple"] {
    print("Apple exists: true, value: \(value)")
}

for (key, value) in fruits {
    print("\(key): \(value)")
}
```
