## Hello World
Basic program structure and main function
```swift
import Foundation

func main() {
    print("Hello, World!")
}

main()
```

----------

## Variable Declaration
Different ways to declare and initialize variables
```swift
// Immutable variables (constants)
let name = "John"
let age = 30
let city: String = "New York"

// Mutable variables
var counter = 0
var temperature = 20.5

counter += 1

// Multiple declarations
let x = 10, y = 20, z = 30

// Type inference
let message = "Swift is awesome"  // String
let pi = 3.14159                  // Double

print("Name: \(name), Age: \(age), Counter: \(counter)")
```

----------

## Constants and Variables
Understanding let vs var
```swift
// Constants (immutable)
let MAX_USERS = 100
let PI = 3.14159
let GREETING = "Hello"

// Variables (mutable)
var score = 0
var playerName = "Player1"

score += 10
playerName = "NewPlayer"

print("Pi: \(PI), Max Users: \(MAX_USERS)")
print("Score: \(score), Player: \(playerName)")
```

----------

## Data Types
Working with basic data types
```swift
// Integer types
let integer8: Int8 = 127
let integer16: Int16 = 32767
let integer32: Int32 = 2147483647
let integer64: Int64 = 9223372036854775807
let unsignedInt: UInt = 100

// Floating point
let float: Float = 3.14
let double: Double = 3.14159265359

// Boolean
let isActive = true
let isComplete: Bool = false

// Character and String
let character: Character = "A"
let emoji: Character = "🚀"
let text = "Hello, Swift!"

// Tuple
let coordinates = (x: 10, y: 20)
let person = (name: "Alice", age: 25)

print("Int: \(integer32), Double: \(double)")
print("Bool: \(isActive), Char: \(character)")
print("Coordinates: \(coordinates.x), \(coordinates.y)")
print("Person: \(person.name), \(person.age)")
```

----------

## Arrays and Collections
Working with arrays
```swift
// Array declaration
var numbers = [1, 2, 3, 4, 5]
var names: [String] = ["Alice", "Bob", "Charlie"]
var emptyArray = [Int]()

// Array operations
numbers.append(6)
numbers.insert(0, at: 0)
let first = numbers.removeFirst()
let last = numbers.removeLast()

// Array methods
let count = numbers.count
let isEmpty = numbers.isEmpty
let contains = numbers.contains(3)

// Iterating
for number in numbers {
    print("Number: \(number)")
}

for (index, name) in names.enumerated() {
    print("\(index): \(name)")
}

// Array slicing
let slice = numbers[1...3]
print("Slice: \(slice)")
```

----------

## Dictionaries
Creating and manipulating dictionaries
```swift
// Dictionary declaration
var person: [String: Any] = [
    "name": "Alice",
    "age": 30,
    "city": "New York"
]

var scores = ["Alice": 95, "Bob": 87, "Charlie": 92]

// Dictionary operations
person["email"] = "alice@example.com"
let age = person["age"] as? Int
person.removeValue(forKey: "city")

// Dictionary methods
let keys = Array(person.keys)
let values = Array(person.values)
let count = person.count

// Iterating
for (key, value) in person {
    print("\(key): \(value)")
}

// Default values
let score = scores["Diana", default: 0]
print("Diana's score: \(score)")
```

----------

## Sets
Working with unique collections
```swift
// Set declaration
var fruits: Set<String> = ["apple", "banana", "orange"]
var numbers: Set<Int> = [1, 2, 3, 4, 5]

// Set operations
fruits.insert("grape")
fruits.remove("banana")
let contains = fruits.contains("apple")

// Set operations
let set1: Set = [1, 2, 3, 4, 5]
let set2: Set = [4, 5, 6, 7, 8]

let union = set1.union(set2)
let intersection = set1.intersection(set2)
let difference = set1.subtracting(set2)
let symmetricDiff = set1.symmetricDifference(set2)

print("Union: \(union)")
print("Intersection: \(intersection)")
print("Difference: \(difference)")
```

