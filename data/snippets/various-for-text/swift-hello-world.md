<!-- METADATA
{
  "title": "Hello World",
  "tags": [
    "swift",
    "basics",
    "control-flow"
  ],
  "language": "swift"
}
-->

# Swift Hello World

## Basic Hello World Program

```swift
print("Hello, World!")
```

That's it! Swift doesn't require a main function or class declaration for simple scripts.

## As a Function

```swift
func sayHello() {
    print("Hello, World!")
}

sayHello()
```

## In a Swift Application (iOS/macOS)

```swift
import SwiftUI

@main
struct HelloWorldApp: App {
    var body: some Scene {
        WindowGroup {
            Text("Hello, World!")
        }
    }
}
```

## Steps to Run

1. **Create the file**: Save the code as `hello.swift`

2. **Run directly**:
   ```bash
   swift hello.swift
   ```

3. **Or compile first, then run**:
   ```bash
   swiftc hello.swift
   ./hello
   ```

## Using Swift REPL

You can also use the interactive Swift REPL:
```bash
swift
```
Then type:
```swift
print("Hello, World!")
```

## Explanation

- `print()`: Outputs text to the console with a newline
- No semicolons required
- No main function needed for simple scripts
- `@main`: Marks the entry point for Swift applications
