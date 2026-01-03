# Complete Swift Programming Course for Console/TUI Development

## Table of Contents

1. [Introduction to Swift](#introduction-to-swift)
2. [Getting Started](#getting-started)
3. [Basic Syntax](#basic-syntax)
4. [Variables and Data Types](#variables-and-data-types)
5. [Operators](#operators)
6. [Control Flow](#control-flow)
7. [Functions and Closures](#functions-and-closures)
8. [Collections](#collections)
9. [Strings](#strings)
10. [Optionals](#optionals)
11. [Enumerations](#enumerations)
12. [Structures and Classes](#structures-and-classes)
13. [Properties and Methods](#properties-and-methods)
14. [Protocols and Extensions](#protocols-and-extensions)
15. [Generics](#generics)
16. [Error Handling](#error-handling)
17. [File I/O and JSON](#file-io-and-json)
18. [Memory Management](#memory-management)
19. [Advanced Topics](#advanced-topics)
20. [Console/TUI Development](#console-tui-development)

---

## Introduction to Swift

Swift is a powerful and intuitive programming language created by Apple for iOS, macOS, watchOS, and tvOS development. It's also excellent for server-side development and command-line applications.

### Key Features

- **Modern and Safe**: Type-safe language with automatic memory management
- **Fast**: Compiled language with performance comparable to C++
- **Expressive**: Clean, readable syntax inspired by multiple programming languages
- **Open Source**: Available for Linux and other platforms
- **Powerful Type System**: With type inference and generics
- **Protocol-Oriented**: Emphasizes protocols over inheritance
- **Functional Programming**: First-class functions, closures, and map/filter/reduce

### Why Choose Swift for Console Applications?

- Fast compilation and execution
- Strong type safety prevents common errors
- Excellent for building CLI tools and TUI applications
- Modern concurrency features (async/await)
- Comprehensive standard library
- Great for scripting and automation

### Use Cases

- Command-line tools and utilities
- Terminal User Interface (TUI) applications
- Server-side applications
- Build tools and automation scripts
- Data processing and analysis
- Network utilities and clients

---

## Getting Started

### Installation

#### On macOS

Swift comes pre-installed with Xcode:
```bash
# Install Xcode from App Store or download from apple.com

# Or install just the command-line tools
xcode-select --install

# Verify installation
swift --version
```

#### On Linux (Ubuntu/Debian)

```bash
# Download Swift from swift.org
wget https://swift.org/builds/swift-5.9-release/ubuntu2204/swift-5.9-RELEASE/swift-5.9-RELEASE-ubuntu22.04.tar.gz

# Extract
tar xzf swift-5.9-RELEASE-ubuntu22.04.tar.gz

# Add to PATH
export PATH=/path/to/swift/usr/bin:$PATH

# Verify
swift --version
```

#### On Windows

Download and install from swift.org or use WSL with the Linux installation.

### Swift REPL (Interactive Mode)

```bash
# Start REPL
swift

# In REPL
1> print("Hello, Swift!")
Hello, Swift!
2> let x = 42
3> x * 2
$R0: Int = 84
4> :quit
```

### Creating Your First Swift Program

Create a file `hello.swift`:
```swift
print("Hello, World!")
```

Run it:
```bash
swift hello.swift
```

### Swift Package Manager

Create a new executable package:
```bash
# Create new package
swift package init --type executable --name MyApp

# Directory structure:
# MyApp/
# ├── Package.swift
# ├── Sources/
# │   └── MyApp/
# │       └── main.swift
# └── Tests/

# Build
swift build

# Run
swift run

# Build for release
swift build -c release
```

### Package.swift Example

```swift
// swift-tools-version:5.9
import PackageDescription

let package = Package(
    name: "MyApp",
    platforms: [
        .macOS(.v13)
    ],
    dependencies: [
        // Add dependencies here
        // .package(url: "https://github.com/...", from: "1.0.0"),
    ],
    targets: [
        .executableTarget(
            name: "MyApp",
            dependencies: []),
        .testTarget(
            name: "MyAppTests",
            dependencies: ["MyApp"]),
    ]
)
```

---

## Basic Syntax

### Your First Swift Program

```swift
// main.swift
import Foundation

print("Hello, World!")

// Multi-line output
print("""
This is a
multi-line
string
""")
```

### Comments

```swift
// Single line comment

/*
Multi-line
comment
*/

/// Documentation comment for the next declaration
/// - Parameter name: The person's name
func greet(name: String) {
    print("Hello, \(name)!")
}

/**
 Multi-line documentation comment

 This function performs addition.

 - Parameters:
   - a: First number
   - b: Second number
 - Returns: Sum of a and b
 */
func add(_ a: Int, _ b: Int) -> Int {
    return a + b
}
```

### Semicolons

```swift
// Semicolons are optional (usually not used)
let x = 10
let y = 20

// But can be used to put multiple statements on one line
let a = 1; let b = 2; let c = 3
```

### Print and String Interpolation

```swift
let name = "Alice"
let age = 30

// Basic print
print("Hello")

// String interpolation
print("Name: \(name), Age: \(age)")

// Multiple items
print("Name:", name, "Age:", age)

// Custom separator and terminator
print("One", "Two", "Three", separator: " | ", terminator: ".\n")

// Print without newline
print("Loading", terminator: "")
print("...", terminator: "")
print("Done!")
```

### Naming Conventions

```swift
// Variables and functions: camelCase
let userName = "Alice"
var itemCount = 0
func calculateTotal() { }

// Types (classes, structs, enums, protocols): PascalCase
struct Person { }
class BankAccount { }
enum Direction { }
protocol Drawable { }

// Constants: Can be camelCase or SCREAMING_SNAKE_CASE
let maxUsers = 100
let MAX_CONNECTIONS = 50

// Private properties: prefix with underscore (optional)
private var _internalState = 0
```

---

## Variables and Data Types

### Variables and Constants

```swift
// Constants (immutable) - preferred when possible
let name = "John"
let pi = 3.14159
// name = "Jane"  // Error: Cannot assign to value

// Variables (mutable)
var age = 30
var score = 0
age = 31  // OK
score += 10  // OK

// Type annotations (usually inferred)
let explicitString: String = "Hello"
let explicitInt: Int = 42
let explicitDouble: Double = 3.14

// Multiple declarations
let x = 1, y = 2, z = 3
var a = 10, b = 20, c = 30
```

### Integer Types

```swift
// Signed integers
let int8: Int8 = -128...127
let int16: Int16 = -32768...32767
let int32: Int32 = -2147483648...2147483647
let int64: Int64 = -9223372036854775808...9223372036854775807

// Unsigned integers
let uint8: UInt8 = 0...255
let uint16: UInt16 = 0...65535
let uint32: UInt32 = 0...4294967295
let uint64: UInt64 = 0...18446744073709551615

// Platform-dependent
let int: Int = 42      // Usually 64-bit on modern systems
let uint: UInt = 100   // Unsigned version

// Integer literals
let decimal = 17
let binary = 0b10001
let octal = 0o21
let hexadecimal = 0x11

// Underscores for readability
let million = 1_000_000
let bytes = 0xFF_00_FF
```

### Floating-Point Types

```swift
// Float: 32-bit floating point
let float: Float = 3.14

// Double: 64-bit floating point (preferred)
let double: Double = 3.14159265359

// Type inference defaults to Double
let inferredDouble = 3.14  // Double

// Scientific notation
let scientific = 1.25e2     // 125.0
let tiny = 1.25e-2          // 0.0125
```

### Boolean Type

```swift
let isActive = true
let isComplete: Bool = false

// Boolean operations
let and = true && false  // false
let or = true || false   // true
let not = !true          // false

// From comparisons
let isAdult = age >= 18
let hasAccess = isActive && isAdult
```

### String and Character Types

```swift
// String
let greeting = "Hello, Swift!"
let multiline = """
This is a
multi-line
string
"""

// Raw string (Swift 5)
let regex = #"^\d{3}-\d{4}$"#
let quote = #"She said "Hello""#

// Character
let char: Character = "A"
let emoji: Character = "😀"

// String is a collection of Characters
for char in greeting {
    print(char)
}
```

### Tuples

```swift
// Basic tuple
let coordinate = (10, 20)
print(coordinate.0, coordinate.1)

// Named elements
let person = (name: "Alice", age: 30)
print(person.name, person.age)

// Decomposition
let (x, y) = coordinate
let (personName, personAge) = person

// Ignore values
let (name, _) = person
```

### Type Aliases

```swift
typealias Coordinate = (x: Int, y: Int)
typealias StringDictionary = [String: String]
typealias CompletionHandler = (Bool) -> Void

let point: Coordinate = (x: 10, y: 20)
let settings: StringDictionary = ["theme": "dark"]
```

### Type Conversion

```swift
// Explicit conversion
let integer = 42
let double = Double(integer)
let float = Float(integer)

let pi = 3.14159
let truncated = Int(pi)  // 3

// String conversion
let number = 123
let text = String(number)

let string = "456"
if let value = Int(string) {
    print("Converted: \(value)")
}

// Type checking
let value: Any = "Hello"
if value is String {
    print("It's a string")
}

// Type casting
if let stringValue = value as? String {
    print("String: \(stringValue)")
}
```

---

## Operators

### Arithmetic Operators

```swift
let a = 10
let b = 3

// Basic arithmetic
let sum = a + b           // 13
let difference = a - b    // 7
let product = a * b       // 30
let quotient = a / b      // 3
let remainder = a % b     // 1

// Unary operators
let positive = +a         // 10
let negative = -a         // -10

// Compound assignment
var x = 10
x += 5   // x = 15
x -= 3   // x = 12
x *= 2   // x = 24
x /= 4   // x = 6
x %= 4   // x = 2
```

### Comparison Operators

```swift
let x = 10
let y = 5

// Comparisons
x == y   // false (equal to)
x != y   // true (not equal to)
x > y    // true (greater than)
x < y    // false (less than)
x >= y   // true (greater than or equal)
x <= y   // false (less than or equal)

// String comparison
"apple" < "banana"  // true (lexicographic)

// Tuple comparison (left to right)
(1, "zebra") < (2, "apple")  // true
(3, "apple") < (3, "bird")   // true
```

### Logical Operators

```swift
let sunny = true
let warm = false

// Logical NOT
!sunny        // false

// Logical AND
sunny && warm  // false

// Logical OR
sunny || warm  // true

// Short-circuit evaluation
let result = false && expensiveOperation()  // expensiveOperation() not called
```

### Range Operators

```swift
// Closed range (includes both ends)
for i in 1...5 {
    print(i)  // 1, 2, 3, 4, 5
}

// Half-open range (excludes upper bound)
for i in 1..<5 {
    print(i)  // 1, 2, 3, 4
}

// One-sided ranges
let numbers = [1, 2, 3, 4, 5]
let firstThree = numbers[...2]     // [1, 2, 3]
let lastTwo = numbers[3...]        // [4, 5]
let allButFirst = numbers[1...]    // [2, 3, 4, 5]

// Range with stride
for i in stride(from: 0, to: 10, by: 2) {
    print(i)  // 0, 2, 4, 6, 8
}
```

### Nil-Coalescing Operator

```swift
let optionalValue: Int? = nil
let defaultValue = 10

// Nil-coalescing
let result = optionalValue ?? defaultValue  // 10

// Chaining
let a: String? = nil
let b: String? = nil
let c: String? = "Hello"
let value = a ?? b ?? c ?? "Default"  // "Hello"
```

### Ternary Conditional Operator

```swift
let age = 20
let status = age >= 18 ? "adult" : "minor"

// Nested ternary (use sparingly)
let score = 85
let grade = score >= 90 ? "A" : score >= 80 ? "B" : "C"
```

---

## Control Flow

### if Statements

```swift
let temperature = 25

// Basic if
if temperature > 30 {
    print("Hot")
}

// if-else
if temperature > 30 {
    print("Hot")
} else {
    print("Not hot")
}

// if-else if-else
if temperature > 30 {
    print("Hot")
} else if temperature > 20 {
    print("Warm")
} else if temperature > 10 {
    print("Cool")
} else {
    print("Cold")
}

// Multiple conditions
let sunny = true
let weekend = false

if sunny && weekend {
    print("Perfect for a picnic!")
} else if sunny || weekend {
    print("Still nice!")
}
```

### guard Statements

```swift
func greet(person: String?) {
    guard let name = person else {
        print("No name provided")
        return
    }
    print("Hello, \(name)!")
}

func divide(_ a: Int, by b: Int) -> Int? {
    guard b != 0 else {
        print("Cannot divide by zero")
        return nil
    }
    return a / b
}

// Multiple conditions
func processUser(name: String?, age: Int?) {
    guard let userName = name,
          let userAge = age,
          userAge >= 18 else {
        print("Invalid user data")
        return
    }
    print("\(userName) is \(userAge) years old")
}
```

### switch Statements

```swift
let number = 3

// Basic switch
switch number {
case 1:
    print("One")
case 2:
    print("Two")
case 3:
    print("Three")
default:
    print("Other")
}

// Multiple values
switch number {
case 1, 2, 3:
    print("One, two, or three")
case 4, 5:
    print("Four or five")
default:
    print("Other")
}

// Range matching
let score = 85
switch score {
case 0..<60:
    print("F")
case 60..<70:
    print("D")
case 70..<80:
    print("C")
case 80..<90:
    print("B")
case 90...100:
    print("A")
default:
    print("Invalid score")
}

// Tuple matching
let point = (1, 1)
switch point {
case (0, 0):
    print("Origin")
case (_, 0):
    print("On X-axis")
case (0, _):
    print("On Y-axis")
case (-2...2, -2...2):
    print("Near origin")
default:
    print("Far from origin")
}

// Value binding
let anotherPoint = (2, 0)
switch anotherPoint {
case (let x, 0):
    print("On X-axis at x=\(x)")
case (0, let y):
    print("On Y-axis at y=\(y)")
case let (x, y):
    print("At (\(x), \(y))")
}

// Where clauses
let yetAnotherPoint = (1, -1)
switch yetAnotherPoint {
case let (x, y) where x == y:
    print("On diagonal")
case let (x, y) where x == -y:
    print("On anti-diagonal")
case let (x, y):
    print("At (\(x), \(y))")
}

// Compound cases
let character: Character = "e"
switch character {
case "a", "e", "i", "o", "u":
    print("Vowel")
default:
    print("Consonant")
}
```

### for-in Loops

```swift
// Range loops
for i in 1...5 {
    print(i)
}

for i in 1..<5 {
    print(i)
}

// Ignore loop variable
for _ in 1...3 {
    print("Hello")
}

// Array iteration
let names = ["Alice", "Bob", "Charlie"]
for name in names {
    print(name)
}

// Dictionary iteration
let ages = ["Alice": 25, "Bob": 30]
for (name, age) in ages {
    print("\(name) is \(age)")
}

// Enumerated iteration
for (index, name) in names.enumerated() {
    print("\(index): \(name)")
}

// Stride
for i in stride(from: 0, to: 10, by: 2) {
    print(i)
}

for i in stride(from: 10, through: 0, by: -2) {
    print(i)
}

// Where clause
for number in 1...10 where number % 2 == 0 {
    print(number)  // Only even numbers
}
```

### while Loops

```swift
var count = 0

// Basic while loop
while count < 5 {
    print(count)
    count += 1
}

// While with complex condition
var playing = true
var score = 0

while playing && score < 100 {
    score += 10
    if score >= 100 {
        playing = false
    }
}
```

### repeat-while Loops

```swift
var number = 0

// Repeat-while (like do-while)
repeat {
    print(number)
    number += 1
} while number < 5

// Executes at least once
var value = 10
repeat {
    print("This runs once even though condition is false")
} while value < 5
```

### Loop Control

```swift
// break: exit loop
for i in 1...10 {
    if i == 5 {
        break
    }
    print(i)  // 1, 2, 3, 4
}

// continue: skip to next iteration
for i in 1...10 {
    if i % 2 == 0 {
        continue
    }
    print(i)  // 1, 3, 5, 7, 9
}

// Labeled statements
outerLoop: for i in 1...3 {
    for j in 1...3 {
        if i == 2 && j == 2 {
            break outerLoop
        }
        print("\(i), \(j)")
    }
}

// fallthrough in switch
let number = 5
switch number {
case 5:
    print("Five")
    fallthrough
case 4:
    print("Four or five")
default:
    print("Other")
}
// Output: "Five" then "Four or five"
```

---

## Functions and Closures

### Function Basics

```swift
// Simple function
func greet() {
    print("Hello!")
}

// Function with parameters
func greet(name: String) {
    print("Hello, \(name)!")
}

// Function with return value
func add(a: Int, b: Int) -> Int {
    return a + b
}

// Implicit return (single expression)
func multiply(_ a: Int, _ b: Int) -> Int {
    a * b
}

// Calling functions
greet()
greet(name: "Alice")
let sum = add(a: 5, b: 3)
```

### Parameter Labels

```swift
// External and internal parameter names
func greet(person name: String, from hometown: String) {
    print("Hello \(name) from \(hometown)!")
}

greet(person: "Alice", from: "Boston")

// Omitting parameter labels
func add(_ a: Int, _ b: Int) -> Int {
    return a + b
}

let result = add(5, 3)  // No labels needed

// Default parameter values
func greet(name: String, greeting: String = "Hello") {
    print("\(greeting), \(name)!")
}

greet(name: "Bob")                    // Uses default "Hello"
greet(name: "Charlie", greeting: "Hi") // Custom greeting
```

### Variadic Parameters

```swift
func sum(_ numbers: Int...) -> Int {
    return numbers.reduce(0, +)
}

let total = sum(1, 2, 3, 4, 5)  // 15

func printNames(_ names: String...) {
    for name in names {
        print(name)
    }
}

printNames("Alice", "Bob", "Charlie")
```

### In-Out Parameters

```swift
func swapValues(_ a: inout Int, _ b: inout Int) {
    let temp = a
    a = b
    b = temp
}

var x = 10
var y = 20
swapValues(&x, &y)
print("x: \(x), y: \(y)")  // x: 20, y: 10

// In-out with modifications
func increment(_ value: inout Int, by amount: Int = 1) {
    value += amount
}

var count = 5
increment(&count)
increment(&count, by: 3)
print(count)  // 9
```

### Function Types

```swift
// Functions are first-class types
func add(_ a: Int, _ b: Int) -> Int {
    a + b
}

func multiply(_ a: Int, _ b: Int) -> Int {
    a * b
}

// Function type variable
var operation: (Int, Int) -> Int = add
print(operation(5, 3))  // 8

operation = multiply
print(operation(5, 3))  // 15

// Function as parameter
func performOperation(_ op: (Int, Int) -> Int, on a: Int, and b: Int) -> Int {
    return op(a, b)
}

let result = performOperation(add, on: 10, and: 5)  // 15

// Function as return type
func chooseOperation(isAdd: Bool) -> (Int, Int) -> Int {
    return isAdd ? add : multiply
}

let selectedOp = chooseOperation(isAdd: true)
print(selectedOp(7, 3))  // 10
```

### Nested Functions

```swift
func calculator(operation: String) -> (Int, Int) -> Int {
    func add(a: Int, b: Int) -> Int {
        return a + b
    }

    func subtract(a: Int, b: Int) -> Int {
        return a - b
    }

    switch operation {
    case "add":
        return add
    case "subtract":
        return subtract
    default:
        return add
    }
}

let addFunc = calculator(operation: "add")
print(addFunc(10, 5))  // 15
```

### Closures

```swift
// Closure expression syntax
let greet = { (name: String) -> String in
    return "Hello, \(name)!"
}

print(greet("Alice"))

// Type inference
let add: (Int, Int) -> Int = { a, b in
    return a + b
}

// Implicit returns
let multiply = { (a: Int, b: Int) -> Int in
    a * b
}

// Shorthand argument names
let divide = { $0 / $1 }
print(divide(10, 2))  // 5

// Trailing closure syntax
func performOperation(_ a: Int, _ b: Int, operation: (Int, Int) -> Int) -> Int {
    return operation(a, b)
}

let result = performOperation(10, 5) { a, b in
    a + b
}

// Multiple trailing closures
func loadData(onSuccess: () -> Void, onFailure: () -> Void) {
    // Simulate loading
    let success = true
    if success {
        onSuccess()
    } else {
        onFailure()
    }
}

loadData {
    print("Success!")
} onFailure: {
    print("Failed!")
}
```

### Capturing Values

```swift
func makeIncrementer(incrementAmount: Int) -> () -> Int {
    var total = 0
    let incrementer: () -> Int = {
        total += incrementAmount
        return total
    }
    return incrementer
}

let incrementByTwo = makeIncrementer(incrementAmount: 2)
print(incrementByTwo())  // 2
print(incrementByTwo())  // 4
print(incrementByTwo())  // 6

// Different instances don't share state
let incrementByTen = makeIncrementer(incrementAmount: 10)
print(incrementByTen())  // 10
print(incrementByTwo())  // 8 (still independent)
```

### Escaping Closures

```swift
var completionHandlers: [() -> Void] = []

func addHandler(handler: @escaping () -> Void) {
    completionHandlers.append(handler)
}

addHandler {
    print("Handler called")
}

// Execute later
for handler in completionHandlers {
    handler()
}

// Async example
func fetchData(completion: @escaping (String) -> Void) {
    DispatchQueue.global().async {
        // Simulate network delay
        sleep(1)
        completion("Data loaded")
    }
}

fetchData { result in
    print(result)
}
```

### Autoclosures

```swift
// Autoclosure delays execution
func logIfTrue(_ condition: @autoclosure () -> Bool) {
    if condition() {
        print("Condition was true")
    }
}

logIfTrue(2 + 2 == 4)  // No need for { }

// Useful for assertions
func assert(_ condition: @autoclosure () -> Bool, _ message: String = "") {
    #if DEBUG
    if !condition() {
        print("Assertion failed: \(message)")
    }
    #endif
}

assert(1 + 1 == 2, "Math is broken")
```

---

## Collections

### Arrays

```swift
// Creating arrays
var numbers: [Int] = [1, 2, 3, 4, 5]
var names = ["Alice", "Bob", "Charlie"]
var emptyArray = [String]()
var anotherEmpty: [Int] = []

// Array with default values
var zeros = Array(repeating: 0, count: 5)  // [0, 0, 0, 0, 0]

// Adding elements
numbers.append(6)
numbers.append(contentsOf: [7, 8, 9])
numbers.insert(0, at: 0)
numbers += [10, 11]

// Removing elements
let last = numbers.removeLast()
let first = numbers.removeFirst()
numbers.remove(at: 2)
numbers.removeAll()

// Accessing elements
let firstElement = names[0]
let lastElement = names[names.count - 1]
let subarray = names[1...2]

// Modifying elements
names[0] = "Alicia"
names[1...2] = ["Robert", "Charles"]

// Array properties
let count = names.count
let isEmpty = names.isEmpty
let capacity = names.capacity

// Iterating
for name in names {
    print(name)
}

for (index, name) in names.enumerated() {
    print("\(index): \(name)")
}

// Array methods
let sorted = numbers.sorted()
let reversed = numbers.reversed()
let filtered = numbers.filter { $0 > 5 }
let mapped = numbers.map { $0 * 2 }
let sum = numbers.reduce(0, +)

// Checking contents
let contains = numbers.contains(5)
let firstGreaterThanFive = numbers.first { $0 > 5 }
let allGreaterThanZero = numbers.allSatisfy { $0 > 0 }

// Combining
let combined = numbers + [100, 200]
```

### Sets

```swift
// Creating sets
var fruits: Set<String> = ["apple", "banana", "orange"]
var numbers: Set = [1, 2, 3, 4, 5]
var emptySet = Set<Int>()

// Adding and removing
fruits.insert("grape")
fruits.remove("banana")
fruits.removeAll()

// Set operations
let set1: Set = [1, 2, 3, 4, 5]
let set2: Set = [4, 5, 6, 7, 8]

let union = set1.union(set2)              // [1, 2, 3, 4, 5, 6, 7, 8]
let intersection = set1.intersection(set2) // [4, 5]
let difference = set1.subtracting(set2)    // [1, 2, 3]
let symmetricDiff = set1.symmetricDifference(set2) // [1, 2, 3, 6, 7, 8]

// Set relationships
let isSubset = set1.isSubset(of: set2)
let isSuperset = set1.isSuperset(of: set2)
let isDisjoint = set1.isDisjoint(with: set2)

// Checking membership
let contains = fruits.contains("apple")

// Iterating (unordered)
for fruit in fruits {
    print(fruit)
}

// Sorted iteration
for fruit in fruits.sorted() {
    print(fruit)
}
```

### Dictionaries

```swift
// Creating dictionaries
var ages: [String: Int] = ["Alice": 25, "Bob": 30]
var scores = ["Math": 95, "English": 88]
var emptyDict = [String: String]()
var anotherEmpty: [Int: String] = [:]

// Adding and updating
ages["Charlie"] = 35
ages.updateValue(26, forKey: "Alice")

// Removing
ages.removeValue(forKey: "Bob")
ages["Charlie"] = nil  // Also removes

// Accessing values
let aliceAge = ages["Alice"]  // Optional Int
let bobAge = ages["Bob", default: 0]  // With default

// Checking for keys
if let age = ages["Alice"] {
    print("Alice is \(age)")
}

// Dictionary properties
let count = ages.count
let isEmpty = ages.isEmpty
let keys = Array(ages.keys)
let values = Array(ages.values)

// Iterating
for (name, age) in ages {
    print("\(name) is \(age)")
}

for name in ages.keys.sorted() {
    print(name)
}

// Dictionary methods
let filtered = ages.filter { $0.value > 25 }
let mapped = ages.mapValues { $0 + 1 }

// Merging dictionaries
var dict1 = ["a": 1, "b": 2]
let dict2 = ["b": 3, "c": 4]

dict1.merge(dict2) { current, new in
    return current  // Keep existing
}
// dict1 is ["a": 1, "b": 2, "c": 4]

// Grouping
let students = ["Alice", "Bob", "Anna", "Charlie", "Amy"]
let grouped = Dictionary(grouping: students) { $0.first! }
// ["A": ["Alice", "Anna", "Amy"], "B": ["Bob"], "C": ["Charlie"]]
```

---

## Strings

### String Basics

```swift
// Creating strings
let simple = "Hello"
let multiline = """
This is a
multi-line
string
"""

// Raw strings (no escaping needed)
let regex = #"^\d{3}-\d{4}$"#
let path = #"C:\Users\Name\Documents"#

// Extended delimiters for interpolation in raw strings
let name = "Alice"
let message = #"Hello, \#(name)!"#

// Empty string
let empty = ""
let alsoEmpty = String()

// String from other types
let number = String(42)
let pi = String(3.14159)
let boolean = String(true)
```

### String Properties and Methods

```swift
let text = "Hello, Swift!"

// Properties
let count = text.count
let isEmpty = text.isEmpty

// Case conversion
let uppercase = text.uppercased()
let lowercase = text.lowercased()

// Checking prefixes and suffixes
let hasPrefix = text.hasPrefix("Hello")
let hasSuffix = text.hasSuffix("!")

// Trimming
let padded = "  Hello  "
let trimmed = padded.trimmingCharacters(in: .whitespaces)

// Replacing
let replaced = text.replacingOccurrences(of: "Swift", with: "World")

// Splitting
let words = text.split(separator: " ")
let components = text.components(separatedBy: ", ")

// Joining
let items = ["apple", "banana", "orange"]
let joined = items.joined(separator: ", ")

// Contains
let contains = text.contains("Swift")
```

### String Indices

```swift
let greeting = "Hello, World!"

// String indices
let startIndex = greeting.startIndex
let endIndex = greeting.endIndex

// Character at index
let firstChar = greeting[startIndex]
let index = greeting.index(startIndex, offsetBy: 7)
let charAtIndex = greeting[index]

// Substring
let start = greeting.index(greeting.startIndex, offsetBy: 7)
let end = greeting.index(greeting.startIndex, offsetBy: 12)
let substring = greeting[start..<end]  // "World"
let substringStr = String(substring)

// Range
if let range = greeting.range(of: "World") {
    let word = greeting[range]
    print(word)
}

// Inserting and removing
var mutable = "Hello"
mutable.insert("!", at: mutable.endIndex)
mutable.insert(contentsOf: " there", at: mutable.index(before: mutable.endIndex))

mutable.remove(at: mutable.index(before: mutable.endIndex))
```

### String Interpolation

```swift
let name = "Alice"
let age = 30
let height = 5.6

// Basic interpolation
let message = "Name: \(name), Age: \(age)"

// Expressions
let calculation = "5 + 3 = \(5 + 3)"

// Formatting
let formatted = String(format: "Pi is approximately %.2f", Double.pi)

// Custom interpolation
extension String.StringInterpolation {
    mutating func appendInterpolation(_ value: Int, format: String) {
        let formatted = String(format: format, value)
        appendLiteral(formatted)
    }
}

let number = 42
let custom = "The number is \(number, format: "%04d")"  // "The number is 0042"
```

### Character Iteration

```swift
let text = "Hello 🌍"

// Iterate characters
for char in text {
    print(char)
}

// Character count
let count = text.count  // 7 (emoji is one character)

// Unicode scalars
for scalar in text.unicodeScalars {
    print("\(scalar.value) ")
}

// UTF-8
for byte in text.utf8 {
    print("\(byte) ")
}
```

---

## Optionals

### Optional Basics

```swift
// Declaring optionals
var name: String? = "Alice"
var age: Int? = nil

// Optional binding with if let
if let unwrappedName = name {
    print("Name is \(unwrappedName)")
} else {
    print("Name is nil")
}

// Multiple optional bindings
var firstName: String? = "John"
var lastName: String? = "Doe"

if let first = firstName, let last = lastName {
    print("Full name: \(first) \(last)")
}

// Optional binding with additional conditions
if let age = age, age >= 18 {
    print("Adult with age \(age)")
}
```

### Guard Statements

```swift
func greet(person: String?) {
    guard let name = person else {
        print("No name provided")
        return
    }
    // name is available here
    print("Hello, \(name)!")
}

// Multiple guards
func processUser(name: String?, age: Int?, email: String?) {
    guard let userName = name else {
        print("Name missing")
        return
    }

    guard let userAge = age, userAge >= 18 else {
        print("Invalid age")
        return
    }

    guard let userEmail = email else {
        print("Email missing")
        return
    }

    print("User: \(userName), \(userAge), \(userEmail)")
}

// Guard with where
func divide(_ a: Int, by b: Int) -> Int? {
    guard b != 0 else {
        return nil
    }
    return a / b
}
```

### Nil-Coalescing Operator

```swift
let optionalName: String? = nil
let defaultName = "Guest"

// Nil-coalescing
let displayName = optionalName ?? defaultName  // "Guest"

// Chaining
let a: String? = nil
let b: String? = nil
let c: String? = "Hello"
let result = a ?? b ?? c ?? "Default"  // "Hello"

// With expressions
let optionalNumber: Int? = nil
let doubled = (optionalNumber ?? 0) * 2
```

### Optional Chaining

```swift
struct Address {
    var street: String
    var city: String
}

struct Person {
    var name: String
    var address: Address?
}

let person = Person(name: "Alice", address: nil)

// Optional chaining
let city = person.address?.city
let street = person.address?.street?.uppercased()

// With method calls
class Car {
    var engine: Engine?
}

class Engine {
    func start() -> Bool {
        return true
    }
}

let car = Car()
let started = car.engine?.start()  // Optional<Bool>

// Dictionary chaining
var dict = ["key": ["nested": "value"]]
let value = dict["key"]?["nested"]  // Optional<String>
```

### Forced Unwrapping

```swift
let optionalString: String? = "Hello"

// Forced unwrapping (use with caution!)
let forcedString = optionalString!  // "Hello"

// Crashes if nil
// let nilValue: String? = nil
// let forced = nilValue!  // Runtime error

// Safe usage after checking
if optionalString != nil {
    print(optionalString!)
}
```

### Implicitly Unwrapped Optionals

```swift
// Declared with !
var assumedString: String! = "Implicitly unwrapped"

// Can be used without unwrapping
print(assumedString)
let length = assumedString.count

// Still optional
assumedString = nil
// print(assumedString)  // Would crash

// Useful for properties that are set after init
class ViewController {
    var label: UILabel!  // Set in viewDidLoad
}
```

### Optional Map and FlatMap

```swift
let optionalNumber: Int? = 5

// Map
let doubled = optionalNumber.map { $0 * 2 }  // Optional(10)
let nilMapped: Int? = nil
let result = nilMapped.map { $0 * 2 }  // nil

// FlatMap
func half(_ value: Int) -> Int? {
    return value % 2 == 0 ? value / 2 : nil
}

let number: Int? = 10
let halved = number.flatMap(half)  // Optional(5)

let oddNumber: Int? = 9
let nilHalf = oddNumber.flatMap(half)  // nil

// Chaining
let result = optionalNumber
    .map { $0 * 2 }
    .flatMap(half)
    .map { $0 + 1 }  // Optional(6)
```

### Optional Pattern Matching

```swift
let optionalValue: Int? = 42

// Switch with optionals
switch optionalValue {
case .none:
    print("No value")
case .some(let value):
    print("Value: \(value)")
}

// If case
if case let .some(value) = optionalValue {
    print("Has value: \(value)")
}

// For case in
let optionals: [Int?] = [1, nil, 3, nil, 5]
for case let .some(value) in optionals {
    print(value)  // 1, 3, 5
}
```

---

## Enumerations

### Basic Enumerations

```swift
enum Direction {
    case north
    case south
    case east
    case west
}

// Shorter syntax
enum CompassPoint {
    case north, south, east, west
}

// Using enums
var direction = Direction.north
direction = .south  // Type inferred

// Switch with enums
switch direction {
case .north:
    print("Going north")
case .south:
    print("Going south")
case .east:
    print("Going east")
case .west:
    print("Going west")
}
```

### Raw Values

```swift
// Int raw values (auto-incremented)
enum Planet: Int {
    case mercury = 1
    case venus      // 2
    case earth      // 3
    case mars       // 4
}

let earth = Planet.earth
print(earth.rawValue)  // 3

// Initialize from raw value
if let planet = Planet(rawValue: 2) {
    print(planet)  // venus
}

// String raw values
enum Size: String {
    case small
    case medium = "M"
    case large = "L"
}

print(Size.small.rawValue)   // "small"
print(Size.medium.rawValue)  // "M"
```

### Associated Values

```swift
enum Barcode {
    case upc(Int, Int, Int, Int)
    case qrCode(String)
}

// Creating instances
let productBarcode = Barcode.upc(8, 85909, 51226, 3)
let urlBarcode = Barcode.qrCode("https://example.com")

// Extracting associated values
switch productBarcode {
case .upc(let numberSystem, let manufacturer, let product, let check):
    print("UPC: \(numberSystem), \(manufacturer), \(product), \(check)")
case .qrCode(let code):
    print("QR Code: \(code)")
}

// Shorthand
switch productBarcode {
case let .upc(a, b, c, d):
    print("UPC: \(a), \(b), \(c), \(d)")
case let .qrCode(code):
    print("QR: \(code)")
}

// More complex example
enum ServerResponse {
    case success(String, Int)
    case failure(String)
    case retry(Int)
}

let response = ServerResponse.success("Data loaded", 200)

switch response {
case .success(let message, let code):
    print("\(message) - Code: \(code)")
case .failure(let error):
    print("Error: \(error)")
case .retry(let attempts):
    print("Retry in \(attempts) seconds")
}
```

### Enum Methods and Properties

```swift
enum TrafficLight {
    case red, yellow, green

    func duration() -> Int {
        switch self {
        case .red:
            return 30
        case .yellow:
            return 5
        case .green:
            return 25
        }
    }

    mutating func next() {
        switch self {
        case .red:
            self = .green
        case .yellow:
            self = .red
        case .green:
            self = .yellow
        }
    }

    var description: String {
        switch self {
        case .red:
            return "Stop"
        case .yellow:
            return "Caution"
        case .green:
            return "Go"
        }
    }
}

var light = TrafficLight.red
print(light.description)      // "Stop"
print(light.duration())       // 30
light.next()
print(light.description)      // "Go"
```

### Recursive Enumerations

```swift
indirect enum ArithmeticExpression {
    case number(Int)
    case addition(ArithmeticExpression, ArithmeticExpression)
    case multiplication(ArithmeticExpression, ArithmeticExpression)
}

// Build: (5 + 4) * 2
let five = ArithmeticExpression.number(5)
let four = ArithmeticExpression.number(4)
let sum = ArithmeticExpression.addition(five, four)
let product = ArithmeticExpression.multiplication(sum, ArithmeticExpression.number(2))

// Evaluate
func evaluate(_ expression: ArithmeticExpression) -> Int {
    switch expression {
    case let .number(value):
        return value
    case let .addition(left, right):
        return evaluate(left) + evaluate(right)
    case let .multiplication(left, right):
        return evaluate(left) * evaluate(right)
    }
}

print(evaluate(product))  // 18
```

### CaseIterable Protocol

```swift
enum Beverage: CaseIterable {
    case coffee, tea, juice, water
}

// Iterate all cases
for beverage in Beverage.allCases {
    print(beverage)
}

print("Number of beverages: \(Beverage.allCases.count)")

// With associated values (manual implementation)
enum Direction: CaseIterable {
    case north, south, east, west

    static var allCases: [Direction] {
        return [.north, .south, .east, .west]
    }
}
```

---

## Structures and Classes

### Defining Structures

```swift
struct Person {
    var name: String
    var age: Int
}

// Creating instances
let person = Person(name: "Alice", age: 30)
var mutablePerson = Person(name: "Bob", age: 25)

// Accessing properties
print(person.name)
print(person.age)

// Modifying (only if var)
mutablePerson.age = 26

// Memberwise initializer (automatic for structs)
let another = Person(name: "Charlie", age: 35)
```

### Defining Classes

```swift
class BankAccount {
    var balance: Double
    var accountNumber: String

    // Classes require explicit initializer
    init(accountNumber: String, balance: Double = 0.0) {
        self.accountNumber = accountNumber
        self.balance = balance
    }

    // Deinitializer
    deinit {
        print("Account \(accountNumber) closed")
    }
}

// Creating instances
let account = BankAccount(accountNumber: "12345", balance: 1000)
```

### Value Types vs Reference Types

```swift
// Structs are value types (copied)
struct Point {
    var x: Int
    var y: Int
}

var point1 = Point(x: 10, y: 20)
var point2 = point1  // Copy
point2.x = 30

print(point1.x)  // 10 (unchanged)
print(point2.x)  // 30

// Classes are reference types (shared)
class Rectangle {
    var width: Int
    var height: Int

    init(width: Int, height: Int) {
        self.width = width
        self.height = height
    }
}

let rect1 = Rectangle(width: 100, height: 50)
let rect2 = rect1  // Same reference
rect2.width = 200

print(rect1.width)  // 200 (changed!)
print(rect2.width)  // 200

// Identity operators for classes
let rect3 = Rectangle(width: 100, height: 50)
rect1 === rect2  // true (same instance)
rect1 === rect3  // false (different instances)
```

### Methods

```swift
struct Counter {
    var count = 0

    // Instance method
    mutating func increment() {
        count += 1
    }

    mutating func increment(by amount: Int) {
        count += amount
    }

    mutating func reset() {
        count = 0
    }

    // Methods can return values
    func doubled() -> Int {
        return count * 2
    }
}

var counter = Counter()
counter.increment()
counter.increment(by: 5)
print(counter.count)  // 6

// Classes don't need mutating
class MutableCounter {
    var count = 0

    func increment() {
        count += 1
    }
}
```

### Type Methods and Properties

```swift
struct Math {
    static let pi = 3.14159

    static func square(_ value: Double) -> Double {
        return value * value
    }

    static func cube(_ value: Double) -> Double {
        return value * value * value
    }
}

print(Math.pi)
print(Math.square(5.0))  // 25.0

// Class type methods
class Player {
    static var highScore = 0

    class func resetHighScore() {
        highScore = 0
    }
}

Player.highScore = 1000
Player.resetHighScore()
print(Player.highScore)  // 0
```

### Subscripts

```swift
struct TimesTable {
    let multiplier: Int

    subscript(index: Int) -> Int {
        return multiplier * index
    }
}

let threeTimesTable = TimesTable(multiplier: 3)
print(threeTimesTable[6])  // 18

// Dictionary-like subscript
struct StringDictionary {
    private var data: [String: String] = [:]

    subscript(key: String) -> String? {
        get {
            return data[key]
        }
        set {
            data[key] = newValue
        }
    }
}

var dict = StringDictionary()
dict["name"] = "Alice"
print(dict["name"] ?? "Unknown")
```

### Inheritance

```swift
// Base class
class Vehicle {
    var speed: Double = 0.0

    func describe() -> String {
        return "Traveling at \(speed) mph"
    }

    func makeNoise() {
        // Default implementation
    }
}

// Subclass
class Car: Vehicle {
    var gear = 1

    override func describe() -> String {
        return super.describe() + " in gear \(gear)"
    }

    override func makeNoise() {
        print("Vroom!")
    }
}

// Another subclass
class Bicycle: Vehicle {
    var hasBasket = false

    override func makeNoise() {
        print("Ring ring!")
    }
}

let car = Car()
car.speed = 60.0
car.gear = 4
print(car.describe())
car.makeNoise()

let bike = Bicycle()
bike.hasBasket = true
bike.speed = 15.0
print(bike.describe())
bike.makeNoise()
```

### Preventing Overrides

```swift
class BaseClass {
    final func cannotOverride() {
        print("This cannot be overridden")
    }

    func canOverride() {
        print("This can be overridden")
    }
}

// Final class (cannot be subclassed)
final class FinalClass {
    func method() {
        print("From final class")
    }
}

// class SubClass: FinalClass { }  // Error!
```

---

## Properties and Methods

### Stored Properties

```swift
struct Person {
    var name: String
    var age: Int
    let id: String  // Constant property

    // Lazy stored property
    lazy var fullProfile: String = {
        return "Profile for \(name), age \(age)"
    }()
}

var person = Person(name: "Alice", age: 30, id: "A123")
person.name = "Alicia"
// person.id = "B456"  // Error: let property
print(person.fullProfile)  // Computed when first accessed
```

### Computed Properties

```swift
struct Rectangle {
    var width: Double
    var height: Double

    // Computed property
    var area: Double {
        return width * height
    }

    // Get and set
    var perimeter: Double {
        get {
            return 2 * (width + height)
        }
        set {
            // newValue is automatic
            let side = newValue / 4
            width = side
            height = side
        }
    }

    // Read-only computed property
    var diagonal: Double {
        return (width * width + height * height).squareRoot()
    }
}

var rect = Rectangle(width: 10, height: 5)
print(rect.area)       // 50
print(rect.perimeter)  // 30
rect.perimeter = 40    // Sets width and height
print(rect.width)      // 10
```

### Property Observers

```swift
class StepCounter {
    var totalSteps: Int = 0 {
        willSet {
            print("About to set totalSteps to \(newValue)")
        }
        didSet {
            if totalSteps > oldValue {
                print("Added \(totalSteps - oldValue) steps")
            }
        }
    }
}

let counter = StepCounter()
counter.totalSteps = 100  // Triggers observers
counter.totalSteps = 150  // Triggers again

// With computed properties
class Temperature {
    var celsius: Double = 0 {
        didSet {
            print("Temperature changed from \(oldValue) to \(celsius)")
        }
    }

    var fahrenheit: Double {
        get {
            return celsius * 1.8 + 32
        }
        set {
            celsius = (newValue - 32) / 1.8
        }
    }
}

let temp = Temperature()
temp.fahrenheit = 72  // Also triggers celsius observer
```

### Property Wrappers

```swift
@propertyWrapper
struct Clamped<T: Comparable> {
    private var value: T
    private let range: ClosedRange<T>

    init(wrappedValue: T, _ range: ClosedRange<T>) {
        self.range = range
        self.value = min(max(wrappedValue, range.lowerBound), range.upperBound)
    }

    var wrappedValue: T {
        get { value }
        set { value = min(max(newValue, range.lowerBound), range.upperBound) }
    }

    var projectedValue: T {
        return value
    }
}

struct Game {
    @Clamped(0...100) var health = 100
    @Clamped(1...10) var level = 1
}

var game = Game()
game.health = 150  // Clamped to 100
print(game.health)  // 100
game.health = -10   // Clamped to 0
print(game.health)  // 0

// Built-in property wrappers
class Settings {
    @Published var username: String = ""  // Combine framework
    @State private var isLoggedIn = false  // SwiftUI
}
```

### Type Properties

```swift
struct Math {
    static let pi = 3.14159
    static var computationCount = 0

    static func square(_ x: Double) -> Double {
        computationCount += 1
        return x * x
    }
}

print(Math.pi)
let result = Math.square(5.0)
print(Math.computationCount)  // 1

// Class type properties
class Player {
    static var highScore = 0
    class var maxPlayers: Int {
        return 4
    }
}
```

### Methods

```swift
class Counter {
    var count = 0

    func increment() {
        count += 1
    }

    func increment(by amount: Int) {
        count += amount
    }

    func reset() {
        count = 0
    }

    // Method with return value
    func doubled() -> Int {
        return count * 2
    }

    // Type method
    static func description() -> String {
        return "A simple counter"
    }
}

let counter = Counter()
counter.increment()
counter.increment(by: 5)
print(counter.doubled())
print(Counter.description())
```

---

## Protocols and Extensions

### Defining Protocols

```swift
protocol Drawable {
    func draw()
}

protocol Named {
    var name: String { get set }
    var fullName: String { get }
}

protocol Identifiable {
    var id: String { get }
    static func generateID() -> String
}

// Protocol with mutating method
protocol Toggleable {
    mutating func toggle()
}
```

### Conforming to Protocols

```swift
struct Circle: Drawable {
    var radius: Double

    func draw() {
        print("Drawing a circle with radius \(radius)")
    }
}

struct Rectangle: Drawable {
    var width: Double
    var height: Double

    func draw() {
        print("Drawing a rectangle \(width)x\(height)")
    }
}

class Person: Named {
    var name: String
    var lastName: String

    var fullName: String {
        return "\(name) \(lastName)"
    }

    init(name: String, lastName: String) {
        self.name = name
        self.lastName = lastName
    }
}

// Using protocol as type
let shapes: [Drawable] = [
    Circle(radius: 5.0),
    Rectangle(width: 10, height: 5)
]

for shape in shapes {
    shape.draw()
}
```

### Protocol Inheritance

```swift
protocol Printable {
    func printDescription()
}

protocol Loggable: Printable {
    func log()
}

struct Document: Loggable {
    var title: String

    func printDescription() {
        print("Document: \(title)")
    }

    func log() {
        print("Logging: \(title)")
    }
}

// Multiple protocol conformance
protocol Saveable {
    func save()
}

struct File: Printable, Saveable {
    var name: String

    func printDescription() {
        print("File: \(name)")
    }

    func save() {
        print("Saving \(name)")
    }
}
```

### Protocol Extensions

```swift
protocol Describable {
    var description: String { get }
}

// Default implementation
extension Describable {
    var description: String {
        return "A describable object"
    }

    func printDescription() {
        print(description)
    }
}

struct Book: Describable {
    var title: String

    // Can override default
    var description: String {
        return "Book: \(title)"
    }
}

let book = Book(title: "Swift Programming")
book.printDescription()

// Constrained extensions
extension Collection where Element: Numeric {
    func sum() -> Element {
        return reduce(0, +)
    }
}

let numbers = [1, 2, 3, 4, 5]
print(numbers.sum())  // 15
```

### Protocol Composition

```swift
protocol Named {
    var name: String { get }
}

protocol Aged {
    var age: Int { get }
}

// Use multiple protocols
func celebrate(person: Named & Aged) {
    print("\(person.name) is turning \(person.age)!")
}

struct Person: Named, Aged {
    var name: String
    var age: Int
}

let person = Person(name: "Alice", age: 30)
celebrate(person: person)

// Protocol composition with class
class Location {
    var latitude: Double
    var longitude: Double

    init(latitude: Double, longitude: Double) {
        self.latitude = latitude
        self.longitude = longitude
    }
}

protocol Locatable {
    var location: Location { get }
}

func show(place: Location & Locatable) {
    // Must be Location class and conform to Locatable
}
```

### Extensions

```swift
// Extend existing types
extension Int {
    func squared() -> Int {
        return self * self
    }

    var isEven: Bool {
        return self % 2 == 0
    }

    func times(_ closure: () -> Void) {
        for _ in 0..<self {
            closure()
        }
    }
}

print(5.squared())  // 25
print(4.isEven)     // true
3.times {
    print("Hello")
}

// Extend String
extension String {
    func reverse() -> String {
        return String(self.reversed())
    }

    var wordCount: Int {
        return split(separator: " ").count
    }
}

let text = "Hello, World!"
print(text.reverse())
print(text.wordCount)

// Add conformance
extension Int: Drawable {
    func draw() {
        print("Drawing number \(self)")
    }
}

5.draw()

// Constrained extensions
extension Array where Element == Int {
    func sum() -> Int {
        return reduce(0, +)
    }
}

let numbers = [1, 2, 3, 4, 5]
print(numbers.sum())
```

### Associated Types

```swift
protocol Container {
    associatedtype Item
    mutating func append(_ item: Item)
    var count: Int { get }
    subscript(i: Int) -> Item { get }
}

struct IntStack: Container {
    typealias Item = Int

    private var items: [Int] = []

    mutating func append(_ item: Int) {
        items.append(item)
    }

    var count: Int {
        return items.count
    }

    subscript(i: Int) -> Int {
        return items[i]
    }
}

// Generic implementation
struct Stack<Element>: Container {
    typealias Item = Element

    private var items: [Element] = []

    mutating func append(_ item: Element) {
        items.append(item)
    }

    var count: Int {
        return items.count
    }

    subscript(i: Int) -> Element {
        return items[i]
    }
}
```

---

## Generics

### Generic Functions

```swift
// Simple generic function
func swapValues<T>(_ a: inout T, _ b: inout T) {
    let temp = a
    a = b
    b = temp
}

var x = 10
var y = 20
swapValues(&x, &y)
print("x: \(x), y: \(y)")

var str1 = "Hello"
var str2 = "World"
swapValues(&str1, &str2)
print("str1: \(str1), str2: \(str2)")

// Multiple type parameters
func zip<T, U>(_ first: T, _ second: U) -> (T, U) {
    return (first, second)
}

let pair = zip(42, "Answer")
```

### Generic Types

```swift
struct Stack<Element> {
    private var items: [Element] = []

    mutating func push(_ item: Element) {
        items.append(item)
    }

    mutating func pop() -> Element? {
        return items.isEmpty ? nil : items.removeLast()
    }

    func peek() -> Element? {
        return items.last
    }

    var isEmpty: Bool {
        return items.isEmpty
    }

    var count: Int {
        return items.count
    }
}

var intStack = Stack<Int>()
intStack.push(1)
intStack.push(2)
intStack.push(3)
print(intStack.pop() ?? 0)

var stringStack = Stack<String>()
stringStack.push("Hello")
stringStack.push("World")
```

### Type Constraints

```swift
// Constraint with protocol
func findIndex<T: Equatable>(of valueToFind: T, in array: [T]) -> Int? {
    for (index, value) in array.enumerated() {
        if value == valueToFind {
            return index
        }
    }
    return nil
}

let numbers = [1, 2, 3, 4, 5]
if let index = findIndex(of: 3, in: numbers) {
    print("Found at index \(index)")
}

// Multiple constraints
func compare<T: Comparable & Numeric>(_ a: T, _ b: T) -> T {
    return a > b ? a : b
}

let max = compare(5, 10)  // 10

// where clause
func allEqual<T: Equatable, C: Collection>(_ collection: C) -> Bool
    where C.Element == T {
    guard let first = collection.first else { return true }
    return collection.allSatisfy { $0 == first }
}
```

### Associated Types in Protocols

```swift
protocol Container {
    associatedtype Item
    mutating func append(_ item: Item)
    var count: Int { get }
    subscript(i: Int) -> Item { get }
}

// Generic type conforming to protocol
struct GenericStack<T>: Container {
    typealias Item = T

    private var items: [T] = []

    mutating func append(_ item: T) {
        items.append(item)
    }

    var count: Int {
        return items.count
    }

    subscript(i: Int) -> T {
        return items[i]
    }
}

// Using associated types in functions
func merge<C1: Container, C2: Container>(_ c1: C1, _ c2: C2) -> [C1.Item]
    where C1.Item == C2.Item, C1.Item: Equatable {
    var result: [C1.Item] = []
    for i in 0..<c1.count {
        result.append(c1[i])
    }
    for i in 0..<c2.count {
        result.append(c2[i])
    }
    return result
}
```

### Generic Where Clauses

```swift
// Extension with where clause
extension Stack where Element: Equatable {
    func isTop(_ item: Element) -> Bool {
        guard let topItem = peek() else {
            return false
        }
        return topItem == item
    }
}

// Contextual where clauses
extension Container {
    func average() -> Double where Item == Int {
        var sum = 0
        for i in 0..<count {
            sum += self[i]
        }
        return count > 0 ? Double(sum) / Double(count) : 0
    }
}

extension Container where Item: Equatable {
    func contains(_ item: Item) -> Bool {
        for i in 0..<count {
            if self[i] == item {
                return true
            }
        }
        return false
    }
}
```

---

## Error Handling

### Defining Errors

```swift
// Error enum
enum NetworkError: Error {
    case badURL
    case requestFailed
    case invalidResponse
    case decodingError
}

// Error with associated values
enum ValidationError: Error {
    case emptyField(String)
    case invalidFormat(String, expected: String)
    case outOfRange(Int, min: Int, max: Int)
}

// Error with custom properties
struct CustomError: Error {
    let message: String
    let code: Int

    var localizedDescription: String {
        return "Error \(code): \(message)"
    }
}
```

### Throwing Functions

```swift
func validateAge(_ age: Int) throws -> Bool {
    if age < 0 {
        throw ValidationError.outOfRange(age, min: 0, max: 150)
    }
    if age > 150 {
        throw ValidationError.outOfRange(age, min: 0, max: 150)
    }
    return true
}

func fetchData(from url: String) throws -> String {
    guard !url.isEmpty else {
        throw NetworkError.badURL
    }
    // Simulate network request
    return "Sample data"
}

// Throwing initializer
struct Document {
    let content: String

    init(path: String) throws {
        guard !path.isEmpty else {
            throw ValidationError.emptyField("path")
        }
        self.content = "File content"
    }
}
```

### Handling Errors with do-catch

```swift
// Basic do-catch
do {
    try validateAge(200)
} catch ValidationError.outOfRange(let value, let min, let max) {
    print("Age \(value) is out of range [\(min)-\(max)]")
} catch {
    print("Other error: \(error)")
}

// Multiple catch blocks
do {
    let data = try fetchData(from: "")
    print(data)
} catch NetworkError.badURL {
    print("Invalid URL")
} catch NetworkError.requestFailed {
    print("Request failed")
} catch {
    print("Unknown error: \(error)")
}

// Pattern matching in catch
do {
    try validateAge(-5)
} catch ValidationError.outOfRange(let value, _, _) where value < 0 {
    print("Negative age: \(value)")
} catch ValidationError.outOfRange {
    print("Age out of range")
} catch {
    print("Other error")
}
```

### Optional Try

```swift
// try? converts error to nil
let data = try? fetchData(from: "https://example.com")
if let result = data {
    print("Got data: \(result)")
} else {
    print("Failed to fetch data")
}

// try! force unwrap (crashes on error)
// Only use when you're absolutely sure it won't fail
// let forcedData = try! fetchData(from: "")

// Chaining optional try
let processed = try? fetchData(from: "https://example.com")
    .uppercased()
    .replacingOccurrences(of: " ", with: "_")
```

### Propagating Errors

```swift
func processUserData(id: String) throws -> String {
    let rawData = try fetchData(from: "https://api.example.com/user/\(id)")
    // Process data...
    return rawData
}

func loadAndProcessUser(id: String) throws -> String {
    // Error is automatically propagated
    let userData = try processUserData(id: id)
    return userData.uppercased()
}

// Using throw in the function
do {
    let user = try loadAndProcessUser(id: "123")
    print("User: \(user)")
} catch {
    print("Failed to load user: \(error)")
}
```

### Cleanup with defer

```swift
func processFile(filename: String) throws {
    let file = try openFile(filename)
    defer {
        closeFile(file)  // Always executed when scope exits
        print("File closed")
    }

    // Process file...
    if someCondition {
        throw CustomError(message: "Processing failed", code: 1)
    }

    // defer is still called even if error is thrown
}

// Multiple defer statements (executed in reverse order)
func example() {
    defer { print("First defer") }
    defer { print("Second defer") }
    defer { print("Third defer") }
    print("Function body")
}
// Output:
// Function body
// Third defer
// Second defer
// First defer

// File operations helper
func openFile(_ name: String) throws -> FileHandle {
    return FileHandle()  // Placeholder
}

func closeFile(_ file: FileHandle) {
    print("Closing file")
}

let someCondition = false
```

### Result Type

```swift
// Function returning Result
func fetchUserData(id: Int) -> Result<String, NetworkError> {
    if id <= 0 {
        return .failure(.badURL)
    }
    return .success("User data for ID \(id)")
}

// Using Result
let result = fetchUserData(id: 123)

switch result {
case .success(let data):
    print("Success: \(data)")
case .failure(let error):
    print("Error: \(error)")
}

// Result with map and flatMap
let processed = fetchUserData(id: 456)
    .map { data in
        data.uppercased()
    }

// Get value or default
let userData = fetchUserData(id: 789).get(default: "No data")

// Converting throws to Result
func loadData() throws -> String {
    return "Data"
}

let resultFromThrows = Result { try loadData() }
```

---

## File I/O and JSON

### Reading Files

```swift
import Foundation

// Read entire file
do {
    let contents = try String(contentsOfFile: "file.txt", encoding: .utf8)
    print(contents)
} catch {
    print("Error reading file: \(error)")
}

// Read with FileManager
let fileManager = FileManager.default
let currentPath = fileManager.currentDirectoryPath

if fileManager.fileExists(atPath: "file.txt") {
    do {
        let data = try Data(contentsOf: URL(fileURLWithPath: "file.txt"))
        if let text = String(data: data, encoding: .utf8) {
            print(text)
        }
    } catch {
        print("Error: \(error)")
    }
}

// Read line by line
if let fileURL = Bundle.main.url(forResource: "file", withExtension: "txt") {
    do {
        let contents = try String(contentsOf: fileURL, encoding: .utf8)
        let lines = contents.components(separatedBy: .newlines)
        for line in lines {
            print(line)
        }
    } catch {
        print("Error: \(error)")
    }
}
```

### Writing Files

```swift
import Foundation

// Write string to file
let content = "Hello, Swift!\nThis is a test file."

do {
    try content.write(toFile: "output.txt", atomically: true, encoding: .utf8)
    print("File written successfully")
} catch {
    print("Error writing file: \(error)")
}

// Write Data
let data = content.data(using: .utf8)!
do {
    try data.write(to: URL(fileURLWithPath: "output.txt"))
    print("Data written successfully")
} catch {
    print("Error: \(error)")
}

// Append to file
if let fileHandle = FileHandle(forWritingAtPath: "output.txt") {
    defer {
        fileHandle.closeFile()
    }

    fileHandle.seekToEndOfFile()
    if let data = "\nAppended line".data(using: .utf8) {
        fileHandle.write(data)
    }
}
```

### File Operations

```swift
import Foundation

let fileManager = FileManager.default

// Get current directory
let currentPath = fileManager.currentDirectoryPath
print("Current path: \(currentPath)")

// Check if file exists
let fileExists = fileManager.fileExists(atPath: "file.txt")
print("File exists: \(fileExists)")

// Create directory
do {
    try fileManager.createDirectory(atPath: "testDir",
                                   withIntermediateDirectories: true)
    print("Directory created")
} catch {
    print("Error creating directory: \(error)")
}

// List directory contents
do {
    let contents = try fileManager.contentsOfDirectory(atPath: ".")
    print("Directory contents:")
    for item in contents {
        print("  - \(item)")
    }
} catch {
    print("Error listing directory: \(error)")
}

// Copy file
do {
    try fileManager.copyItem(atPath: "source.txt", toPath: "destination.txt")
    print("File copied")
} catch {
    print("Error copying file: \(error)")
}

// Move file
do {
    try fileManager.moveItem(atPath: "old.txt", toPath: "new.txt")
    print("File moved")
} catch {
    print("Error moving file: \(error)")
}

// Delete file
do {
    try fileManager.removeItem(atPath: "file.txt")
    print("File deleted")
} catch {
    print("Error deleting file: \(error)")
}

// Get file attributes
do {
    let attributes = try fileManager.attributesOfItem(atPath: "file.txt")
    if let size = attributes[.size] as? Int {
        print("File size: \(size) bytes")
    }
    if let modDate = attributes[.modificationDate] as? Date {
        print("Modified: \(modDate)")
    }
} catch {
    print("Error getting attributes: \(error)")
}
```

### JSON Encoding and Decoding

```swift
import Foundation

// Define Codable struct
struct Person: Codable {
    var name: String
    var age: Int
    var email: String?
    var hobbies: [String]
}

// Encoding to JSON
let person = Person(
    name: "Alice",
    age: 30,
    email: "alice@example.com",
    hobbies: ["reading", "coding", "hiking"]
)

let encoder = JSONEncoder()
encoder.outputFormatting = .prettyPrinted

do {
    let jsonData = try encoder.encode(person)
    if let jsonString = String(data: jsonData, encoding: .utf8) {
        print("JSON:\n\(jsonString)")

        // Write to file
        try jsonString.write(toFile: "person.json", atomically: true, encoding: .utf8)
    }
} catch {
    print("Encoding error: \(error)")
}

// Decoding from JSON
do {
    let jsonData = try Data(contentsOf: URL(fileURLWithPath: "person.json"))
    let decoder = JSONDecoder()
    let decodedPerson = try decoder.decode(Person.self, from: jsonData)
    print("\nDecoded: \(decodedPerson.name), \(decodedPerson.age)")
} catch {
    print("Decoding error: \(error)")
}

// Custom keys
struct User: Codable {
    var username: String
    var fullName: String
    var age: Int

    enum CodingKeys: String, CodingKey {
        case username
        case fullName = "full_name"  // Different JSON key
        case age
    }
}

// Manual encoding/decoding
extension User {
    init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: CodingKeys.self)
        username = try container.decode(String.self, forKey: .username)
        fullName = try container.decode(String.self, forKey: .fullName)
        age = try container.decode(Int.self, forKey: .age)
    }

    func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encode(username, forKey: .username)
        try container.encode(fullName, forKey: .fullName)
        try container.encode(age, forKey: .age)
    }
}

// Array of objects
let people = [person, person, person]
do {
    let jsonData = try encoder.encode(people)
    if let jsonString = String(data: jsonData, encoding: .utf8) {
        print("\nArray JSON:\n\(jsonString)")
    }
} catch {
    print("Error: \(error)")
}
```

---

## Memory Management

### ARC (Automatic Reference Counting)

```swift
class Person {
    let name: String

    init(name: String) {
        self.name = name
        print("\(name) is initialized")
    }

    deinit {
        print("\(name) is being deinitialized")
    }
}

// Strong references
var person1: Person? = Person(name: "Alice")
var person2 = person1
var person3 = person1

person1 = nil
person2 = nil
// Person still exists (person3 holds strong reference)
person3 = nil
// Now Person is deinitialized
```

### Strong Reference Cycles

```swift
class Apartment {
    let unit: String
    var tenant: Person?

    init(unit: String) {
        self.unit = unit
    }

    deinit {
        print("Apartment \(unit) is being deinitialized")
    }
}

class Person {
    let name: String
    var apartment: Apartment?

    init(name: String) {
        self.name = name
    }

    deinit {
        print("\(name) is being deinitialized")
    }
}

// Strong reference cycle (memory leak)
var john: Person? = Person(name: "John")
var unit4A: Apartment? = Apartment(unit: "4A")

john?.apartment = unit4A
unit4A?.tenant = john

john = nil
unit4A = nil
// Neither is deinitialized (strong reference cycle)
```

### Weak References

```swift
class Person {
    let name: String
    var apartment: Apartment?

    init(name: String) {
        self.name = name
    }

    deinit {
        print("\(name) is being deinitialized")
    }
}

class Apartment {
    let unit: String
    weak var tenant: Person?  // Weak reference

    init(unit: String) {
        self.unit = unit
    }

    deinit {
        print("Apartment \(unit) is being deinitialized")
    }
}

var john: Person? = Person(name: "John")
var unit4A: Apartment? = Apartment(unit: "4A")

john?.apartment = unit4A
unit4A?.tenant = john

john = nil  // Person is deinitialized
unit4A = nil  // Apartment is deinitialized
```

### Unowned References

```swift
class Customer {
    let name: String
    var card: CreditCard?

    init(name: String) {
        self.name = name
    }

    deinit {
        print("\(name) is being deinitialized")
    }
}

class CreditCard {
    let number: UInt64
    unowned let customer: Customer  // Unowned reference

    init(number: UInt64, customer: Customer) {
        self.number = number
        self.customer = customer
    }

    deinit {
        print("Card #\(number) is being deinitialized")
    }
}

var john: Customer? = Customer(name: "John Appleseed")
john?.card = CreditCard(number: 1234_5678_9012_3456, customer: john!)

john = nil
// Both are deinitialized
```

### Closure Capture Lists

```swift
class HTMLElement {
    let name: String
    let text: String?

    // Strong reference cycle with closure
    lazy var asHTML: () -> String = {
        if let text = self.text {
            return "<\(self.name)>\(text)</\(self.name)>"
        } else {
            return "<\(self.name) />"
        }
    }

    init(name: String, text: String? = nil) {
        self.name = name
        self.text = text
    }

    deinit {
        print("\(name) is being deinitialized")
    }
}

var heading: HTMLElement? = HTMLElement(name: "h1", text: "Hello")
print(heading?.asHTML() ?? "")
heading = nil  // Not deinitialized (strong reference cycle)

// Fix with capture list
class HTMLElementFixed {
    let name: String
    let text: String?

    lazy var asHTML: () -> String = { [weak self] in
        guard let self = self else {
            return ""
        }
        if let text = self.text {
            return "<\(self.name)>\(text)</\(self.name)>"
        } else {
            return "<\(self.name) />"
        }
    }

    init(name: String, text: String? = nil) {
        self.name = name
        self.text = text
    }

    deinit {
        print("\(name) is being deinitialized")
    }
}

var paragraph: HTMLElementFixed? = HTMLElementFixed(name: "p", text: "Hello")
print(paragraph?.asHTML() ?? "")
paragraph = nil  // Properly deinitialized

// Unowned in closures
class NumberTracker {
    var number: Int

    lazy var incrementer: () -> Int = { [unowned self] in
        self.number += 1
        return self.number
    }

    init(number: Int) {
        self.number = number
    }

    deinit {
        print("Tracker deinitialized")
    }
}
```

---

## Advanced Topics

### Nested Types

```swift
struct BlackjackCard {
    enum Suit: Character {
        case spades = "♠", hearts = "♡", diamonds = "♢", clubs = "♣"
    }

    enum Rank: Int {
        case two = 2, three, four, five, six, seven, eight, nine, ten
        case jack, queen, king, ace

        struct Values {
            let first: Int
            let second: Int?
        }

        var values: Values {
            switch self {
            case .ace:
                return Values(first: 1, second: 11)
            case .jack, .queen, .king:
                return Values(first: 10, second: nil)
            default:
                return Values(first: self.rawValue, second: nil)
            }
        }
    }

    let rank: Rank
    let suit: Suit

    var description: String {
        var output = "\(rank.values.first)"
        if let second = rank.values.second {
            output += " or \(second)"
        }
        output += " of \(suit.rawValue)"
        return output
    }
}

let card = BlackjackCard(rank: .ace, suit: .spades)
print(card.description)
```

### Type Casting

```swift
class MediaItem {
    var name: String
    init(name: String) {
        self.name = name
    }
}

class Movie: MediaItem {
    var director: String
    init(name: String, director: String) {
        self.director = director
        super.init(name: name)
    }
}

class Song: MediaItem {
    var artist: String
    init(name: String, artist: String) {
        self.artist = artist
        super.init(name: name)
    }
}

let library: [MediaItem] = [
    Movie(name: "Inception", director: "Christopher Nolan"),
    Song(name: "Bohemian Rhapsody", artist: "Queen"),
    Movie(name: "The Matrix", director: "Wachowskis"),
    Song(name: "Imagine", artist: "John Lennon")
]

// Type checking
var movieCount = 0
var songCount = 0

for item in library {
    if item is Movie {
        movieCount += 1
    } else if item is Song {
        songCount += 1
    }
}

print("Movies: \(movieCount), Songs: \(songCount)")

// Downcasting
for item in library {
    if let movie = item as? Movie {
        print("Movie: \(movie.name), dir. \(movie.director)")
    } else if let song = item as? Song {
        print("Song: \(song.name), by \(song.artist)")
    }
}

// Any and AnyObject
var things: [Any] = []
things.append(0)
things.append(0.0)
things.append(42)
things.append(3.14159)
things.append("hello")
things.append((3.0, 5.0))
things.append(Movie(name: "Ghostbusters", director: "Ivan Reitman"))

for thing in things {
    switch thing {
    case let someInt as Int:
        print("Int: \(someInt)")
    case let someDouble as Double:
        print("Double: \(someDouble)")
    case let someString as String:
        print("String: \(someString)")
    case let (x, y) as (Double, Double):
        print("Point: (\(x), \(y))")
    case let movie as Movie:
        print("Movie: \(movie.name)")
    default:
        print("Something else")
    }
}
```

### Opaque Types

```swift
protocol Shape {
    func draw() -> String
}

struct Square: Shape {
    var size: Int
    func draw() -> String {
        return "Square (\(size)x\(size))"
    }
}

struct Circle: Shape {
    var radius: Int
    func draw() -> String {
        return "Circle (r=\(radius))"
    }
}

// Opaque return type
func makeShape(isCircle: Bool) -> some Shape {
    if isCircle {
        return Circle(radius: 10)
    } else {
        return Square(size: 10)
    }
}

let shape = makeShape(isCircle: true)
print(shape.draw())

// Without opaque type (error: different return types)
// func makeShapeProtocol(isCircle: Bool) -> Shape {
//     if isCircle {
//         return Circle(radius: 10)
//     } else {
//         return Square(size: 10)
//     }
// }
```

### KeyPaths

```swift
struct Person {
    var name: String
    var age: Int
}

let alice = Person(name: "Alice", age: 30)

// KeyPath
let nameKeyPath = \Person.name
let ageKeyPath = \Person.age

print(alice[keyPath: nameKeyPath])  // "Alice"
print(alice[keyPath: ageKeyPath])   // 30

// WritableKeyPath
var bob = Person(name: "Bob", age: 25)
bob[keyPath: nameKeyPath] = "Robert"
print(bob.name)  // "Robert"

// Using KeyPaths with functions
func getValue<T, V>(from object: T, keyPath: KeyPath<T, V>) -> V {
    return object[keyPath: keyPath]
}

let name = getValue(from: alice, keyPath: \.name)
print(name)

// KeyPath in collections
let people = [
    Person(name: "Alice", age: 30),
    Person(name: "Bob", age: 25),
    Person(name: "Charlie", age: 35)
]

let names = people.map(\.name)
let ages = people.map(\.age)
let sorted = people.sorted(by: \.age, <)
```

### Result Builders

```swift
@resultBuilder
struct StringBuilder {
    static func buildBlock(_ components: String...) -> String {
        components.joined(separator: "\n")
    }

    static func buildOptional(_ component: String?) -> String {
        component ?? ""
    }

    static func buildEither(first component: String) -> String {
        component
    }

    static func buildEither(second component: String) -> String {
        component
    }

    static func buildArray(_ components: [String]) -> String {
        components.joined(separator: "\n")
    }
}

@StringBuilder
func makeGreeting(name: String) -> String {
    "Hello, \(name)!"
    "Welcome to Swift."
    "Enjoy your stay."
}

print(makeGreeting(name: "Alice"))

// With conditionals
@StringBuilder
func makeReport(score: Int) -> String {
    "Score Report"
    "Score: \(score)"

    if score >= 90 {
        "Grade: A"
    } else if score >= 80 {
        "Grade: B"
    } else {
        "Grade: C"
    }
}

print(makeReport(score: 95))
```

---

## Console/TUI Development

### Reading User Input

```swift
import Foundation

// Basic input
print("Enter your name: ", terminator: "")
if let name = readLine() {
    print("Hello, \(name)!")
}

// With validation
func getInteger(prompt: String) -> Int? {
    print(prompt, terminator: "")
    if let input = readLine(),
       let number = Int(input) {
        return number
    }
    return nil
}

if let age = getInteger(prompt: "Enter your age: ") {
    print("You are \(age) years old")
}

// Password input (hidden)
func getPassword(prompt: String) -> String {
    print(prompt, terminator: "")
    var password = ""
    // Note: This doesn't actually hide input in standard Swift
    // Would need termios for true hidden input
    if let input = readLine() {
        password = input
    }
    return password
}
```

### Command Line Arguments

```swift
import Foundation

// Access command line arguments
let arguments = CommandLine.arguments

print("Program: \(arguments[0])")

if arguments.count > 1 {
    print("Arguments:")
    for (index, arg) in arguments.dropFirst().enumerated() {
        print("  \(index + 1): \(arg)")
    }
}

// Simple argument parser
struct Arguments {
    var name: String?
    var age: Int?
    var verbose: Bool = false

    init(args: [String]) {
        var i = 1
        while i < args.count {
            let arg = args[i]

            switch arg {
            case "-n", "--name":
                if i + 1 < args.count {
                    name = args[i + 1]
                    i += 1
                }
            case "-a", "--age":
                if i + 1 < args.count {
                    age = Int(args[i + 1])
                    i += 1
                }
            case "-v", "--verbose":
                verbose = true
            case "-h", "--help":
                printHelp()
                exit(0)
            default:
                print("Unknown argument: \(arg)")
            }

            i += 1
        }
    }

    func printHelp() {
        print("""
        Usage: program [options]

        Options:
          -n, --name NAME      Specify name
          -a, --age AGE        Specify age
          -v, --verbose        Enable verbose output
          -h, --help           Show this help
        """)
    }
}

let parsedArgs = Arguments(args: CommandLine.arguments)

if parsedArgs.verbose {
    print("Verbose mode enabled")
    print("Name: \(parsedArgs.name ?? "not provided")")
    print("Age: \(parsedArgs.age?.description ?? "not provided")")
}

// Usage: swift program.swift -n Alice -a 30 -v
```

### Creating a Simple TUI

```swift
import Foundation

// Clear screen
func clearScreen() {
    print("\u{001B}[2J")
    print("\u{001B}[H")
}

// Colors
enum Color: String {
    case black = "\u{001B}[0;30m"
    case red = "\u{001B}[0;31m"
    case green = "\u{001B}[0;32m"
    case yellow = "\u{001B}[0;33m"
    case blue = "\u{001B}[0;34m"
    case magenta = "\u{001B}[0;35m"
    case cyan = "\u{001B}[0;36m"
    case white = "\u{001B}[0;37m"
    case reset = "\u{001B}[0m"
}

func colorize(_ text: String, color: Color) -> String {
    return "\(color.rawValue)\(text)\(Color.reset.rawValue)"
}

// Box drawing
func drawBox(width: Int, height: Int, title: String = "") {
    let topBorder = "┌" + String(repeating: "─", count: width - 2) + "┐"
    let bottomBorder = "└" + String(repeating: "─", count: width - 2) + "┘"
    let emptyLine = "│" + String(repeating: " ", count: width - 2) + "│"

    var titleLine = emptyLine
    if !title.isEmpty {
        let padding = (width - 2 - title.count) / 2
        let spaces = String(repeating: " ", count: padding)
        titleLine = "│\(spaces)\(title)\(spaces)"
        if title.count % 2 == 1 {
            titleLine += " "
        }
        titleLine += "│"
    }

    print(topBorder)
    if !title.isEmpty {
        print(titleLine)
        print("│" + String(repeating: "─", count: width - 2) + "│")
    }
    for _ in 0..<(height - 3) {
        print(emptyLine)
    }
    print(bottomBorder)
}

// Progress bar
func drawProgressBar(progress: Double, width: Int = 40) {
    let filled = Int(progress * Double(width))
    let empty = width - filled
    let bar = String(repeating: "█", count: filled) + String(repeating: "░", count: empty)
    let percentage = Int(progress * 100)
    print("[\(bar)] \(percentage)%")
}

// Menu
func showMenu(title: String, options: [String]) -> Int {
    clearScreen()
    print(colorize("=== \(title) ===", color: .cyan))
    print()

    for (index, option) in options.enumerated() {
        print("\(index + 1). \(option)")
    }

    print()
    print("Enter choice: ", terminator: "")

    if let input = readLine(),
       let choice = Int(input),
       choice > 0 && choice <= options.count {
        return choice
    }

    return 0
}

// Simple TUI application
struct TodoApp {
    var todos: [String] = []

    mutating func run() {
        var running = true

        while running {
            let choice = showMenu(
                title: "Todo List",
                options: [
                    "View todos",
                    "Add todo",
                    "Remove todo",
                    "Exit"
                ]
            )

            switch choice {
            case 1:
                viewTodos()
            case 2:
                addTodo()
            case 3:
                removeTodo()
            case 4:
                running = false
            default:
                print("Invalid choice!")
                sleep(1)
            }
        }

        print("Goodbye!")
    }

    func viewTodos() {
        clearScreen()
        print(colorize("=== Your Todos ===", color: .green))
        print()

        if todos.isEmpty {
            print("No todos yet!")
        } else {
            for (index, todo) in todos.enumerated() {
                print("\(index + 1). \(todo)")
            }
        }

        print("\nPress Enter to continue...")
        _ = readLine()
    }

    mutating func addTodo() {
        clearScreen()
        print(colorize("=== Add Todo ===", color: .yellow))
        print()
        print("Enter todo: ", terminator: "")

        if let todo = readLine(), !todo.isEmpty {
            todos.append(todo)
            print(colorize("✓ Todo added!", color: .green))
        } else {
            print(colorize("✗ Invalid input!", color: .red))
        }

        sleep(1)
    }

    mutating func removeTodo() {
        clearScreen()
        print(colorize("=== Remove Todo ===", color: .red))
        print()

        if todos.isEmpty {
            print("No todos to remove!")
            sleep(1)
            return
        }

        for (index, todo) in todos.enumerated() {
            print("\(index + 1). \(todo)")
        }

        print("\nEnter number to remove: ", terminator: "")

        if let input = readLine(),
           let index = Int(input),
           index > 0 && index <= todos.count {
            let removed = todos.remove(at: index - 1)
            print(colorize("✓ Removed: \(removed)", color: .green))
        } else {
            print(colorize("✗ Invalid number!", color: .red))
        }

        sleep(1)
    }
}

// Run the app
// var app = TodoApp()
// app.run()
```

### Working with Process and Pipes

```swift
import Foundation

// Run shell command
func runCommand(_ command: String, args: [String] = []) -> String {
    let process = Process()
    let pipe = Pipe()

    process.executableURL = URL(fileURLWithPath: "/usr/bin/env")
    process.arguments = [command] + args
    process.standardOutput = pipe

    do {
        try process.run()
        process.waitUntilExit()

        let data = pipe.fileHandleForReading.readDataToEndOfFile()
        if let output = String(data: data, encoding: .utf8) {
            return output.trimmingCharacters(in: .newlines)
        }
    } catch {
        print("Error running command: \(error)")
    }

    return ""
}

// Example usage
let output = runCommand("ls", args: ["-la"])
print(output)

// Get system information
func getSystemInfo() -> [String: String] {
    var info: [String: String] = [:]

    info["hostname"] = runCommand("hostname")
    info["user"] = runCommand("whoami")
    info["date"] = runCommand("date")
    info["uptime"] = runCommand("uptime")

    return info
}

let sysInfo = getSystemInfo()
for (key, value) in sysInfo {
    print("\(key): \(value)")
}

// Interactive process
func interactiveShell() {
    let process = Process()
    process.executableURL = URL(fileURLWithPath: "/bin/bash")
    process.arguments = ["-l"]

    let inputPipe = Pipe()
    let outputPipe = Pipe()

    process.standardInput = inputPipe
    process.standardOutput = outputPipe

    do {
        try process.run()

        // Send command
        let command = "echo 'Hello from Swift'\n"
        inputPipe.fileHandleForWriting.write(command.data(using: .utf8)!)

        // Read output
        let data = outputPipe.fileHandleForReading.availableData
        if let output = String(data: data, encoding: .utf8) {
            print(output)
        }

        process.terminate()
    } catch {
        print("Error: \(error)")
    }
}
```

### Timer and Scheduling

```swift
import Foundation

// Simple timer
func startTimer(duration: TimeInterval, interval: TimeInterval = 1.0) {
    var elapsed: TimeInterval = 0

    while elapsed < duration {
        print("Elapsed: \(Int(elapsed))s / \(Int(duration))s")
        Thread.sleep(forTimeInterval: interval)
        elapsed += interval
    }

    print("Timer finished!")
}

// Countdown timer
func countdown(from seconds: Int) {
    for i in (0...seconds).reversed() {
        clearScreen()
        print(colorize("=== Countdown ===", color: .cyan))
        print()
        print(colorize("\(i)", color: .yellow))
        sleep(1)
    }

    print(colorize("Time's up!", color: .red))
}

// Progress bar with timer
func processWithProgress(duration: TimeInterval, steps: Int = 100) {
    let stepDuration = duration / Double(steps)

    for i in 0...steps {
        let progress = Double(i) / Double(steps)
        clearScreen()
        print("Processing...")
        print()
        drawProgressBar(progress: progress)
        Thread.sleep(forTimeInterval: stepDuration)
    }

    print()
    print(colorize("✓ Complete!", color: .green))
}

// Usage examples:
// startTimer(duration: 10)
// countdown(from: 5)
// processWithProgress(duration: 5.0)
```

---

## Best Practices and Style Guide

### Naming Conventions

```swift
// Types: PascalCase
struct PersonInfo { }
class DatabaseManager { }
enum Direction { }
protocol Drawable { }

// Variables and functions: camelCase
let firstName = "John"
var itemCount = 0
func calculateTotal() { }

// Constants: Can be camelCase or SCREAMING_SNAKE_CASE
let maxConnections = 100
let API_KEY = "secret"

// Private members: Optional underscore prefix
private var _internalState = 0

// Boolean variables: Use "is", "has", "should" prefix
let isActive = true
let hasPermission = false
let shouldUpdate = true
```

### Code Organization

```swift
// 1. Imports
import Foundation
import SwiftUI

// 2. Type definitions
struct User {
    // 3. Properties
    let id: String
    var name: String
    var email: String

    // 4. Initializers
    init(id: String, name: String, email: String) {
        self.id = id
        self.name = name
        self.email = email
    }

    // 5. Methods
    func displayName() -> String {
        return "\(name) <\(email)>"
    }
}

// 6. Extensions (separate concerns)
extension User: Codable { }

extension User: CustomStringConvertible {
    var description: String {
        return displayName()
    }
}

// 7. Protocol conformance in extensions
protocol Identifiable {
    var id: String { get }
}

extension User: Identifiable { }
```

### Documentation

```swift
/// A structure representing a user in the system.
///
/// Use this structure to store and manage user information.
///
/// Example:
/// ```swift
/// let user = User(id: "123", name: "Alice", email: "alice@example.com")
/// print(user.displayName())
/// ```
struct User {
    /// The unique identifier for the user.
    let id: String

    /// The user's full name.
    var name: String

    /// The user's email address.
    var email: String

    /// Creates a new user with the specified information.
    ///
    /// - Parameters:
    ///   - id: The unique identifier
    ///   - name: The user's name
    ///   - email: The user's email address
    init(id: String, name: String, email: String) {
        self.id = id
        self.name = name
        self.email = email
    }

    /// Returns a formatted display name with email.
    ///
    /// - Returns: A string in the format "Name <email>"
    func displayName() -> String {
        return "\(name) <\(email)>"
    }
}
```

### Error Handling

```swift
// Prefer Result type for clear error handling
func fetchData() -> Result<String, NetworkError> {
    // Implementation
    return .success("Data")
}

// Use guard for early returns
func process(data: String?) {
    guard let validData = data else {
        print("No data provided")
        return
    }
    // Process validData
}

// Use defer for cleanup
func readFile() throws {
    let file = openFile()
    defer {
        closeFile(file)
    }
    // Process file
}
```

### Performance Tips

```swift
// Use lazy properties for expensive computations
struct DataProcessor {
    let data: [Int]

    lazy var sortedData: [Int] = {
        return data.sorted()
    }()
}

// Prefer value types (structs) over reference types (classes)
// when you don't need inheritance or reference semantics

// Use copy-on-write for large structs
struct LargeDataSet {
    private var storage: [String]

    init(storage: [String]) {
        self.storage = storage
    }

    mutating func append(_ item: String) {
        if !isKnownUniquelyReferenced(&storage) {
            storage = storage  // Make a copy
        }
        storage.append(item)
    }
}
```

---

## Conclusion

This course has covered Swift programming from basics to advanced topics, with a focus on console and TUI application development. You now have the knowledge to:

- Build command-line tools and utilities
- Create interactive terminal applications
- Work with files, JSON, and system processes
- Handle errors gracefully
- Write clean, idiomatic Swift code
- Manage memory effectively
- Use advanced Swift features

### Next Steps

1. **Practice** - Build small CLI tools and gradually increase complexity
2. **Explore Libraries** - Look into popular Swift packages:
   - ArgumentParser - Command-line argument parsing
   - SwiftLog - Logging framework
   - AsyncHTTPClient - Async HTTP requests
   - SwiftNIO - Non-blocking I/O
3. **Learn Server-Side Swift** - Vapor, Kitura for web services
4. **Study Concurrency** - async/await, actors, structured concurrency
5. **Contribute to Open Source** - Join Swift community projects

### Resources

- Official Swift Documentation: swift.org/documentation
- Swift Forums: forums.swift.org
- Swift Evolution: github.com/apple/swift-evolution
- Swift Package Index: swiftpackageindex.com
- Awesome Swift: github.com/matteocrippa/awesome-swift

Happy Swift coding!