----------

## Control Flow - If/Else
Conditional statements
```swift
let age = 25
let score = 85

// Basic if-else
if age >= 18 {
    print("Adult")
} else {
    print("Minor")
}

// If-else if-else
let grade: String
if score >= 90 {
    grade = "A"
} else if score >= 80 {
    grade = "B"
} else if score >= 70 {
    grade = "C"
} else {
    grade = "F"
}

// Ternary operator
let status = age >= 18 ? "adult" : "minor"

// Multiple conditions
if age >= 18 && score >= 80 {
    print("Eligible for scholarship")
}

print("Status: \(status), Grade: \(grade)")
```

----------

## Switch Statements
Pattern matching with switch
```swift
let number = 3
let point = (x: 2, y: 3)

// Basic switch
switch number {
case 1:
    print("One")
case 2, 3:
    print("Two or Three")
case 4...10:
    print("Four to Ten")
default:
    print("Something else")
}

// Switch with tuples
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
    print("Somewhere else")
}

// Switch with value binding
let value = (1, "hello")
switch value {
case (0, let text):
    print("Zero with text: \(text)")
case (let number, "hello"):
    print("Hello with number: \(number)")
default:
    break
}
```

----------

## Loops
Different types of loops
```swift
// For-in loop with range
for i in 0..<5 {
    print("Number: \(i)")
}

// For-in loop with array
let fruits = ["apple", "banana", "orange"]
for fruit in fruits {
    print("Fruit: \(fruit)")
}

// For-in with enumerated
for (index, fruit) in fruits.enumerated() {
    print("\(index): \(fruit)")
}

// While loop
var counter = 0
while counter < 3 {
    print("Counter: \(counter)")
    counter += 1
}

// Repeat-while loop (do-while)
var number = 0
repeat {
    print("Number: \(number)")
    number += 1
} while number < 3

// Stride for custom steps
for i in stride(from: 0, to: 10, by: 2) {
    print("Even: \(i)")
}
```

----------

## Functions
Function declaration and parameters
```swift
// Basic function
func greet() {
    print("Hello, World!")
}

// Function with parameters
func greet(name: String) {
    print("Hello, \(name)!")
}

// Function with return value
func add(a: Int, b: Int) -> Int {
    return a + b
}

// Function with external parameter names
func greet(person name: String, from hometown: String) {
    print("Hello \(name) from \(hometown)!")
}

// Function with default parameters
func greet(name: String, greeting: String = "Hello") {
    print("\(greeting), \(name)!")
}

// Variadic parameters
func sum(_ numbers: Int...) -> Int {
    return numbers.reduce(0, +)
}

// Multiple return values (tuple)
func minMax(array: [Int]) -> (min: Int, max: Int)? {
    guard let first = array.first else { return nil }
    var min = first
    var max = first

    for value in array {
        if value < min { min = value }
        if value > max { max = value }
    }

    return (min, max)
}

// Usage
greet()
greet(name: "Alice")
let result = add(a: 5, b: 3)
greet(person: "Bob", from: "Boston")
let total = sum(1, 2, 3, 4, 5)

if let bounds = minMax(array: [1, 5, 3, 9, 2]) {
    print("Min: \(bounds.min), Max: \(bounds.max)")
}
```

----------

## Closures
Working with closures (lambdas)
```swift
// Basic closure
let greet = { (name: String) -> String in
    return "Hello, \(name)!"
}

// Closure with inferred types
let add: (Int, Int) -> Int = { a, b in
    return a + b
}

// Shorthand closure
let multiply = { $0 * $1 }

// Closures with arrays
let numbers = [1, 2, 3, 4, 5]

let doubled = numbers.map { $0 * 2 }
let evens = numbers.filter { $0 % 2 == 0 }
let sum = numbers.reduce(0) { $0 + $1 }

// Trailing closure syntax
func performOperation(_ operation: (Int, Int) -> Int, on a: Int, and b: Int) -> Int {
    return operation(a, b)
}

let result1 = performOperation({ $0 + $1 }, on: 5, and: 3)
let result2 = performOperation(on: 5, and: 3) { $0 * $1 }

// Capturing values
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

print("Doubled: \(doubled)")
print("Evens: \(evens)")
print("Sum: \(sum)")
```

