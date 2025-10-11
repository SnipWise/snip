<!-- METADATA
{
  "title": "Swift Functions",
  "tags": [
    "swift",
    "functions",
    "io"
  ],
  "language": "swift"
}
-->

## Functions
Function declaration and return values
```swift
import Foundation

func add(_ a: Int, _ b: Int) -> Int {
    return a + b
}

enum DivisionError: Error {
    case divisionByZero
}

func divide(_ a: Double, _ b: Double) throws -> Double {
    if b == 0 {
        throw DivisionError.divisionByZero
    }
    return a / b
}

let sum = add(5, 3)
do {
    let result = try divide(10, 2)
    print("Sum: \(sum), Division: \(result)")
} catch {
    print("Error: \(error)")
}
```
