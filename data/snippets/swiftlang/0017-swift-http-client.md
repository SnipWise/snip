<!-- METADATA
{
  "title": "Swift Http Client",
  "tags": [
    "swift",
    "http",
    "networking"
  ],
  "language": "swift"
}
-->

## HTTP Client
Making HTTP requests
```swift
import Foundation

Task {
    guard let url = URL(string: "https://httpbin.org/json") else { return }

    do {
        let (data, response) = try await URLSession.shared.data(from: url)

        if let httpResponse = response as? HTTPURLResponse {
            print("Status: \(httpResponse.statusCode)")
        }

        if let body = String(data: data, encoding: .utf8) {
            print("Response: \(body)")
        }
    } catch {
        print("Error: \(error)")
    }
}

// Keep program alive for demo
try? await Task.sleep(nanoseconds: 2_000_000_000)
```
