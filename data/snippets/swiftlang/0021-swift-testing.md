<!-- METADATA
{
  "title": "Swift Testing",
  "tags": [
    "swift",
    "testing"
  ],
  "language": "swift"
}
-->

## Testing
Unit tests with XCTest
```swift
// Math.swift
func add(_ a: Int, _ b: Int) -> Int {
    return a + b
}

// MathTests.swift
import XCTest

class MathTests: XCTestCase {
    func testAdd() {
        let tests = [
            (a: 2, b: 3, expected: 5),
            (a: 0, b: 0, expected: 0),
            (a: -1, b: 1, expected: 0),
        ]

        for test in tests {
            let result = add(test.a, test.b)
            XCTAssertEqual(result, test.expected,
                "add(\(test.a), \(test.b)) = \(result); want \(test.expected)")
        }
    }
}
```