----------

## Structs
Defining and using structures
```swift
struct Person {
    var name: String
    var age: Int

    // Method
    func greet() -> String {
        return "Hello, I'm \(name)"
    }

    // Mutating method
    mutating func haveBirthday() {
        age += 1
    }

    // Static method
    static func defaultPerson() -> Person {
        return Person(name: "Unknown", age: 0)
    }
}

// Rectangle with computed properties
struct Rectangle {
    var width: Double
    var height: Double

    var area: Double {
        return width * height
    }

    var perimeter: Double {
        return 2 * (width + height)
    }

    init(width: Double, height: Double) {
        self.width = width
        self.height = height
    }
}

// Usage
var person = Person(name: "Alice", age: 30)
print(person.greet())
person.haveBirthday()
print("New age: \(person.age)")

let defaultPerson = Person.defaultPerson()
print("Default: \(defaultPerson.name)")

let rect = Rectangle(width: 10, height: 5)
print("Area: \(rect.area), Perimeter: \(rect.perimeter)")
```

----------

## Classes
Object-oriented programming with classes
```swift
class Animal {
    var name: String
    var species: String

    init(name: String, species: String) {
        self.name = name
        self.species = species
    }

    func makeSound() -> String {
        return "Some generic animal sound"
    }

    func getInfo() -> String {
        return "\(name) is a \(species)"
    }

    deinit {
        print("\(name) is being deallocated")
    }
}

class Dog: Animal {
    var breed: String

    init(name: String, breed: String) {
        self.breed = breed
        super.init(name: name, species: "Dog")
    }

    override func makeSound() -> String {
        return "Woof!"
    }

    func fetch() -> String {
        return "\(name) is fetching the ball!"
    }
}

class Cat: Animal {
    var isIndoor: Bool

    init(name: String, isIndoor: Bool = true) {
        self.isIndoor = isIndoor
        super.init(name: name, species: "Cat")
    }

    override func makeSound() -> String {
        return "Meow!"
    }
}

// Usage
let dog = Dog(name: "Buddy", breed: "Golden Retriever")
let cat = Cat(name: "Whiskers")

print(dog.getInfo())
print(dog.makeSound())
print(dog.fetch())

print(cat.getInfo())
print(cat.makeSound())
```

----------

## Enums
Defining and using enumerations
```swift
// Basic enum
enum Direction {
    case north
    case south
    case east
    case west
}

// Enum with associated values
enum Barcode {
    case upc(Int, Int, Int, Int)
    case qrCode(String)
}

// Enum with raw values
enum Planet: Int {
    case mercury = 1
    case venus
    case earth
    case mars
}

// Enum with methods
enum CompassPoint: String {
    case north = "N"
    case south = "S"
    case east = "E"
    case west = "W"

    func opposite() -> CompassPoint {
        switch self {
        case .north: return .south
        case .south: return .north
        case .east: return .west
        case .west: return .east
        }
    }
}

// Usage
let direction = Direction.north

let productBarcode = Barcode.upc(8, 85909, 51226, 3)
let qrBarcode = Barcode.qrCode("ABCDEFG")

// Switch with associated values
switch productBarcode {
case .upc(let a, let b, let c, let d):
    print("UPC: \(a), \(b), \(c), \(d)")
case .qrCode(let code):
    print("QR Code: \(code)")
}

let earth = Planet.earth
print("Earth's position: \(earth.rawValue)")

let point = CompassPoint.north
print("North opposite: \(point.opposite().rawValue)")
```

----------

