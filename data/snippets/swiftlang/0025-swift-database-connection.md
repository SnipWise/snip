<!-- METADATA
{
  "title": "Swift Database Connection",
  "tags": [
    "swift",
    "io",
    "database"
  ],
  "language": "swift"
}
-->

## Database Connection
Basic database operations with PostgreSQL (using PostgresNIO)
```swift
import PostgresNIO

struct User {
    let id: Int
    let name: String
    let email: String
}

Task {
    do {
        let config = PostgresConnection.Configuration(
            host: "localhost",
            username: "user",
            password: "password",
            database: "test"
        )

        let connection = try await PostgresConnection.connect(
            configuration: config,
            id: 1,
            logger: .init(label: "postgres")
        )
        defer { try? await connection.close() }

        let rows = try await connection.query(
            "SELECT id, name, email FROM users WHERE id = $1",
            [1]
        )

        for try await row in rows {
            let id = try row.decode(Int.self, context: .default)
            let name = try row.decode(String.self, context: .default)
            let email = try row.decode(String.self, context: .default)
            print("User: id=\(id), name=\(name), email=\(email)")
        }
    } catch {
        print("Error: \(error)")
    }
}
```
