# Complete MoonBit Programming Course

## Table of Contents

1. [Introduction to MoonBit](#introduction-to-moonbit)
2. [Getting Started](#getting-started)
3. [Basic Syntax](#basic-syntax)
4. [Variables and Data Types](#variables-and-data-types)
5. [Control Flow](#control-flow)
6. [Functions](#functions)
7. [Structs and Enums](#structs-and-enums)
8. [Pattern Matching](#pattern-matching)
9. [Error Handling](#error-handling)
10. [Collections](#collections)
11. [Iterators and Functional Programming](#iterators-and-functional-programming)
12. [Modules and Packages](#modules-and-packages)
13. [Traits](#traits)
14. [Generics](#generics)
15. [Advanced Features](#advanced-features)
16. [Async Programming](#async-programming)
17. [Testing](#testing)
18. [Best Practices](#best-practices)

---

## Introduction to MoonBit

MoonBit is a next-generation programming language designed for cloud and edge computing with WebAssembly as a primary compilation target. It combines functional, imperative, and object-oriented programming paradigms in a modern, efficient language.

### Key Features

- **WebAssembly First**: Native compilation to WebAssembly with excellent performance
- **Multi-Backend**: Compiles to WebAssembly, JavaScript, and native code
- **AI-Native Design**: Minimizes LLM hallucinations with consistent, predictable syntax
- **Built-in Async**: No explicit async/await keywords needed
- **Checked Exceptions**: Compile-time verification of error handling
- **Strong Type System**: Static typing with powerful type inference
- **Fast Compilation**: Extremely fast build times
- **Small Output**: Generates compact WebAssembly modules

### Use Cases

- WebAssembly applications
- Cloud and edge computing
- High-performance web applications
- Systems programming
- Data processing
- Blockchain and smart contracts

---

## Getting Started

### Installation

Install MoonBit using the official installer:

```bash
# On Unix-like systems (macOS, Linux)
curl -fsSL https://cli.moonbitlang.com/install/unix.sh | bash

# On Windows
powershell -Command "iwr https://cli.moonbitlang.com/install/windows.ps1 -useb | iex"
```

Verify installation:

```bash
moon version
```

### Creating Your First Project

```bash
# Create a new project
moon new hello_world
cd hello_world

# Build the project
moon build

# Run the project
moon run main
```

### Project Structure

```
hello_world/
├── moon.mod.json       # Package manifest
├── src/
│   └── main.mbt        # Main source file
└── target/             # Build output
```

### Your First Program

Create `src/main.mbt`:

```moonbit
fn main {
  println("Hello, MoonBit!")
}
```

Build and run:

```bash
moon build
moon run main
```

---

## Basic Syntax

### Comments

```moonbit
// Single line comment

/*
Multi-line
comment
*/

/// Documentation comment for the next item
fn documented_function() -> Unit {
  ()
}
```

### Expressions vs Statements

MoonBit is expression-based. Most constructs are expressions that return values.

```moonbit
fn main {
  // If is an expression
  let value = if true { 1 } else { 2 }

  // Match is an expression
  let result = match value {
    1 => "one"
    _ => "other"
  }

  // Blocks are expressions
  let x = {
    let a = 1
    let b = 2
    a + b  // Last expression is the value
  }

  println("x = \{x}")  // x = 3
}
```

### Indentation and Style

MoonBit uses braces for blocks and doesn't require semicolons at the end of statements.

```moonbit
fn main {
  let x = 1
  let y = 2

  if x < y {
    println("x is less than y")
  }

  for i = 0; i < 5; i = i + 1 {
    println(i)
  }
}
```

### String Interpolation

```moonbit
fn main {
  let name = "Alice"
  let age = 30

  println("My name is \{name} and I'm \{age} years old")

  // Expressions in interpolation
  println("Next year I'll be \{age + 1}")
}
```

---

## Variables and Data Types

### Variables

```moonbit
fn main {
  // Immutable by default
  let x = 5
  // x = 6  // Error: cannot assign to immutable variable

  // Mutable variables
  let mut y = 5
  y = 6  // OK

  // Type annotations
  let z : Int = 10
  let name : String = "Alice"
}
```

### Primitive Types

```moonbit
fn main {
  // Unit type (similar to void)
  let unit : Unit = ()

  // Boolean
  let is_valid : Bool = true
  let is_false : Bool = false

  // Integers (64-bit signed)
  let age : Int = 25
  let negative : Int = -100

  // Floating point (64-bit)
  let pi : Double = 3.14159
  let price : Double = 19.99

  // Character (Unicode)
  let letter : Char = 'A'
  let emoji : Char = '😊'

  // String
  let greeting : String = "Hello"
  let multiline : String = """
    This is a
    multiline string
  """

  // Bytes
  let data : Bytes = b"binary data"
}
```

### Compound Types

```moonbit
fn main {
  // Tuples
  let point : (Int, Int) = (3, 4)
  let person : (String, Int, String) = ("Alice", 30, "NYC")

  // Destructuring
  let (x, y) = point
  let (name, age, city) = person

  // Access by index
  let first = point.0
  let second = point.1

  println("Point: (\{x}, \{y})")
}
```

### Type Inference

MoonBit has powerful type inference:

```moonbit
fn main {
  let x = 42              // Inferred as Int
  let y = 3.14            // Inferred as Double
  let name = "Alice"      // Inferred as String
  let values = [1, 2, 3]  // Inferred as Array[Int]

  // Sometimes you need explicit types
  let empty : Array[Int] = []
}
```

### Overloaded Literals

MoonBit supports overloaded literals that adapt based on context:

```moonbit
fn main {
  // Same literal, different types
  let x : Int = 42
  let y : Int64 = 42
  let z : BigInt = 42

  // Array literals
  let arr1 : Array[Int] = [1, 2, 3]
  let arr2 : List[Int] = [1, 2, 3]

  // String-like literals
  let s1 : String = "hello"
  let s2 : StringBuilder = "hello"
}
```

---

## Control Flow

### If Expressions

```moonbit
fn main {
  let age = 18

  // Basic if
  if age >= 18 {
    println("Adult")
  }

  // If-else
  if age >= 18 {
    println("Adult")
  } else {
    println("Minor")
  }

  // If as expression
  let status = if age >= 18 { "adult" } else { "minor" }

  // Chained if-else
  let score = 85
  let grade = if score >= 90 {
    "A"
  } else if score >= 80 {
    "B"
  } else if score >= 70 {
    "C"
  } else {
    "D"
  }

  println("Status: \{status}, Grade: \{grade}")
}
```

### Match Expressions

```moonbit
fn main {
  let number = 3

  // Basic match
  let result = match number {
    1 => "one"
    2 => "two"
    3 => "three"
    _ => "other"
  }

  // Multiple patterns
  match number {
    1 | 2 | 3 => println("Between 1 and 3")
    _ => println("Other")
  }

  // Range patterns
  match number {
    1..=10 => println("Between 1 and 10")
    11..=20 => println("Between 11 and 20")
    _ => println("Out of range")
  }

  // Matching with extraction
  let point = (3, 4)
  match point {
    (0, 0) => println("Origin")
    (x, 0) => println("On x-axis at \{x}")
    (0, y) => println("On y-axis at \{y}")
    (x, y) => println("At (\{x}, \{y})")
  }
}
```

### Guard Expressions

```moonbit
fn classify(value : Int) -> String {
  match value {
    x if x < 0 => "negative"
    x if x == 0 => "zero"
    x if x % 2 == 0 => "positive even"
    x if x % 2 == 1 => "positive odd"
    _ => "unknown"
  }
}

fn main {
  println(classify(-5))  // "negative"
  println(classify(0))   // "zero"
  println(classify(6))   // "positive even"
  println(classify(7))   // "positive odd"
}
```

### Loops

#### Loop (Infinite)

```moonbit
fn main {
  let mut counter = 0

  loop {
    if counter >= 5 {
      break
    }
    println("Counter: \{counter}")
    counter = counter + 1
  }
}
```

#### While Loop

```moonbit
fn main {
  let mut n = 5

  while n > 0 {
    println(n)
    n = n - 1
  }

  println("Blast off!")
}
```

#### For Loop

```moonbit
fn main {
  // C-style for loop
  for i = 0; i < 5; i = i + 1 {
    println("Iteration: \{i}")
  }

  // For-in loop
  let numbers = [1, 2, 3, 4, 5]
  for num in numbers {
    println("Number: \{num}")
  }
}
```

#### Loop Control

```moonbit
fn main {
  // Break
  for i = 0; i < 10; i = i + 1 {
    if i == 5 {
      break
    }
    println(i)
  }

  // Continue
  for i = 0; i < 10; i = i + 1 {
    if i % 2 == 0 {
      continue
    }
    println(i)  // Only prints odd numbers
  }
}
```

---

## Functions

### Basic Functions

```moonbit
// Function with no parameters
fn greet() -> Unit {
  println("Hello!")
}

// Function with parameters
fn greet_person(name : String) -> Unit {
  println("Hello, \{name}!")
}

// Function with return value
fn add(a : Int, b : Int) -> Int {
  a + b  // Last expression is returned
}

// Explicit return
fn early_return(x : Int) -> Int {
  if x < 0 {
    return 0
  }
  x * 2
}

fn main {
  greet()
  greet_person("Alice")

  let sum = add(5, 3)
  println("Sum: \{sum}")

  println("Early return: \{early_return(-5)}")
}
```

### Labelled Arguments

```moonbit
fn create_user(~name : String, ~age : Int, ~email : String = "unknown") -> String {
  "User: \{name}, \{age}, \{email}"
}

fn main {
  // All arguments
  let user1 = create_user(name="Alice", age=30, email="alice@example.com")

  // Using default
  let user2 = create_user(name="Bob", age=25)

  // Order doesn't matter with labels
  let user3 = create_user(age=35, name="Charlie")

  println(user1)
  println(user2)
  println(user3)
}
```

### Optional Arguments

```moonbit
fn log_message(message : String, ~level : String = "INFO", ~timestamp : Bool = true) -> Unit {
  let prefix = if timestamp {
    "[2024-01-01] "
  } else {
    ""
  }
  println("\{prefix}[\{level}] \{message}")
}

fn main {
  log_message("Application started")
  log_message("Warning", level="WARN")
  log_message("Debug info", level="DEBUG", timestamp=false)
}
```

### Higher-Order Functions

```moonbit
fn apply_twice(f : (Int) -> Int, x : Int) -> Int {
  f(f(x))
}

fn double(x : Int) -> Int {
  x * 2
}

fn main {
  let result = apply_twice(double, 5)
  println("Result: \{result}")  // 20

  // With anonymous function
  let result2 = apply_twice(fn(x) { x + 1 }, 5)
  println("Result2: \{result2}")  // 7
}
```

### Closures

```moonbit
fn main {
  let multiplier = 10

  // Closure captures multiplier
  let multiply_by_ten = fn(x : Int) -> Int {
    x * multiplier
  }

  println(multiply_by_ten(5))  // 50

  // Closures in array operations
  let numbers = [1, 2, 3, 4, 5]
  let doubled = numbers.map(fn(x) { x * 2 })
  println("Doubled: \{doubled}")
}
```

### Partial Application

```moonbit
fn add(a : Int, b : Int) -> Int {
  a + b
}

fn multiply(a : Int, b : Int, c : Int) -> Int {
  a * b * c
}

fn main {
  // Using _ for partial application
  let add_five = add(5, _)
  println("5 + 3 = \{add_five(3)}")

  let double = multiply(2, _, 1)
  println("Double 7: \{double(7)}")

  // Multiple placeholders
  let multiply_by_6 = multiply(_, 2, 3)
  println("4 * 2 * 3 = \{multiply_by_6(4)}")
}
```

---

## Structs and Enums

### Structs

```moonbit
// Basic struct
struct Person {
  name : String
  age : Int
}

// Struct with methods
fn Person::new(name : String, age : Int) -> Person {
  { name, age }
}

fn greet(self : Person) -> String {
  "Hello, I'm \{self.name} and I'm \{self.age} years old"
}

fn have_birthday(self : Person) -> Person {
  { ..self, age: self.age + 1 }
}

fn main {
  let person = Person::new("Alice", 30)
  println(person.greet())

  let older_person = person.have_birthday()
  println("After birthday: \{older_person.greet()}")
}
```

### Mutable Structs

```moonbit
struct Counter {
  mut value : Int
}

fn increment(self : Counter) -> Unit {
  self.value = self.value + 1
}

fn get_value(self : Counter) -> Int {
  self.value
}

fn main {
  let counter = { value: 0 }

  counter.increment()
  counter.increment()

  println("Counter value: \{counter.get_value()}")
}
```

### Tuple Structs

```moonbit
struct Point(Int, Int)
struct Color(Int, Int, Int)

fn main {
  let point = Point(3, 4)
  let color = Color(255, 128, 0)

  // Access fields by index
  println("Point: (\{point.0}, \{point.1})")
  println("Color: RGB(\{color.0}, \{color.1}, \{color.2})")
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

enum Option[T] {
  Some(T)
  None
}

enum Result[T, E] {
  Ok(T)
  Err(E)
}

fn main {
  let red = Red
  let custom = RGB(255, 128, 0)

  match custom {
    Red => println("Pure red")
    Green => println("Pure green")
    Blue => println("Pure blue")
    RGB(r, g, b) => println("Custom RGB: \{r}, \{g}, \{b}")
  }

  let value : Option[Int] = Some(42)
  match value {
    Some(x) => println("Got value: \{x}")
    None => println("No value")
  }
}
```

### Enum Methods

```moonbit
enum Message {
  Quit
  Move(Int, Int)
  Write(String)
  ChangeColor(Int, Int, Int)
}

fn process(self : Message) -> Unit {
  match self {
    Quit => println("Quitting")
    Move(x, y) => println("Moving to (\{x}, \{y})")
    Write(text) => println("Writing: \{text}")
    ChangeColor(r, g, b) => println("Changing color to RGB(\{r}, \{g}, \{b})")
  }
}

fn main {
  let msg1 = Move(10, 20)
  let msg2 = Write("Hello")

  msg1.process()
  msg2.process()
}
```

---

## Pattern Matching

### Basic Patterns

```moonbit
fn main {
  let x = 5

  // Literal patterns
  match x {
    0 => println("zero")
    1 => println("one")
    2 => println("two")
    _ => println("other")
  }

  // Multiple patterns (OR)
  match x {
    1 | 2 | 3 => println("one, two, or three")
    _ => println("other")
  }

  // Range patterns
  match x {
    0..=10 => println("between 0 and 10")
    11..=20 => println("between 11 and 20")
    _ => println("out of range")
  }
}
```

### Destructuring Patterns

```moonbit
fn main {
  // Tuple destructuring
  let point = (3, 4)
  match point {
    (0, 0) => println("origin")
    (x, 0) => println("on x-axis: \{x}")
    (0, y) => println("on y-axis: \{y}")
    (x, y) => println("point: (\{x}, \{y})")
  }

  // Struct destructuring
  struct Person { name : String; age : Int }
  let person = { name: "Alice", age: 30 }

  match person {
    { name: "Alice", age } => println("It's Alice, age \{age}")
    { name, age: 30 } => println("\{name} is 30 years old")
    { name, age } => println("\{name} is \{age} years old")
  }
}
```

### Array Patterns

```moonbit
fn main {
  let arr = [1, 2, 3]

  match arr {
    [] => println("empty")
    [x] => println("single element: \{x}")
    [x, y] => println("two elements: \{x}, \{y}")
    [x, y, z] => println("three elements: \{x}, \{y}, \{z}")
    _ => println("more than three elements")
  }

  // With spread
  match arr {
    [first, .. ] => println("first: \{first}")
    _ => println("empty")
  }

  match arr {
    [first, .., last] => println("first: \{first}, last: \{last}")
    _ => println("empty or single element")
  }
}
```

### Map Patterns

```moonbit
fn main {
  let map = { "name": "Alice", "age": 30 }

  match map {
    { "name": String(name), "age": Number(age) } => {
      println("Person: \{name}, age: \{age}")
    }
    _ => println("invalid format")
  }
}
```

### JSON Patterns

```moonbit
fn main {
  let json = {
    "type": "user",
    "data": {
      "name": "Alice",
      "age": 30
    }
  }

  match json {
    {
      "type": String("user"),
      "data": {
        "name": String(name),
        "age": Number(age)
      }
    } => {
      println("User: \{name}, age: \{age}")
    }
    _ => println("unknown format")
  }
}
```

### Guard Patterns

```moonbit
fn classify(x : Int) -> String {
  match x {
    n if n < 0 => "negative"
    0 => "zero"
    n if n % 2 == 0 => "positive even"
    n if n % 2 == 1 => "positive odd"
    _ => "unknown"
  }
}

fn main {
  println(classify(-5))  // negative
  println(classify(0))   // zero
  println(classify(4))   // positive even
  println(classify(7))   // positive odd
}
```

---

## Error Handling

### The Error Type System

MoonBit has a unique checked error handling system. Functions that can raise errors have a `!` in their type signature.

```moonbit
// Function that can raise a String error
fn divide(a : Double, b : Double) -> Double!String {
  if b == 0.0 {
    raise "Division by zero"
  }
  a / b
}

// Function that can raise multiple error types
fn parse_and_divide(a : String, b : String) -> Double!ParseError {
  let num_a = parse!(a)
  let num_b = parse!(b)
  divide!(num_a, num_b)
}
```

### Try-Catch

```moonbit
fn main {
  // Basic try-catch
  try {
    let result = divide!(10.0, 0.0)
    println("Result: \{result}")
  } catch {
    error => println("Error: \{error}")
  }

  // Catching specific errors
  try {
    let result = risky_operation!()
    println("Success: \{result}")
  } catch {
    "file_not_found" => println("File not found")
    "permission_denied" => println("Permission denied")
    error => println("Other error: \{error}")
  }
}

fn risky_operation() -> String!String {
  raise "file_not_found"
}
```

### Error Propagation with !

```moonbit
// The ! suffix propagates errors up the call stack
fn read_and_process(filename : String) -> String!String {
  let content = read_file!(filename)  // Propagates error
  let processed = process_content!(content)  // Propagates error
  processed
}

fn read_file(filename : String) -> String!String {
  if filename == "" {
    raise "Empty filename"
  }
  "file content"
}

fn process_content(content : String) -> String!String {
  if content == "" {
    raise "Empty content"
  }
  "processed: \{content}"
}

fn main {
  try {
    let result = read_and_process!("data.txt")
    println(result)
  } catch {
    error => println("Error: \{error}")
  }
}
```

### Try? for Result Conversion

```moonbit
fn main {
  // Convert error to Result type
  let result : Result[Double, String] = divide(10.0, 0.0)

  match result {
    Ok(value) => println("Result: \{value}")
    Err(error) => println("Error: \{error}")
  }

  // Using map and default
  let safe_result = result.or(0.0)
  println("Safe result: \{safe_result}")
}
```

### Custom Error Types

```moonbit
enum FileError {
  NotFound(String)
  PermissionDenied(String)
  InvalidFormat(String)
}

fn read_config(path : String) -> String!FileError {
  if path == "" {
    raise NotFound("Empty path")
  }
  if not_accessible(path) {
    raise PermissionDenied(path)
  }
  // ... read file
  "config content"
}

fn not_accessible(path : String) -> Bool {
  false  // Simplified
}

fn main {
  try {
    let config = read_config!("config.json")
    println("Config: \{config}")
  } catch {
    NotFound(path) => println("File not found: \{path}")
    PermissionDenied(path) => println("Permission denied: \{path}")
    InvalidFormat(msg) => println("Invalid format: \{msg}")
  }
}
```

### Error Context

```moonbit
fn process_file(path : String) -> Unit!String {
  try {
    let content = read_file!(path)
    let data = parse_json!(content)
    save_data!(data)
  } catch {
    error => {
      // Add context to errors
      raise "Failed to process file \{path}: \{error}"
    }
  }
}

fn read_file(path : String) -> String!String {
  raise "file not found"
}

fn parse_json(content : String) -> Json!String {
  raise "invalid json"
}

fn save_data(data : Json) -> Unit!String {
  ()
}

fn main {
  try {
    process_file!("data.json")
  } catch {
    error => println("Error: \{error}")
  }
}
```

---

## Collections

### Arrays

```moonbit
fn main {
  // Creating arrays
  let arr1 : Array[Int] = []
  let arr2 = [1, 2, 3, 4, 5]
  let arr3 : Array[Int] = Array::make(5, 0)  // [0, 0, 0, 0, 0]

  // Accessing elements
  let first = arr2[0]
  let last = arr2[arr2.length() - 1]

  // Modifying arrays (mutable)
  let mut arr = [1, 2, 3]
  arr.push(4)
  arr[0] = 10

  println("Array: \{arr}")  // [10, 2, 3, 4]

  // Array operations
  let length = arr.length()
  let is_empty = arr.is_empty()
  let contains_3 = arr.contains(3)

  println("Length: \{length}, Empty: \{is_empty}, Contains 3: \{contains_3}")
}
```

### Array Methods

```moonbit
fn main {
  let numbers = [1, 2, 3, 4, 5]

  // Map
  let doubled = numbers.map(fn(x) { x * 2 })
  println("Doubled: \{doubled}")

  // Filter
  let evens = numbers.filter(fn(x) { x % 2 == 0 })
  println("Evens: \{evens}")

  // Fold
  let sum = numbers.fold(init=0, fn(acc, x) { acc + x })
  println("Sum: \{sum}")

  // Each
  numbers.each(fn(x) { println("Number: \{x}") })

  // Zip
  let letters = ["a", "b", "c"]
  let zipped = numbers.zip(letters)
  println("Zipped: \{zipped}")  // [(1, "a"), (2, "b"), (3, "c")]
}
```

### Lists

```moonbit
fn main {
  // Linked list
  let list1 : List[Int] = Nil
  let list2 = Cons(1, Cons(2, Cons(3, Nil)))

  // List operations
  let list3 = [1, 2, 3]  // Can also create with literal
  let head = list3.head()  // Some(1)
  let tail = list3.tail()  // [2, 3]

  println("List: \{list3}")
  println("Head: \{head}")
  println("Tail: \{tail}")
}
```

### Maps

```moonbit
fn main {
  // Creating maps
  let mut map : Map[String, Int] = Map::new()

  // Adding entries
  map.set("apple", 5)
  map.set("banana", 3)
  map.set("orange", 7)

  // Getting values
  let apple_count = map.get("apple")  // Some(5)
  let grape_count = map.get("grape")  // None

  match apple_count {
    Some(count) => println("Apples: \{count}")
    None => println("No apples")
  }

  // Checking existence
  if map.contains("banana") {
    println("We have bananas")
  }

  // Iterating
  map.iter().each(fn(entry) {
    let (key, value) = entry
    println("\{key}: \{value}")
  })

  // Removing
  map.remove("orange")

  // Size
  println("Map size: \{map.size()}")
}
```

### Sets

```moonbit
fn main {
  // Creating sets
  let mut set : Set[Int] = Set::new()

  // Adding elements
  set.insert(1)
  set.insert(2)
  set.insert(3)
  set.insert(2)  // Duplicates ignored

  println("Set: \{set}")  // {1, 2, 3}

  // Set operations
  let set2 = Set::from_array([2, 3, 4, 5])

  let union = set.union(set2)
  let intersection = set.intersect(set2)
  let difference = set.diff(set2)

  println("Union: \{union}")          // {1, 2, 3, 4, 5}
  println("Intersection: \{intersection}")  // {2, 3}
  println("Difference: \{difference}")      // {1}
}
```

---

## Iterators and Functional Programming

### Creating Iterators

```moonbit
fn main {
  let numbers = [1, 2, 3, 4, 5]

  // Get an iterator
  let iter = numbers.iter()

  // Iterate
  iter.each(fn(x) { println(x) })

  // Or use for-in
  for num in numbers {
    println(num)
  }
}
```

### Iterator Methods

```moonbit
fn main {
  let numbers = [1, 2, 3, 4, 5]

  // Map - transform each element
  let doubled = numbers.iter()
    .map(fn(x) { x * 2 })
    .collect()
  println("Doubled: \{doubled}")

  // Filter - keep only matching elements
  let evens = numbers.iter()
    .filter(fn(x) { x % 2 == 0 })
    .collect()
  println("Evens: \{evens}")

  // Fold - reduce to single value
  let sum = numbers.iter()
    .fold(init=0, fn(acc, x) { acc + x })
  println("Sum: \{sum}")

  // Take - take first n elements
  let first_three = numbers.iter()
    .take(3)
    .collect()
  println("First three: \{first_three}")

  // Drop - skip first n elements
  let after_two = numbers.iter()
    .drop(2)
    .collect()
  println("After two: \{after_two}")
}
```

### Chaining Operations

```moonbit
fn main {
  let numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

  // Chain multiple operations
  let result = numbers.iter()
    .filter(fn(x) { x % 2 == 0 })    // Keep evens
    .map(fn(x) { x * x })             // Square them
    .take(3)                          // Take first 3
    .collect()

  println("Result: \{result}")  // [4, 16, 36]

  // Sum of squares of even numbers
  let sum_of_squares = numbers.iter()
    .filter(fn(x) { x % 2 == 0 })
    .map(fn(x) { x * x })
    .fold(init=0, fn(acc, x) { acc + x })

  println("Sum of squares: \{sum_of_squares}")
}
```

### Pipe Operator

```moonbit
fn double(x : Int) -> Int { x * 2 }
fn add_ten(x : Int) -> Int { x + 10 }
fn square(x : Int) -> Int { x * x }

fn main {
  // Traditional function composition
  let result1 = square(add_ten(double(5)))
  println("Result1: \{result1}")

  // Using pipe operator
  let result2 = 5
    |> double
    |> add_ten
    |> square

  println("Result2: \{result2}")

  // With arrays
  let numbers = [1, 2, 3, 4, 5]
  let processed = numbers
    |> fn(arr) { arr.filter(fn(x) { x % 2 == 0 }) }
    |> fn(arr) { arr.map(fn(x) { x * x }) }

  println("Processed: \{processed}")
}
```

### Functional Programming Patterns

```moonbit
// Higher-order functions
fn apply_twice(f : (Int) -> Int, x : Int) -> Int {
  f(f(x))
}

fn compose[A, B, C](f : (B) -> C, g : (A) -> B) -> (A) -> C {
  fn(x) { f(g(x)) }
}

fn main {
  // Using apply_twice
  let double = fn(x : Int) -> Int { x * 2 }
  let result = apply_twice(double, 5)
  println("Apply twice: \{result}")  // 20

  // Using compose
  let add_one = fn(x : Int) -> Int { x + 1 }
  let add_one_then_double = compose(double, add_one)
  println("Composed: \{add_one_then_double(5)}")  // 12

  // Currying with partial application
  fn add(a : Int, b : Int) -> Int { a + b }
  let add_five = add(5, _)
  println("Add five: \{add_five(3)}")  // 8
}
```

---

## Modules and Packages

### Package Structure

A MoonBit package is defined by a `moon.mod.json` file:

```json
{
  "name": "username/mypackage",
  "version": "0.1.0",
  "deps": {
    "username/otherpackage": "0.2.0"
  }
}
```

### Creating Modules

Each `.mbt` file is a module. Organize your code into multiple files:

```
mypackage/
├── moon.mod.json
├── lib/
│   ├── math.mbt
│   ├── string.mbt
│   └── utils.mbt
└── main/
    └── main.mbt
```

### Visibility

```moonbit
// math.mbt

// Private function (default)
fn internal_helper(x : Int) -> Int {
  x * 2
}

// Public function
pub fn add(a : Int, b : Int) -> Int {
  a + b
}

// Public for all packages
pub(all) fn multiply(a : Int, b : Int) -> Int {
  a * b
}

// Public struct
pub struct Point {
  pub x : Int
  pub y : Int
}
```

### Importing

```moonbit
// main.mbt

// Import from same package
use math::add
use string::capitalize

// Import with alias
use math::multiply as mult

// Import multiple items
use utils::{format_date, format_time}

fn main {
  let sum = add(5, 3)
  let product = mult(4, 5)

  println("Sum: \{sum}, Product: \{product}")
}
```

### Package Dependencies

Install packages from mooncakes.io:

```bash
# Add dependency
moon add username/package

# Update dependencies
moon update

# Install dependencies
moon install
```

### Re-exports

```moonbit
// lib.mbt

// Re-export from submodules
pub use math::*
pub use string::*

// Selective re-export
pub use utils::{format_date, parse_date}
```

---

## Traits

### Defining Traits

```moonbit
trait Show {
  show(Self) -> String
}

trait Eq {
  op_equal(Self, Self) -> Bool
}

trait Compare : Eq {
  op_less(Self, Self) -> Bool
  op_less_equal(Self, Self) -> Bool
  op_greater(Self, Self) -> Bool
  op_greater_equal(Self, Self) -> Bool
}
```

### Implementing Traits

```moonbit
struct Point {
  x : Int
  y : Int
}

fn show(self : Point) -> String {
  "(\{self.x}, \{self.y})"
}

fn op_equal(self : Point, other : Point) -> Bool {
  self.x == other.x && self.y == other.y
}

fn main {
  let p1 = { x: 3, y: 4 }
  let p2 = { x: 3, y: 4 }

  println(p1.show())
  println("Equal: \{p1 == p2}")
}
```

### Generic Traits

```moonbit
trait Container[T] {
  add(Self, T) -> Unit
  get(Self, Int) -> T?
  size(Self) -> Int
}

struct Box[T] {
  mut items : Array[T]
}

fn Box::new[T]() -> Box[T] {
  { items: [] }
}

fn add[T](self : Box[T], item : T) -> Unit {
  self.items.push(item)
}

fn get[T](self : Box[T], index : Int) -> T? {
  if index >= 0 && index < self.items.length() {
    Some(self.items[index])
  } else {
    None
  }
}

fn size[T](self : Box[T]) -> Int {
  self.items.length()
}

fn main {
  let box : Box[Int] = Box::new()
  box.add(1)
  box.add(2)
  box.add(3)

  println("Size: \{box.size()}")
  match box.get(1) {
    Some(value) => println("Value: \{value}")
    None => println("Not found")
  }
}
```

### Trait Bounds

```moonbit
fn find_max[T : Compare](arr : Array[T]) -> T {
  let mut max = arr[0]
  for item in arr {
    if item > max {
      max = item
    }
  }
  max
}

fn print_all[T : Show](items : Array[T]) -> Unit {
  for item in items {
    println(item.show())
  }
}

fn main {
  let numbers = [3, 7, 2, 9, 1]
  let max = find_max(numbers)
  println("Max: \{max}")
}
```

### Derived Traits

```moonbit
// Automatically derive common traits
struct Point {
  x : Int
  y : Int
} derive(Show, Eq, Compare)

fn main {
  let p1 = { x: 3, y: 4 }
  let p2 = { x: 5, y: 6 }

  println(p1.show())
  println("Equal: \{p1 == p2}")
  println("Less than: \{p1 < p2}")
}
```

---

## Generics

### Generic Functions

```moonbit
fn identity[T](x : T) -> T {
  x
}

fn first[T](arr : Array[T]) -> T? {
  if arr.length() > 0 {
    Some(arr[0])
  } else {
    None
  }
}

fn swap[T](tuple : (T, T)) -> (T, T) {
  let (a, b) = tuple
  (b, a)
}

fn main {
  println(identity(42))
  println(identity("hello"))

  let numbers = [1, 2, 3]
  match first(numbers) {
    Some(x) => println("First: \{x}")
    None => println("Empty array")
  }

  let swapped = swap((1, 2))
  println("Swapped: \{swapped}")
}
```

### Generic Structs

```moonbit
struct Pair[T, U] {
  first : T
  second : U
}

fn Pair::new[T, U](first : T, second : U) -> Pair[T, U] {
  { first, second }
}

fn swap[T, U](self : Pair[T, U]) -> Pair[U, T] {
  { first: self.second, second: self.first }
}

fn main {
  let pair1 = Pair::new(1, "hello")
  let pair2 = pair1.swap()

  println("Original: (\{pair1.first}, \{pair1.second})")
  println("Swapped: (\{pair2.first}, \{pair2.second})")
}
```

### Generic Enums

```moonbit
enum Option[T] {
  Some(T)
  None
}

enum Result[T, E] {
  Ok(T)
  Err(E)
}

enum Tree[T] {
  Leaf
  Node(T, Tree[T], Tree[T])
}

fn Tree::insert[T : Compare](self : Tree[T], value : T) -> Tree[T] {
  match self {
    Leaf => Node(value, Leaf, Leaf)
    Node(v, left, right) => {
      if value < v {
        Node(v, left.insert(value), right)
      } else if value > v {
        Node(v, left, right.insert(value))
      } else {
        self
      }
    }
  }
}

fn main {
  let tree : Tree[Int] = Leaf
  let tree = tree.insert(5)
  let tree = tree.insert(3)
  let tree = tree.insert(7)

  println("Tree created")
}
```

### Constrained Generics

```moonbit
fn sort[T : Compare](arr : Array[T]) -> Array[T] {
  // Bubble sort implementation
  let mut result = arr.copy()
  for i = 0; i < result.length(); i = i + 1 {
    for j = 0; j < result.length() - i - 1; j = j + 1 {
      if result[j] > result[j + 1] {
        let temp = result[j]
        result[j] = result[j + 1]
        result[j + 1] = temp
      }
    }
  }
  result
}

fn print_if_positive[T : Compare + Show](value : T, zero : T) -> Unit {
  if value > zero {
    println(value.show())
  }
}

fn main {
  let numbers = [3, 1, 4, 1, 5, 9, 2, 6]
  let sorted = sort(numbers)
  println("Sorted: \{sorted}")

  print_if_positive(5, 0)
  print_if_positive(-3, 0)
}
```

---

## Advanced Features

### Newtype Pattern

```moonbit
struct UserId(Int)
struct Email(String)

fn UserId::new(id : Int) -> UserId {
  UserId(id)
}

fn Email::new(email : String) -> Email!String {
  if email.contains("@") {
    Email(email)
  } else {
    raise "Invalid email"
  }
}

fn send_email(user_id : UserId, email : Email) -> Unit {
  println("Sending email to user \{user_id.0} at \{email.0}")
}

fn main {
  let user_id = UserId::new(123)

  try {
    let email = Email::new!("user@example.com")
    send_email(user_id, email)
  } catch {
    error => println("Error: \{error}")
  }
}
```

### Builder Pattern

```moonbit
struct Config {
  mut host : String
  mut port : Int
  mut timeout : Int
  mut retries : Int
}

fn Config::new() -> Config {
  {
    host: "localhost",
    port: 8080,
    timeout: 30,
    retries: 3
  }
}

fn with_host(self : Config, host : String) -> Config {
  self.host = host
  self
}

fn with_port(self : Config, port : Int) -> Config {
  self.port = port
  self
}

fn with_timeout(self : Config, timeout : Int) -> Config {
  self.timeout = timeout
  self
}

fn build(self : Config) -> Config {
  self
}

fn main {
  let config = Config::new()
    ..with_host("example.com")
    ..with_port(9000)
    ..with_timeout(60)

  println("Config: \{config.host}:\{config.port}")
}
```

### Phantom Types

```moonbit
struct Validated
struct Unvalidated

struct User[State] {
  name : String
  email : String
}

fn User::new(name : String, email : String) -> User[Unvalidated] {
  { name, email }
}

fn validate(self : User[Unvalidated]) -> User[Validated]!String {
  if self.email.contains("@") {
    { name: self.name, email: self.email }
  } else {
    raise "Invalid email"
  }
}

fn save(self : User[Validated]) -> Unit {
  println("Saving validated user: \{self.name}")
}

fn main {
  let user = User::new("Alice", "alice@example.com")

  // Cannot save unvalidated user
  // user.save()  // Compile error!

  try {
    let validated_user = user.validate!()
    validated_user.save()
  } catch {
    error => println("Validation error: \{error}")
  }
}
```

### Defer

```moonbit
fn process_file(filename : String) -> Unit!String {
  println("Opening file: \{filename}")

  defer {
    println("Closing file: \{filename}")
  }

  if filename == "" {
    raise "Empty filename"
  }

  println("Processing file...")

  // Defer runs before function returns
}

fn main {
  try {
    process_file!("data.txt")
  } catch {
    error => println("Error: \{error}")
  }

  println("Done")
}
```

### Method Chaining with Cascade

```moonbit
struct StringBuilder {
  mut content : String
}

fn StringBuilder::new() -> StringBuilder {
  { content: "" }
}

fn append(self : StringBuilder, text : String) -> Unit {
  self.content = self.content + text
}

fn append_line(self : StringBuilder, text : String) -> Unit {
  self.content = self.content + text + "\n"
}

fn to_string(self : StringBuilder) -> String {
  self.content
}

fn main {
  let builder = StringBuilder::new()

  // Cascade operator for method chaining
  builder
    ..append("Hello")
    ..append(", ")
    ..append("World")
    ..append_line("!")
    ..append_line("How are you?")

  println(builder.to_string())
}
```

---

## Async Programming

### Built-in Async

MoonBit has built-in async support. Async functions look like regular functions:

```moonbit
// This function is async (IDE shows it in italics)
fn fetch_data(url : String) -> String {
  // Simulated async operation
  "Data from \{url}"
}

fn process_data(data : String) -> String {
  "Processed: \{data}"
}

fn main {
  // Async functions are called normally
  let data = fetch_data("https://api.example.com")
  let result = process_data(data)

  println(result)
}
```

### Concurrent Operations

```moonbit
fn fetch_user(id : Int) -> String {
  // Async operation
  "User \{id}"
}

fn fetch_posts(user_id : Int) -> Array[String] {
  // Async operation
  ["Post 1", "Post 2", "Post 3"]
}

fn main {
  // These run concurrently
  let users = [1, 2, 3, 4, 5].map(fetch_user)

  println("Users: \{users}")

  // Sequential chaining
  let user = fetch_user(1)
  let posts = fetch_posts(1)

  println("User \{user} has posts: \{posts}")
}
```

### Error Handling in Async

```moonbit
fn fetch_data(url : String) -> String!String {
  if url == "" {
    raise "Empty URL"
  }
  // Async fetch
  "Data from \{url}"
}

fn main {
  try {
    let data = fetch_data!("https://api.example.com")
    println(data)
  } catch {
    error => println("Error: \{error}")
  }
}
```

---

## Testing

### Unit Tests

```moonbit
fn add(a : Int, b : Int) -> Int {
  a + b
}

fn divide(a : Double, b : Double) -> Double!String {
  if b == 0.0 {
    raise "Division by zero"
  }
  a / b
}

test "add function" {
  assert_eq!(add(2, 3), 5)
  assert_eq!(add(-1, 1), 0)
  assert_eq!(add(0, 0), 0)
}

test "divide function" {
  try {
    assert_eq!(divide!(10.0, 2.0), 5.0)
    assert_eq!(divide!(9.0, 3.0), 3.0)
  } catch {
    _ => fail!("Should not raise error")
  }
}

test "divide by zero" {
  let result : Result[Double, String] = divide(10.0, 0.0)
  match result {
    Ok(_) => fail!("Should have raised error")
    Err(msg) => assert_eq!(msg, "Division by zero")
  }
}
```

### Test Assertions

```moonbit
test "assertions" {
  // Equality
  assert_eq!(1 + 1, 2)
  assert_ne!(1, 2)

  // Boolean
  assert!(true)
  assert!(1 < 2)

  // Failure
  // fail!("This test should fail")
}

test "array testing" {
  let arr = [1, 2, 3]

  assert_eq!(arr.length(), 3)
  assert_eq!(arr[0], 1)
  assert!(arr.contains(2))
}
```

### Property-Based Testing

```moonbit
fn reverse[T](arr : Array[T]) -> Array[T] {
  let mut result = []
  for i = arr.length() - 1; i >= 0; i = i - 1 {
    result.push(arr[i])
  }
  result
}

test "reverse twice is identity" {
  let arr = [1, 2, 3, 4, 5]
  let reversed_twice = reverse(reverse(arr))

  assert_eq!(arr, reversed_twice)
}

test "reverse preserves length" {
  let arr = [1, 2, 3, 4, 5]
  let reversed = reverse(arr)

  assert_eq!(arr.length(), reversed.length())
}
```

### Running Tests

```bash
# Run all tests
moon test

# Run specific test
moon test --file src/math_test.mbt

# Run with verbose output
moon test --verbose
```

---

## Best Practices

### Code Organization

```moonbit
// Group related functions
// math.mbt
pub fn add(a : Int, b : Int) -> Int { a + b }
pub fn subtract(a : Int, b : Int) -> Int { a - b }
pub fn multiply(a : Int, b : Int) -> Int { a * b }
pub fn divide(a : Int, b : Int) -> Int!String {
  if b == 0 { raise "Division by zero" }
  a / b
}

// Use descriptive names
fn calculate_total_price(items : Array[Item]) -> Double {
  items.iter()
    .map(fn(item) { item.price })
    .fold(init=0.0, fn(acc, price) { acc + price })
}

// Keep functions small and focused
fn validate_email(email : String) -> Bool {
  email.contains("@") && email.contains(".")
}

fn validate_age(age : Int) -> Bool {
  age >= 0 && age <= 150
}

fn validate_user(user : User) -> Bool!String {
  if not(validate_email(user.email)) {
    raise "Invalid email"
  }
  if not(validate_age(user.age)) {
    raise "Invalid age"
  }
  true
}
```

### Error Handling Best Practices

```moonbit
// Use specific error types
enum UserError {
  InvalidEmail(String)
  InvalidAge(Int)
  UserNotFound(Int)
}

fn find_user(id : Int) -> User!UserError {
  // Implementation
  raise UserNotFound(id)
}

// Provide context in error messages
fn process_user(id : Int) -> Unit!String {
  try {
    let user = find_user!(id)
    // Process user
  } catch {
    error => raise "Failed to process user \{id}: \{error}"
  }
}

// Use Result for recoverable errors
fn parse_config(text : String) -> Result[Config, String] {
  // Implementation
  Err("Invalid config")
}
```

### Type Design

```moonbit
// Use newtypes for type safety
struct UserId(Int)
struct ProductId(Int)

// This prevents accidentally mixing up IDs
fn get_user(id : UserId) -> User { /* ... */ }
fn get_product(id : ProductId) -> Product { /* ... */ }

// Use enums for state machines
enum OrderState {
  Pending
  Processing
  Shipped
  Delivered
  Cancelled
}

struct Order {
  id : Int
  state : OrderState
}

// Use builder pattern for complex objects
struct Query {
  mut table : String
  mut where_clause : String
  mut limit : Int
}

fn Query::new(table : String) -> Query {
  { table, where_clause: "", limit: 100 }
}

fn where(self : Query, clause : String) -> Query {
  self.where_clause = clause
  self
}

fn limit(self : Query, n : Int) -> Query {
  self.limit = n
  self
}
```

### Performance Tips

```moonbit
// Use iterators instead of creating intermediate arrays
fn process_large_array(arr : Array[Int]) -> Array[Int] {
  // Good: single pass
  arr.iter()
    .filter(fn(x) { x > 0 })
    .map(fn(x) { x * 2 })
    .collect()

  // Avoid: multiple passes creating intermediate arrays
  // let filtered = arr.filter(fn(x) { x > 0 })
  // let mapped = filtered.map(fn(x) { x * 2 })
}

// Use early returns to avoid unnecessary work
fn find_first_match(arr : Array[Int], target : Int) -> Int? {
  for i = 0; i < arr.length(); i = i + 1 {
    if arr[i] == target {
      return Some(i)
    }
  }
  None
}

// Prefer immutable data structures when possible
fn add_item(items : Array[Int], item : Int) -> Array[Int] {
  let mut new_items = items.copy()
  new_items.push(item)
  new_items
}
```

### Documentation

```moonbit
/// Calculates the factorial of a number
///
/// # Arguments
/// * `n` - The number to calculate factorial for
///
/// # Returns
/// The factorial of n
///
/// # Examples
/// ```
/// assert_eq!(factorial(5), 120)
/// ```
pub fn factorial(n : Int) -> Int {
  if n <= 1 {
    1
  } else {
    n * factorial(n - 1)
  }
}

/// A user in the system
pub struct User {
  /// The user's unique identifier
  pub id : UserId

  /// The user's email address
  pub email : String

  /// The user's age in years
  pub age : Int
}
```

### Testing Best Practices

```moonbit
// Test one thing at a time
test "add returns correct sum" {
  assert_eq!(add(2, 3), 5)
}

test "add handles negative numbers" {
  assert_eq!(add(-2, -3), -5)
  assert_eq!(add(-2, 3), 1)
}

// Use descriptive test names
test "user validation fails with invalid email" {
  let user = { name: "Alice", email: "invalid", age: 30 }
  let result = validate_user(user)

  match result {
    Ok(_) => fail!("Should have failed")
    Err(_) => ()
  }
}

// Test edge cases
test "divide handles zero" {
  let result = divide(10.0, 0.0)
  assert!(result.is_error())
}

test "array handles empty input" {
  let arr : Array[Int] = []
  assert_eq!(arr.length(), 0)
  assert_eq!(first(arr), None)
}
```

---

## Resources and Next Steps

### Essential Tools

- **moon**: The MoonBit build tool and package manager
- **MoonBit IDE**: VS Code extension with language support
- **mooncakes.io**: Package registry

### Useful Commands

```bash
moon new project_name      # Create new project
moon build                 # Build project
moon run main             # Run project
moon test                 # Run tests
moon check                # Type check without building
moon fmt                  # Format code
moon doc                  # Generate documentation
moon publish              # Publish package
moon add package          # Add dependency
```

### Learning Resources

- [Official Documentation](https://docs.moonbitlang.com/)
- [MoonBit Language Tour](https://www.moonbitlang.com/tour/)
- [GitHub Repository](https://github.com/moonbitlang/moonbit-docs)
- [Package Registry](https://mooncakes.io/)

### Community

- [GitHub Discussions](https://github.com/moonbitlang/moonbit-docs/discussions)
- [Discord Server](https://discord.gg/moonbitlang)

### Popular Packages

- **moonbitlang/core**: Core library with essential data structures
- **moonbitlang/json**: JSON parsing and serialization
- **moonbitlang/http**: HTTP client and server
- **moonbitlang/test**: Testing utilities

### Next Steps

1. **Practice with small projects** - Build simple applications
2. **Explore the standard library** - Learn the core data structures
3. **Read existing code** - Study open-source MoonBit projects
4. **Contribute to packages** - Help improve the ecosystem
5. **Build WebAssembly apps** - Leverage MoonBit's primary target

### Key Differences from Other Languages

**From Rust:**
- No explicit lifetimes in most cases
- Built-in async without keywords
- Checked exceptions instead of Result everywhere
- Simpler syntax, less verbose

**From TypeScript:**
- Stronger type system
- No null/undefined issues (use Option)
- Compiles to WebAssembly, not just JavaScript
- Better performance

**From Go:**
- More powerful type system with generics
- Functional programming features
- Pattern matching
- Expression-based syntax

**From Python:**
- Static typing with inference
- Much faster execution
- Compile-time error checking
- Smaller deployment size

This course provides a comprehensive foundation in MoonBit programming. Practice these concepts by building projects and exploring the ecosystem!
