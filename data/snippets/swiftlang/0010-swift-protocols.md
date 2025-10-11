<!-- METADATA
{
  "title": "Swift Protocols",
  "tags": [
    "swift",
    "protocols"
  ],
  "language": "swift"
}
-->

## Protocols
Defining and implementing protocols
```swift
import Foundation

protocol Shape {
    func area() -> Double
}

struct Rectangle: Shape {
    let width: Double
    let height: Double

    func area() -> Double {
        return width * height
    }
}

let shape: Shape = Rectangle(width: 5, height: 3)
print("Area: \(String(format: "%.2f", shape.area()))")
```