## Optionals
Working with optional values
```swift
// Optional declaration
var name: String? = "Alice"
var age: Int? = nil

// Optional binding (if let)
if let unwrappedName = name {
    print("Name is \(unwrappedName)")
} else {
    print("Name is nil")
}

// Optional binding with multiple conditions
if let name = name, let age = age {
    print("\(name) is \(age) years old")
}

// Guard statement
func greet(person: String?) {
    guard let name = person else {
        print("No name provided")
        return
    }
    print("Hello, \(name)!")
}

// Nil coalescing operator
let displayName = name ?? "Guest"
print("Display name: \(displayName)")

// Optional chaining
struct Address {
    var street: String
    var city: String
}

struct Person {
    var name: String
    var address: Address?
}

let person = Person(name: "Bob", address: nil)
let city = person.address?.city ?? "Unknown"
print("City: \(city)")

// Forced unwrapping (use with caution!)
let forcedName = name!  // Crashes if nil

// Implicitly unwrapped optionals
var assumedString: String! = "An implicitly unwrapped optional string"
print(assumedString)  // No need to unwrap
```

----------

## Error Handling
Try-catch blocks and error handling
```swift
enum NetworkError: Error {
    case badURL
    case requestFailed
    case invalidResponse
    case decodingError
}

enum ValidationError: Error {
    case emptyString
    case tooShort
    case tooLong
}

// Function that throws
func validatePassword(_ password: String) throws -> Bool {
    if password.isEmpty {
        throw ValidationError.emptyString
    }
    if password.count < 6 {
        throw ValidationError.tooShort
    }
    if password.count > 20 {
        throw ValidationError.tooLong
    }
    return true
}

func fetchData(from url: String) throws -> String {
    if url.isEmpty {
        throw NetworkError.badURL
    }
    // Simulate network call
    return "Sample data"
}

// Using do-catch
do {
    try validatePassword("abc")
    print("Password is valid")
} catch ValidationError.emptyString {
    print("Password cannot be empty")
} catch ValidationError.tooShort {
    print("Password is too short")
} catch ValidationError.tooLong {
    print("Password is too long")
} catch {
    print("Unknown error: \(error)")
}

// Try? (converts to optional)
let data = try? fetchData(from: "https://example.com")
print("Data: \(data ?? "nil")")

// Try! (force try - crashes on error)
// let forcedData = try! fetchData(from: "")  // Crashes

// Defer statement
func processFile(filename: String) {
    print("Opening file")
    defer {
        print("Closing file")  // Executed when function exits
    }
    print("Processing file")
}

processFile(filename: "test.txt")
```

----------

## File I/O
Reading and writing files
```swift
import Foundation

// Writing to file
let content = "Hello, Swift!\nThis is a test file."
let filename = "test.txt"

do {
    try content.write(toFile: filename, atomically: true, encoding: .utf8)
    print("File written successfully")
} catch {
    print("Error writing file: \(error)")
}

// Reading from file
do {
    let readContent = try String(contentsOfFile: filename, encoding: .utf8)
    print("File content:\n\(readContent)")
} catch {
    print("Error reading file: \(error)")
}

// Working with FileManager
let fileManager = FileManager.default
let currentPath = fileManager.currentDirectoryPath

// Check if file exists
if fileManager.fileExists(atPath: filename) {
    print("File exists at: \(currentPath)/\(filename)")
}

// Create directory
let directoryName = "testDir"
do {
    try fileManager.createDirectory(atPath: directoryName,
                                    withIntermediateDirectories: true)
    print("Directory created")
} catch {
    print("Error creating directory: \(error)")
}

// List directory contents
do {
    let contents = try fileManager.contentsOfDirectory(atPath: ".")
    print("Directory contents: \(contents)")
} catch {
    print("Error listing directory: \(error)")
}

// Clean up
try? fileManager.removeItem(atPath: filename)
try? fileManager.removeItem(atPath: directoryName)
```

----------

