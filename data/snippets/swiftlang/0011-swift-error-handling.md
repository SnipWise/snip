<!-- METADATA
{
  "title": "Swift Error Handling",
  "tags": [
    "swift",
    "error-handling"
  ],
  "language": "swift"
}
-->

## Error Handling
Proper error handling patterns
```swift
import Foundation

enum DivisionError: Error {
    case divisionByZero
}

func safeDivide(_ a: Double, _ b: Double) throws -> Double {
    if b == 0 {
        throw DivisionError.divisionByZero
    }
    return a / b
}

do {
    let result = try safeDivide(10, 0)
    print("Result: \(result)")
} catch {
    print("Error: \(error)")
}
```
