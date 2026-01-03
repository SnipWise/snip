## Hello World
Basic program structure and main function
```moonbit
fn main {
  println("Hello, World!")
}
```

----------

## Variable Declaration
Different ways to declare and initialize variables
```moonbit
fn main {
  let name = "John"           // Immutable (default)
  let mut age = 30            // Mutable
  let city : String = "New York"  // With type annotation

  age = age + 1  // Modify mutable variable

  println("Name: \{name}, Age: \{age}, City: \{city}")
}
```

----------

## Constants
Declaring constants and top-level values
```moonbit
const PI : Double = 3.14159
const MAX_USERS : Int = 100
let GREETING : String = "Hello"  // Top-level let

fn main {
  println("Pi: \{PI}, Max Users: \{MAX_USERS}, Greeting: \{GREETING}")
}
```

----------

## Data Types
Working with basic data types
```moonbit
fn main {
  let integer : Int = 42
  let float : Double = 3.14
  let boolean : Bool = true
  let character : Char = '🦀'
  let string : String = "MoonBit"
  let bytes : Bytes = b"binary data"
  let unit : Unit = ()

  println("int: \{integer}, float: \{float}, bool: \{boolean}, char: \{character}")
  println("String: \{string}")
}
```

----------

## Arrays and Lists
Working with arrays and lists
```moonbit
fn main {
  let arr : Array[Int] = [1, 2, 3]
  let mut vec = [1, 2, 3, 4]

  vec.push(5)
  let last = vec.pop()

  println("Array: \{arr}")
  println("Vector: \{vec}, len: \{vec.length()}")

  // Iteration
  for i = 0; i < vec.length(); i = i + 1 {
    println("Index \{i}: \{vec[i]}")
  }

  // Functional iteration
  vec.iter().each(fn(value) { println("Value: \{value}") })
}
```

----------

## Maps
Creating and manipulating hash maps
```moonbit
fn main {
  let mut map : Map[String, Int] = Map::new()
  map.set("apple", 5)
  map.set("banana", 3)

  // Check if key exists
  match map.get("apple") {
    Some(value) => println("Apple: \{value}")
    None => println("Apple not found")
  }

  // Iterate over map
  map.iter().each(fn(entry) {
    let (key, value) = entry
    println("\{key}: \{value}")
  })
}
```

----------

## Control Flow - If/Else
Conditional statements
```moonbit
fn main {
  let age = 25

  if age >= 18 {
    println("Adult")
  } else {
    println("Minor")
  }

  // If as expression
  let status = if age >= 18 { "adult" } else { "minor" }
  println("Status: \{status}")

  let score = 85
  let grade = if score >= 90 {
    "A"
  } else if score >= 80 {
    "B"
  } else {
    "C"
  }
  println("Grade: \{grade}")
}
```

----------

## Loops
Different types of loops
```moonbit
fn main {
  // Loop with break
  let mut counter = 0
  loop {
    if counter >= 3 {
      break
    }
    println("Loop: \{counter}")
    counter = counter + 1
  }

  // While loop
  let mut n = 3
  while n > 0 {
    println("While: \{n}")
    n = n - 1
  }

  // For loop
  for i = 0; i < 3; i = i + 1 {
    println("For: \{i}")
  }

  // For-in loop
  let numbers = [1, 2, 3]
  for num in numbers {
    println("For-in: \{num}")
  }
}
```

----------

## Match Statements
Pattern matching with match
```moonbit
fn main {
  let number = 3

  match number {
    1 => println("One")
    2 | 3 => println("Two or Three")
    4..=10 => println("Four to Ten")
    _ => println("Something else")
  }

  let option_value : Int? = Some(5)
  match option_value {
    Some(x) => println("Got value: \{x}")
    None => println("No value")
  }
}
```

----------

## Functions
Function declaration and parameters
```moonbit
fn add(a : Int, b : Int) -> Int {
  a + b  // Last expression is return value
}

fn divide(a : Double, b : Double) -> Double!String {
  if b == 0.0 {
    raise "Division by zero"
  }
  a / b
}

fn main {
  let sum = add(5, 3)
  println("Sum: \{sum}")

  try {
    let result = divide!(10.0, 2.0)
    println("Division: \{result}")
  } catch {
    error => println("Error: \{error}")
  }
}
```

----------

## Structs
Defining and using structs
```moonbit
struct Person {
  name : String
  age : Int
}

fn Person::new(name : String, age : Int) -> Person {
  { name, age }
}

fn greet(self : Person) -> String {
  "Hello, I'm \{self.name}"
}

fn main {
  let person = Person::new("Alice", 30)
  println("Person: \{person.name} (\{person.age})")
  println(person.greet())
}
```

----------

## Enums
Defining and using enums
```moonbit
enum Color {
  Red
  Green
  Blue
  RGB(Int, Int, Int)
}

fn main {
  let red = Red
  let custom = RGB(255, 128, 0)

  match red {
    Red => println("It's red!")
    Green => println("It's green!")
    Blue => println("It's blue!")
    RGB(r, g, b) => println("RGB: \{r}, \{g}, \{b}")
  }
}
```