## JSON Handling
Working with JSON data
```swift
import Foundation

// Codable struct
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
    }

    // Decoding from JSON
    let decoder = JSONDecoder()
    let decodedPerson = try decoder.decode(Person.self, from: jsonData)
    print("\nDecoded: \(decodedPerson.name), \(decodedPerson.age)")

} catch {
    print("Error: \(error)")
}

// Manual JSON parsing
let jsonString = """
{
    "name": "Bob",
    "age": 25,
    "scores": [85, 90, 95]
}
"""

if let jsonData = jsonString.data(using: .utf8) {
    do {
        if let json = try JSONSerialization.jsonObject(with: jsonData) as? [String: Any] {
            let name = json["name"] as? String
            let age = json["age"] as? Int
            let scores = json["scores"] as? [Int]

            print("\nManual parsing:")
            print("Name: \(name ?? "nil")")
            print("Age: \(age ?? 0)")
            print("Scores: \(scores ?? [])")
        }
    } catch {
        print("Error parsing JSON: \(error)")
    }
}
```

----------

## String Manipulation
Common string operations
```swift
let text = "  Hello, World!  "
let name = "Alice"
let age = 30

// Basic operations
print("Original: '\(text)'")
print("Trimmed: '\(text.trimmingCharacters(in: .whitespaces))'")
print("Uppercase: \(text.uppercased())")
print("Lowercase: \(text.lowercased())")

// String interpolation
let message = "My name is \(name) and I'm \(age) years old"
print(message)

// String methods
let sentence = "Swift is awesome"
let words = sentence.split(separator: " ")
let contains = sentence.contains("Swift")
let hasPrefix = sentence.hasPrefix("Swift")
let hasSuffix = sentence.hasSuffix("awesome")

print("Words: \(words)")
print("Contains 'Swift': \(contains)")
print("Starts with 'Swift': \(hasPrefix)")

// Replacing
let replaced = sentence.replacingOccurrences(of: "awesome", with: "great")
print("Replaced: \(replaced)")

// String length
let length = sentence.count
print("Length: \(length)")

// Substring
let index = sentence.index(sentence.startIndex, offsetBy: 5)
let substring = sentence[..<index]
print("First 5 chars: \(substring)")

// Joining
let items = ["apple", "banana", "orange"]
let joined = items.joined(separator: ", ")
print("Joined: \(joined)")

// Multiline strings
let multiline = """
This is a
multiline
string
"""
print("Multiline:\n\(multiline)")
```

----------

## Regular Expressions
Pattern matching with regex
```swift
import Foundation

let text = "Contact us at info@company.com or support@company.com. Phone: (555) 123-4567"

// Email pattern
let emailPattern = #"[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}"#

do {
    let emailRegex = try NSRegularExpression(pattern: emailPattern)
    let matches = emailRegex.matches(in: text, range: NSRange(text.startIndex..., in: text))

    print("Found emails:")
    for match in matches {
        if let range = Range(match.range, in: text) {
            print("  - \(text[range])")
        }
    }
} catch {
    print("Invalid regex: \(error)")
}

// Phone pattern
let phonePattern = #"\((\d{3})\)\s(\d{3})-(\d{4})"#

do {
    let phoneRegex = try NSRegularExpression(pattern: phonePattern)
    if let match = phoneRegex.firstMatch(in: text, range: NSRange(text.startIndex..., in: text)) {
        if let range = Range(match.range, in: text) {
            print("\nPhone found: \(text[range])")
        }
    }
} catch {
    print("Invalid regex: \(error)")
}

// Swift 5.7+ native regex (if available)
#if swift(>=5.7)
if #available(macOS 13.0, iOS 16.0, *) {
    let emailRegex = /[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/
    let emailMatches = text.matches(of: emailRegex)
    print("\nNative regex emails:")
    for match in emailMatches {
        print("  - \(match.output)")
    }
}
#endif
```

----------

