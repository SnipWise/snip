<!-- METADATA
{
  "title": "Swift Json Handling",
  "tags": [
    "swift",
    "json"
  ],
  "language": "swift"
}
-->

## JSON Handling
Encode and decode JSON data
```swift
import Foundation

struct Person: Codable {
    let name: String
    let age: Int
    let email: String?
}

let person = Person(name: "Alice", age: 30, email: "alice@example.com")

do {
    // Encode to JSON
    let encoder = JSONEncoder()
    encoder.outputFormatting = .prettyPrinted
    let jsonData = try encoder.encode(person)
    print("JSON: \(String(data: jsonData, encoding: .utf8) ?? "")")

    // Decode from JSON
    let decoder = JSONDecoder()
    let newPerson = try decoder.decode(Person.self, from: jsonData)
    print("Decoded: \(newPerson)")
} catch {
    print("Error: \(error)")
}
```
