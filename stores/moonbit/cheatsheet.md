# MoonBit Quick Reference Cheatsheet

A concise reference for MoonBit syntax and common patterns.

## Basic Syntax

### Variables

```moonbit
let x = 5                    // Immutable
let mut y = 10               // Mutable
let z : Int = 15             // With type annotation
const PI : Double = 3.14159  // Constant
```

### Data Types

```moonbit
// Primitives
let unit : Unit = ()
let bool : Bool = true
let int : Int = 42
let float : Double = 3.14
let char : Char = 'A'
let string : String = "Hello"
let bytes : Bytes = b"data"

// Collections
let array : Array[Int] = [1, 2, 3]
let tuple : (Int, String) = (1, "a")
let map : Map[String, Int] = Map::new()
```

### String Interpolation

```moonbit
let name = "Alice"
let age = 30
println("Hello, \{name}! You are \{age} years old")
```

## Functions

### Basic Functions

```moonbit
fn add(a : Int, b : Int) -> Int {
  a + b
}

fn greet(name : String) -> Unit {
  println("Hello, \{name}!")
}
```

### Labelled Arguments

```moonbit
fn create_user(~name : String, ~age : Int, ~email : String = "unknown") -> String {
  "User: \{name}"
}

// Call with labels
create_user(name="Alice", age=30)
create_user(age=25, name="Bob", email="bob@example.com")
```

### Anonymous Functions / Closures

```moonbit
let double = fn(x : Int) -> Int { x * 2 }
let add = fn(a, b) { a + b }  // Type inferred

// With captures
let multiplier = 10
let multiply_by_ten = fn(x) { x * multiplier }
```

### Partial Application

```moonbit
fn add(a : Int, b : Int) -> Int { a + b }

let add_five = add(5, _)
println(add_five(3))  // 8
```

## Control Flow

### If-Else

```moonbit
// As statement
if condition {
  // ...
} else if other_condition {
  // ...
} else {
  // ...
}

// As expression
let value = if condition { 1 } else { 2 }
```

### Match

```moonbit
match value {
  1 => "one"
  2 | 3 => "two or three"
  4..=10 => "four to ten"
  _ => "other"
}

// With destructuring
match (x, y) {
  (0, 0) => "origin"
  (x, 0) => "x-axis"
  (0, y) => "y-axis"
  (x, y) => "other"
}

// With guards
match x {
  n if n < 0 => "negative"
  0 => "zero"
  n if n > 0 => "positive"
  _ => "unknown"
}
```

### Loops

```moonbit
// Loop (infinite)
loop {
  if condition { break }
}

// While
while condition {
  // ...
}

// For
for i = 0; i < 10; i = i + 1 {
  println(i)
}

// For-in
for item in collection {
  println(item)
}
```

## Structs and Enums

### Structs

```moonbit
struct Person {
  name : String
  age : Int
}

// Create
let person = { name: "Alice", age: 30 }

// With constructor
fn Person::new(name : String, age : Int) -> Person {
  { name, age }
}

// Methods
fn greet(self : Person) -> String {
  "Hello, I'm \{self.name}"
}
```

### Enums

```moonbit
enum Color {
  Red
  Green
  Blue
  RGB(Int, Int, Int)
}

let red = Red
let custom = RGB(255, 128, 0)

match color {
  Red => "red"
  Green => "green"
  Blue => "blue"
  RGB(r, g, b) => "custom: \{r}, \{g}, \{b}"
}
```

### Option and Result

```moonbit
// Option
let some : Int? = Some(42)
let none : Int? = None

match optional {
  Some(x) => println("Value: \{x}")
  None => println("No value")
}

// Result
let ok : Result[Int, String] = Ok(42)
let err : Result[Int, String] = Err("error")

match result {
  Ok(value) => println("Success: \{value}")
  Err(error) => println("Error: \{error}")
}
```

## Error Handling

### Checked Exceptions

```moonbit
// Function that can raise errors
fn divide(a : Double, b : Double) -> Double!String {
  if b == 0.0 {
    raise "Division by zero"
  }
  a / b
}

// Try-catch
try {
  let result = divide!(10.0, 2.0)
  println(result)
} catch {
  error => println("Error: \{error}")
}

// Error propagation with !
fn calculate(a : Double, b : Double) -> Double!String {
  let x = divide!(a, b)
  let y = divide!(x, 2.0)
  y
}

// Convert to Result
let result : Result[Double, String] = divide(10.0, 0.0)
```

## Collections

### Arrays

```moonbit
let mut arr = [1, 2, 3]

arr.push(4)              // Add element
let last = arr.pop()     // Remove last
let first = arr[0]       // Access by index
let len = arr.length()   // Get length

arr[0] = 10             // Modify element
```

### Array Methods

```moonbit
let numbers = [1, 2, 3, 4, 5]

// Map
numbers.map(fn(x) { x * 2 })

// Filter
numbers.filter(fn(x) { x % 2 == 0 })

// Fold
numbers.fold(init=0, fn(acc, x) { acc + x })

// Each
numbers.each(fn(x) { println(x) })

// Chain operations
numbers.iter()
  .filter(fn(x) { x > 2 })
  .map(fn(x) { x * 2 })
  .collect()
```

### Maps

```moonbit
let mut map : Map[String, Int] = Map::new()

map.set("key", 42)        // Insert
let value = map.get("key") // Get (returns Option)
map.remove("key")         // Remove
let exists = map.contains("key")

// Iterate
map.iter().each(fn(entry) {
  let (key, value) = entry
  println("\{key}: \{value}")
})
```

## Pattern Matching

### Destructuring