## Extensions
Extending existing types
```swift
// Extend Int
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

// Extend String
extension String {
    func reverse() -> String {
        return String(self.reversed())
    }

    var wordCount: Int {
        return self.split(separator: " ").count
    }

    func truncate(length: Int) -> String {
        if self.count > length {
            let index = self.index(self.startIndex, offsetBy: length)
            return String(self[..<index]) + "..."
        }
        return self
    }
}

// Extend Array
extension Array where Element: Numeric {
    func sum() -> Element {
        return self.reduce(0, +)
    }

    func average() -> Double where Element == Int {
        return isEmpty ? 0 : Double(sum()) / Double(count)
    }
}

// Usage
let number = 5
print("5 squared: \(number.squared())")
print("5 is even: \(number.isEven)")

3.times {
    print("Hello!")
}

let text = "Hello, Swift!"
print("Reversed: \(text.reverse())")
print("Word count: \(text.wordCount)")
print("Truncated: \(text.truncate(length: 8))")

let numbers = [1, 2, 3, 4, 5]
print("Sum: \(numbers.sum())")
print("Average: \(numbers.average())")
```

----------

## Protocols
Defining and conforming to protocols
```swift
// Basic protocol
protocol Drawable {
    func draw() -> String
}

// Protocol with properties
protocol Named {
    var name: String { get set }
    var fullName: String { get }
}

// Protocol with initializer
protocol Identifiable {
    var id: String { get }
    init(id: String)
}

// Conforming to protocols
struct Circle: Drawable {
    var radius: Double

    func draw() -> String {
        return "Drawing a circle with radius \(radius)"
    }
}

struct Rectangle: Drawable {
    var width: Double
    var height: Double

    func draw() -> String {
        return "Drawing a rectangle \(width)x\(height)"
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

// Protocol composition
protocol Printable {
    func printDescription()
}

struct Document: Named, Printable {
    var name: String
    var fullName: String {
        return "Document: \(name)"
    }

    func printDescription() {
        print(fullName)
    }
}

// Usage
let shapes: [Drawable] = [
    Circle(radius: 5.0),
    Rectangle(width: 10, height: 5)
]

for shape in shapes {
    print(shape.draw())
}

let person = Person(name: "John", lastName: "Doe")
print("Full name: \(person.fullName)")
```

----------

## Generics
Generic types and functions
```swift
// Generic function
func swapValues<T>(_ a: inout T, _ b: inout T) {
    let temp = a
    a = b
    b = temp
}

// Generic type
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

// Generic with constraints
func findIndex<T: Equatable>(of value: T, in array: [T]) -> Int? {
    for (index, item) in array.enumerated() {
        if item == value {
            return index
        }
    }
    return nil
}

// Generic protocol
protocol Container {
    associatedtype Item
    mutating func append(_ item: Item)
    var count: Int { get }
    subscript(i: Int) -> Item { get }
}

struct GenericArray<T>: Container {
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

// Usage
var a = 5, b = 10
swapValues(&a, &b)
print("Swapped: a=\(a), b=\(b)")

var stack = Stack<String>()
stack.push("First")
stack.push("Second")
stack.push("Third")
print("Stack count: \(stack.count)")
print("Popped: \(stack.pop() ?? "nil")")

let numbers = [1, 2, 3, 4, 5]
if let index = findIndex(of: 3, in: numbers) {
    print("Found 3 at index \(index)")
}
```

----------

## Property Wrappers
Custom property wrappers
```swift
// Property wrapper for clamping values
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
}

// Property wrapper for capitalization
@propertyWrapper
struct Capitalized {
    private var value: String

    init(wrappedValue: String) {
        self.value = wrappedValue.capitalized
    }

    var wrappedValue: String {
        get { value }
        set { value = newValue.capitalized }
    }
}

// Usage
struct Game {
    @Clamped(0...100) var health: Int = 100
    @Clamped(0...10) var level: Int = 1
    @Capitalized var playerName: String
}

var game = Game(playerName: "john doe")
print("Player: \(game.playerName)")  // John Doe

game.health = 150  // Will be clamped to 100
print("Health: \(game.health)")

game.health = -10  // Will be clamped to 0
print("Health after damage: \(game.health)")
```

----------

