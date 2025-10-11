<!-- METADATA
{
  "title": "Swift Command Line Arguments",
  "tags": [
    "swift",
    "cli"
  ],
  "language": "swift"
}
-->

## Command Line Arguments
Parsing command line arguments
```swift
import Foundation

// Using CommandLine
let arguments = CommandLine.arguments

print("Program name: \(arguments[0])")

if arguments.count > 1 {
    print("Arguments:")
    for (index, arg) in arguments.dropFirst().enumerated() {
        print("  \(index): \(arg)")
    }
}

// Manual parsing example
var name = "World"
var verbose = false

for (i, arg) in arguments.enumerated() {
    if arg == "--name" && i + 1 < arguments.count {
        name = arguments[i + 1]
    } else if arg == "--verbose" {
        verbose = true
    }
}

if verbose {
    print("Verbose mode enabled")
}

print("Hello, \(name)!")
```