----------

## Option and Result
Working with Option and Result types
```moonbit
fn find_user(id : Int) -> String? {
  if id == 1 {
    Some("Alice")
  } else {
    None
  }
}

fn parse_number(s : String) -> Int!String {
  // Simplified - would use proper parsing
  if s == "42" {
    42
  } else {
    raise "Parse error"
  }
}

fn main {
  match find_user(1) {
    Some(name) => println("Found user: \{name}")
    None => println("User not found")
  }

  try {
    let n = parse_number!("42")
    println("Parsed: \{n}")
  } catch {
    e => println("Parse error: \{e}")
  }
}
```

----------

## Error Handling
Proper error handling with checked exceptions
```moonbit
fn read_file_content(path : String) -> String!String {
  // Simplified - would use actual file I/O
  if path == "test.txt" {
    "File content"
  } else {
    raise "File not found"
  }
}

fn main {
  try {
    let content = read_file_content!("test.txt")
    println("File content: \{content}")
  } catch {
    e => println("Error reading file: \{e}")
  }

  // Using try? to convert to Result
  let result : Result[String, String] = read_file_content("test.txt")
  match result {
    Ok(content) => println("Got content: \{content}")
    Err(e) => println("Error: \{e}")
  }
}
```

----------

## Traits
Defining and implementing traits
```moonbit
trait Drawable {
  draw(Self) -> Unit
}

struct Circle {
  radius : Double
}

struct Rectangle {
  width : Double
  height : Double
}

fn draw(self : Circle) -> Unit {
  println("Drawing circle with radius \{self.radius}")
}

fn draw(self : Rectangle) -> Unit {
  println("Drawing rectangle \{self.width}x\{self.height}")
}

fn main {
  let circle = { radius: 5.0 }
  let rect = { width: 3.0, height: 4.0 }

  circle.draw()
  rect.draw()
}
```

----------

## Generics
Using generic types and functions
```moonbit
fn find_largest[T : Compare](list : Array[T]) -> T {
  let mut largest = list[0]
  for item in list {
    if item > largest {
      largest = item
    }
  }
  largest
}

struct Point[T] {
  x : T
  y : T
}

fn main {
  let numbers = [34, 50, 25, 100, 65]
  let largest = find_largest(numbers)
  println("Largest number: \{largest}")

  let point : Point[Int] = { x: 5, y: 10 }
  println("Point: (\{point.x}, \{point.y})")
}
```

----------

## Closures
Using closures and functional programming
```moonbit
fn main {
  let numbers = [1, 2, 3, 4, 5]

  // Closure that captures environment
  let multiplier = 2
  let doubled = numbers.map(fn(x) { x * multiplier })
  println("Doubled: \{doubled}")

  // Filter and collect
  let evens = numbers.filter(fn(x) { x % 2 == 0 })
  println("Evens: \{evens}")

  // Using closure as parameter
  let add_one = fn(x) { x + 1 }
  println("5 + 1 = \{add_one(5)}")
}
```

----------

## Iterators
Working with iterators
```moonbit
fn main {
  let vec = [1, 2, 3, 4, 5]

  // Iterator methods
  let sum = vec.iter().fold(init=0, fn(acc, x) { acc + x })
  println("Sum: \{sum}")

  // Chain operations
  let result = vec.iter()
    .filter(fn(x) { x > 2 })
    .map(fn(x) { x * 2 })
    .collect()
  println("Filtered and doubled: \{result}")

  // Enumerate
  vec.iter().eachi(fn(i, value) {
    println("Index \{i}: \{value}")
  })
}
```

----------

## Pipe Operator
Using the pipe operator for function chaining
```moonbit
fn double(x : Int) -> Int { x * 2 }
fn add_ten(x : Int) -> Int { x + 10 }
fn to_string(x : Int) -> String { x.to_string() }

fn main {
  let result = 5
    |> double
    |> add_ten
    |> to_string

  println("Result: \{result}")  // "Result: 20"

  // With inline functions
  let processed = [1, 2, 3]
    |> fn(arr) { arr.map(fn(x) { x * 2 }) }
    |> fn(arr) { arr.filter(fn(x) { x > 2 }) }

  println("Processed: \{processed}")
}
```

----------

## JSON Handling
Working with JSON data
```moonbit
fn main {
  // Create JSON
  let json = {
    "name": "Alice",
    "age": 30,
    "email": "alice@example.com"
  }

  // Access JSON fields
  match json["name"] {
    String(name) => println("Name: \{name}")
    _ => println("Name not found")
  }

  // Pattern matching on JSON
  match json {
    { "name": String(name), "age": Number(age), .. } => {
      println("Person: \{name}, age: \{age}")
    }
    _ => println("Invalid format")
  }

  // Convert to string
  let json_str = json.to_string()
  println("JSON: \{json_str}")
}
```

