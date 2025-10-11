<!-- METADATA
{
  "title": "Swift Configuration Management",
  "tags": [
    "swift",
    "io",
    "configuration"
  ],
  "language": "swift"
}
-->

## Configuration Management
Managing application configuration
```swift
import Foundation

struct ServerConfig: Codable {
    let host: String
    let port: Int
}

struct DatabaseConfig: Codable {
    let host: String
    let username: String
}

struct Config: Codable {
    let server: ServerConfig
    let database: DatabaseConfig
}

func loadConfig(from filename: String) throws -> Config {
    let url = URL(fileURLWithPath: filename)
    let data = try Data(contentsOf: url)
    let decoder = JSONDecoder()
    return try decoder.decode(Config.self, from: data)
}

do {
    let config = try loadConfig(from: "config.json")
    print("Server: \(config.server.host):\(config.server.port)")
} catch {
    print("Error loading config: \(error)")
}
```