## Result Type
Handling success and failure
```swift
enum NetworkError: Error {
    case invalidURL
    case requestFailed
    case decodingError
}

// Function returning Result
func fetchUserData(id: Int) -> Result<String, NetworkError> {
    if id <= 0 {
        return .failure(.invalidURL)
    }

    // Simulate successful fetch
    return .success("User data for ID \(id)")
}

func parseJSON(_ data: String) -> Result<[String: Any], NetworkError> {
    if data.isEmpty {
        return .failure(.decodingError)
    }

    return .success(["name": "Alice", "age": 30])
}

// Using Result
let result1 = fetchUserData(id: 123)

switch result1 {
case .success(let data):
    print("Success: \(data)")
case .failure(let error):
    print("Error: \(error)")
}

// Result with map and flatMap
let result2 = fetchUserData(id: 456)
    .map { data in
        data.uppercased()
    }

// Get value or provide default
let userData = fetchUserData(id: 789).get(default: "No data")
print("User data: \(userData)")

// Chaining results
let finalResult = fetchUserData(id: 100)
    .flatMap { data in
        parseJSON(data)
    }

switch finalResult {
case .success(let json):
    print("Parsed JSON: \(json)")
case .failure(let error):
    print("Failed: \(error)")
}
```

----------

## Command Line Arguments
Processing command line arguments
```swift
import Foundation

// Access command line arguments
let arguments = CommandLine.arguments
print("Program name: \(arguments[0])")

if arguments.count > 1 {
    print("Arguments:")
    for (index, arg) in arguments.dropFirst().enumerated() {
        print("  \(index + 1): \(arg)")
    }
} else {
    print("No arguments provided")
}

// Simple argument parser
struct Arguments {
    var name: String?
    var age: Int?
    var verbose: Bool = false

    init(args: [String]) {
        var i = 1  // Skip program name
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
            default:
                break
            }

            i += 1
        }
    }
}

let parsedArgs = Arguments(args: CommandLine.arguments)

if parsedArgs.verbose {
    print("\nParsed arguments:")
    print("  Name: \(parsedArgs.name ?? "not provided")")
    print("  Age: \(parsedArgs.age?.description ?? "not provided")")
    print("  Verbose: \(parsedArgs.verbose)")
}

// Usage example:
// swift program.swift -n Alice -a 30 -v
```

----------

## Date and Time
Working with dates and times
```swift
import Foundation

// Current date and time
let now = Date()
print("Now: \(now)")

// Date components
let calendar = Calendar.current
let components = calendar.dateComponents([.year, .month, .day, .hour, .minute], from: now)

print("Year: \(components.year ?? 0)")
print("Month: \(components.month ?? 0)")
print("Day: \(components.day ?? 0)")

// Date formatting
let formatter = DateFormatter()
formatter.dateStyle = .medium
formatter.timeStyle = .medium
print("Formatted: \(formatter.string(from: now))")

// Custom format
formatter.dateFormat = "yyyy-MM-dd HH:mm:ss"
print("Custom format: \(formatter.string(from: now))")

// Date arithmetic
if let tomorrow = calendar.date(byAdding: .day, value: 1, to: now),
   let nextWeek = calendar.date(byAdding: .weekOfYear, value: 1, to: now),
   let oneHourAgo = calendar.date(byAdding: .hour, value: -1, to: now) {

    print("\nTomorrow: \(formatter.string(from: tomorrow))")
    print("Next week: \(formatter.string(from: nextWeek))")
    print("One hour ago: \(formatter.string(from: oneHourAgo))")
}

// Parse date string
formatter.dateFormat = "yyyy-MM-dd"
if let parsedDate = formatter.date(from: "2023-12-25") {
    print("\nParsed date: \(parsedDate)")
}

// Time interval
let timestamp = now.timeIntervalSince1970
print("Timestamp: \(timestamp)")

let futureDate = Date(timeIntervalSinceNow: 3600)  // 1 hour from now
print("One hour from now: \(formatter.string(from: futureDate))")
```
