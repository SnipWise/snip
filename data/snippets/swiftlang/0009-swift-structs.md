<!-- METADATA
{
  "title": "Swift Structs",
  "tags": [
    "swift",
    "structs"
  ],
  "language": "swift"
}
-->

## Structs
Defining and using structs
```swift
import Foundation

struct Person {
    let name: String
    let age: Int

    func greet() -> String {
        return "Hello, I'm \(name)"
    }
}

let person = Person(name: "Alice", age: 30)
print("Person: \(person)")
print(person.greet())
```