----------

## String Manipulation
Common string operations
```moonbit
fn main {
  let text = "  Hello, World!  "

  println("Original: '\{text}'")
  println("Trimmed: '\{text.trim()}'")
  println("Uppercase: \{text.to_upper()}")
  println("Lowercase: \{text.to_lower()}")

  // String interpolation
  let name = "Alice"
  let age = 30
  let message = "My name is \{name} and I'm \{age} years old"
  println(message)

  // String methods
  let sentence = "moonbit is awesome"
  println("Contains 'moonbit': \{sentence.contains("moonbit")}")
  println("Starts with 'moon': \{sentence.starts_with("moon")}")
  println("Length: \{sentence.length()}")

  // Split and join
  let words = sentence.split(" ")
  println("Words: \{words}")
}
```

----------

## Testing
Unit tests and test-driven development
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

test "add function works correctly" {
  assert_eq!(add(2, 3), 5)
  assert_eq!(add(-1, 1), 0)
}

test "divide function" {
  try {
    assert_eq!(divide!(10.0, 2.0), 5.0)
  } catch {
    _ => fail!("Should not raise error")
  }
}

test "divide by zero raises error" {
  let result : Result[Double, String] = divide(10.0, 0.0)
  match result {
    Ok(_) => fail!("Should have raised error")
    Err(_) => ()  // Expected
  }
}
```

----------

## Async Programming
Asynchronous operations (built-in async)
```moonbit
// Async functions don't need special keywords
// The compiler handles async automatically
fn fetch_data(id : Int) -> String {
  // Simulated async operation
  // In real code, this would be an async I/O operation
  "Data for ID: \{id}"
}

fn process_multiple() -> Unit {
  let tasks = [1, 2, 3]

  // Process concurrently
  let results = tasks.map(fetch_data)

  for result in results {
    println(result)
  }
}

fn main {
  println("Starting async operations...")
  process_multiple()
  println("All operations completed")
}
```

----------

## Defer
Cleanup with defer expressions
```moonbit
fn process_file(path : String) -> Unit!String {
  println("Opening file: \{path}")
  defer { println("Closing file: \{path}") }

  // File processing logic
  if path == "" {
    raise "Empty path"
  }

  println("Processing file...")
  // Defer will run even if we return early or raise an error
}

fn main {
  try {
    process_file!("test.txt")
  } catch {
    e => println("Error: \{e}")
  }
  // "Closing file" is printed before this point
}
```

----------

## Partial Application
Using underscore for partial application
```moonbit
fn add(a : Int, b : Int) -> Int {
  a + b
}

fn multiply(a : Int, b : Int, c : Int) -> Int {
  a * b * c
}

fn main {
  // Partial application with _
  let add_five = add(5, _)
  println("5 + 3 = \{add_five(3)}")

  let double = multiply(2, _, 1)
  println("Double 7: \{double(7)}")

  // Multiple placeholders
  let multiply_by_2_and_3 = multiply(_, 2, 3)
  println("4 * 2 * 3 = \{multiply_by_2_and_3(4)}")
}
```

----------

## Guard Expressions
Pattern matching with guards
```moonbit
fn classify_number(n : Int) -> String {
  match n {
    x if x < 0 => "negative"
    x if x == 0 => "zero"
    x if x < 10 => "single digit"
    x if x < 100 => "double digit"
    _ => "large number"
  }
}

fn main {
  println(classify_number(-5))   // "negative"
  println(classify_number(0))    // "zero"
  println(classify_number(7))    // "single digit"
  println(classify_number(42))   // "double digit"
  println(classify_number(150))  // "large number"
}
```

----------

## Labelled Arguments
Functions with labelled parameters
```moonbit
fn create_profile(~name : String, ~age : Int, ~city : String = "Unknown") -> String {
  "Name: \{name}, Age: \{age}, City: \{city}"
}

fn main {
  // All arguments
  let profile1 = create_profile(name="Alice", age=30, city="New York")
  println(profile1)

  // Using default
  let profile2 = create_profile(name="Bob", age=25)
  println(profile2)

  // Different order
  let profile3 = create_profile(age=35, name="Charlie", city="Boston")
  println(profile3)
}
```

----------

## Cascade Operator
Chaining mutable operations
```moonbit
struct Builder {
  mut name : String
  mut age : Int
  mut city : String
}

fn set_name(self : Builder, name : String) -> Unit {
  self.name = name
}

fn set_age(self : Builder, age : Int) -> Unit {
  self.age = age
}

fn set_city(self : Builder, city : String) -> Unit {
  self.city = city
}

fn main {
  let builder = { name: "", age: 0, city: "" }

  // Cascade operator for chaining
  builder
    ..set_name("Alice")
    ..set_age(30)
    ..set_city("New York")

  println("Built: \{builder.name}, \{builder.age}, \{builder.city}")
}
```