```moonbit
// Tuples
let (x, y, z) = (1, 2, 3)

// Structs
let Person { name, age } = person

// Arrays
match arr {
  [] => "empty"
  [x] => "one element"
  [x, y] => "two elements"
  [first, .., last] => "first and last"
  _ => "other"
}
```

### JSON Patterns

```moonbit
match json {
  { "name": String(name), "age": Number(age) } => {
    println("\{name} is \{age}")
  }
  _ => println("invalid")
}
```

## Traits

### Define and Implement

```moonbit
trait Show {
  show(Self) -> String
}

struct Point { x : Int; y : Int }

fn show(self : Point) -> String {
  "(\{self.x}, \{self.y})"
}

// Use
let p = { x: 3, y: 4 }
println(p.show())
```

### Trait Bounds

```moonbit
fn print_all[T : Show](items : Array[T]) -> Unit {
  for item in items {
    println(item.show())
  }
}
```

## Generics

### Generic Functions

```moonbit
fn identity[T](x : T) -> T { x }

fn first[T](arr : Array[T]) -> T? {
  if arr.length() > 0 {
    Some(arr[0])
  } else {
    None
  }
}
```

### Generic Types

```moonbit
struct Box[T] {
  value : T
}

enum Option[T] {
  Some(T)
  None
}
```

## Advanced Features

### Pipe Operator

```moonbit
let result = value
  |> double
  |> add_ten
  |> square
```

### Cascade Operator

```moonbit
builder
  ..set_name("Alice")
  ..set_age(30)
  ..build()
```

### Defer

```moonbit
fn process() -> Unit {
  println("Start")
  defer { println("End") }
  println("Middle")
}
// Output: Start, Middle, End
```

## Testing

```moonbit
test "test name" {
  assert_eq!(1 + 1, 2)
  assert_ne!(1, 2)
  assert!(true)
}

test "error handling" {
  let result = risky_operation()
  match result {
    Ok(_) => ()
    Err(_) => fail!("Should succeed")
  }
}
```

## Common Patterns

### Builder Pattern

```moonbit
struct Config {
  mut host : String
  mut port : Int
}

fn Config::new() -> Config {
  { host: "localhost", port: 8080 }
}

fn with_host(self : Config, host : String) -> Config {
  self.host = host
  self
}

let config = Config::new()..with_host("example.com")
```

### Newtype Pattern

```moonbit
struct UserId(Int)
struct Email(String)

fn send_email(user_id : UserId, email : Email) -> Unit {
  // Type safety!
}
```

### Iterator Chaining

```moonbit
collection.iter()
  .filter(predicate)
  .map(transform)
  .take(n)
  .collect()
```

### Option Chaining

```moonbit
user
  .get_address()
  .and_then(fn(addr) { addr.get_zipcode() })
  .or("00000")
```

## Operators

### Arithmetic

```moonbit
a + b    // Addition
a - b    // Subtraction
a * b    // Multiplication
a / b    // Division
a % b    // Modulo
```

### Comparison

```moonbit
a == b   // Equal
a != b   // Not equal
a < b    // Less than
a > b    // Greater than
a <= b   // Less than or equal
a >= b   // Greater than or equal
```

### Logical

```moonbit
a && b   // AND
a || b   // OR
!a       // NOT
```

### Special

```moonbit
a |> b   // Pipe
a..b     // Cascade
[a, b, ..c]  // Spread
```

## Module System

### Visibility

```moonbit
fn private_function() {}        // Private
pub fn public_function() {}     // Public to package
pub(all) fn all_public() {}     // Public to all
```

### Imports

```moonbit
use module::function
use module::{func1, func2}
use module::name as alias
```

## CLI Commands

```bash
moon new <name>      # New project
moon build           # Build
moon run <target>    # Run
moon test            # Test
moon check           # Type check
moon fmt             # Format
moon add <pkg>       # Add package
moon doc             # Generate docs
```

## Type Annotations

```moonbit
let x : Int = 42
let f : (Int, Int) -> Int = add
let arr : Array[String] = []
let opt : Int? = Some(5)
let res : Result[Int, String] = Ok(42)
let func : (Int) -> Int = fn(x) { x * 2 }
```

## Common Functions

```moonbit
// String
s.length()
s.trim()
s.to_upper()
s.to_lower()
s.contains(substr)
s.split(separator)

// Array
arr.length()
arr.push(item)
arr.pop()
arr.map(fn)
arr.filter(fn)
arr.fold(init, fn)
arr.iter()

// Option
opt.is_some()
opt.is_none()
opt.unwrap()
opt.or(default)
opt.and_then(fn)

// Result
res.is_ok()
res.is_error()
res.unwrap()
res.or(default)
```

## Common Idioms

### Safe Division

```moonbit
fn safe_divide(a : Double, b : Double) -> Double? {
  if b == 0.0 { None } else { Some(a / b) }
}
```

### Parse with Error Handling

```moonbit
fn parse_int(s : String) -> Int!String {
  // parsing logic
  if valid { value } else { raise "Invalid" }
}
```

### Validate and Process

```moonbit
fn process_user(data : UserData) -> User!ValidationError {
  validate!(data)
  create_user(data)
}
```

---

## Quick Tips

1. **Use `|>` for pipelines** instead of nested function calls
2. **Prefer immutable** data structures when possible
3. **Use pattern matching** instead of multiple if-else
4. **Leverage type inference** but add annotations for clarity
5. **Use `defer`** for cleanup operations
6. **Chain iterators** instead of creating intermediate arrays
7. **Use checked exceptions** for better error tracking
8. **Write tests** with descriptive names

---

**For more details, see:**
- [snippets-moonbit.md](./snippets-moonbit.md) - Code examples
- [moonbit-course.md](./moonbit-course.md) - Complete course
- [getting-started.md](./getting-started.md) - Beginner guide
