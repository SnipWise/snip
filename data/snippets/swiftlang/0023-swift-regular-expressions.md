<!-- METADATA
{
  "title": "Swift Regular Expressions",
  "tags": [
    "swift",
    "io",
    "regex"
  ],
  "language": "swift"
}
-->

## Regular Expressions
Pattern matching with regex
```swift
import Foundation

let emailPattern = #"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$"#
let emails = ["user@example.com", "invalid.email", "test@domain.co.uk"]

for email in emails {
    let isValid = email.range(of: emailPattern, options: .regularExpression) != nil
    print("\(email): \(isValid)")
}

let text = "Contact us at info@company.com"
if let range = text.range(of: emailPattern, options: .regularExpression) {
    let foundEmail = String(text[range])
    print("Found email: \(foundEmail)")
}
```
