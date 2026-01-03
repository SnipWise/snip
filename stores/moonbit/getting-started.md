# MoonBit Getting Started Guide

A quick-start guide to get you up and running with MoonBit in minutes.

## Installation

### macOS / Linux

```bash
curl -fsSL https://cli.moonbitlang.com/install/unix.sh | bash
```

### Windows

```powershell
powershell -Command "iwr https://cli.moonbitlang.com/install/windows.ps1 -useb | iex"
```

### Verify Installation

```bash
moon version
```

You should see output like:
```
moon 0.1.x
```

## IDE Setup

### VS Code (Recommended)

1. Install [Visual Studio Code](https://code.visualstudio.com/)
2. Install the MoonBit extension from the marketplace
3. The extension provides:
   - Syntax highlighting
   - Code completion
   - Error checking
   - Format on save
   - Async function indicators (italics)
   - Error-throwing function indicators (underlines)

## Your First Project

### Create a New Project

```bash
# Create a new project
moon new hello_world

# Navigate to the project
cd hello_world
```

This creates:
```
hello_world/
├── moon.mod.json    # Package manifest
└── src/
    └── main.mbt     # Main source file
```

### Hello World

Open `src/main.mbt` and you'll see:

```moonbit
fn main {
  println("Hello, MoonBit!")
}
```

### Build and Run

```bash
# Build the project
moon build

# Run the project
moon run main
```

Output:
```
Hello, MoonBit!
```

## Basic Syntax

### Variables

```moonbit
fn main {
  // Immutable (default)
  let name = "Alice"

  // Mutable
  let mut age = 30
  age = age + 1

  // With type annotation
  let city : String = "New York"

  println("Name: \{name}, Age: \{age}, City: \{city}")
}
```

### Data Types

```moonbit
fn main {
  let integer : Int = 42
  let float : Double = 3.14
  let boolean : Bool = true
  let character : Char = 'A'
  let text : String = "Hello"

  let array = [1, 2, 3, 4, 5]
  let tuple = (10, "hello", true)

  println("Integer: \{integer}")
  println("Array: \{array}")
}
```

### Functions

```moonbit
// Function with parameters and return type
fn add(a : Int, b : Int) -> Int {
  a + b  // Last expression is returned
}

// Function with no return value (Unit)
fn greet(name : String) -> Unit {
  println("Hello, \{name}!")
}

fn main {
  let sum = add(5, 3)
  println("Sum: \{sum}")

  greet("Alice")
}
```

### Control Flow

```moonbit
fn main {
  let age = 25

  // If-else
  if age >= 18 {
    println("Adult")
  } else {
    println("Minor")
  }

  // Match
  let grade = match age {
    0..=12 => "Child"
    13..=17 => "Teenager"
    18..=64 => "Adult"
    _ => "Senior"
  }

  println("Category: \{grade}")
}
```

### Loops

```moonbit
fn main {
  // For loop
  for i = 0; i < 5; i = i + 1 {
    println("Count: \{i}")
  }

  // For-in loop
  let numbers = [1, 2, 3, 4, 5]
  for num in numbers {
    println("Number: \{num}")
  }

  // While loop
  let mut count = 0
  while count < 3 {
    println("Count: \{count}")
    count = count + 1
  }
}
```

## Working with Collections

### Arrays

```moonbit
fn main {
  let mut numbers = [1, 2, 3]

  // Add element
  numbers.push(4)

  // Access element
  let first = numbers[0]

  // Length
  let size = numbers.length()

  println("Array: \{numbers}, First: \{first}, Size: \{size}")
}
```

### Maps

```moonbit
fn main {
  let mut scores : Map[String, Int] = Map::new()

  // Add entries
  scores.set("Alice", 95)
  scores.set("Bob", 87)

  // Get value
  match scores.get("Alice") {
    Some(score) => println("Alice's score: \{score}")
    None => println("Score not found")
  }

  // Iterate
  scores.iter().each(fn(entry) {
    let (name, score) = entry
    println("\{name}: \{score}")
  })
}
```

## Error Handling

### Basic Error Handling

```moonbit
fn divide(a : Double, b : Double) -> Double!String {
  if b == 0.0 {
    raise "Division by zero"
  }
  a / b
}

fn main {
  try {
    let result = divide!(10.0, 2.0)
    println("Result: \{result}")
  } catch {
    error => println("Error: \{error}")
  }
}
```

### Error Propagation

```moonbit
fn calculate(a : Double, b : Double) -> Double!String {
  let result1 = divide!(a, b)
  let result2 = divide!(result1, 2.0)
  result2
}

fn main {
  try {
    let result = calculate!(10.0, 2.0)
    println("Result: \{result}")
  } catch {
    error => println("Calculation error: \{error}")
  }
}
```

## Pattern Matching

```moonbit
enum Shape {
  Circle(Double)
  Rectangle(Double, Double)
  Triangle(Double, Double, Double)
}

fn area(shape : Shape) -> Double {
  match shape {
    Circle(radius) => 3.14159 * radius * radius
    Rectangle(width, height) => width * height
    Triangle(a, b, c) => {
      let s = (a + b + c) / 2.0
      (s * (s - a) * (s - b) * (s - c)).sqrt()
    }
  }
}

fn main {
  let circle = Circle(5.0)
  let rect = Rectangle(4.0, 6.0)

  println("Circle area: \{area(circle)}")
  println("Rectangle area: \{area(rect)}")
}
```

## Structs

```moonbit
struct Person {
  name : String
  age : Int
  email : String
}

fn Person::new(name : String, age : Int, email : String) -> Person {
  { name, age, email }
}

fn greet(self : Person) -> String {
  "Hello, I'm \{self.name} and I'm \{self.age} years old"
}

fn main {
  let person = Person::new("Alice", 30, "alice@example.com")
  println(person.greet())
  println("Email: \{person.email}")
}
```

## Functional Programming

### Map, Filter, Fold

```moonbit
fn main {
  let numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

  // Map: transform each element
  let doubled = numbers.map(fn(x) { x * 2 })

  // Filter: keep only matching elements
  let evens = numbers.filter(fn(x) { x % 2 == 0 })

  // Fold: reduce to single value
  let sum = numbers.fold(init=0, fn(acc, x) { acc + x })

  println("Doubled: \{doubled}")
  println("Evens: \{evens}")
  println("Sum: \{sum}")
}
```

### Pipe Operator

```moonbit
fn double(x : Int) -> Int { x * 2 }
fn add_ten(x : Int) -> Int { x + 10 }
fn to_string(x : Int) -> String { x.to_string() }

fn main {
  let result = 5
    |> double      // 10
    |> add_ten     // 20
    |> to_string   // "20"

  println("Result: \{result}")
}
```

## Testing

### Writing Tests

Create `src/math.mbt`:

```moonbit
pub fn add(a : Int, b : Int) -> Int {
  a + b
}

pub fn multiply(a : Int, b : Int) -> Int {
  a * b
}

test "add function" {
  assert_eq!(add(2, 3), 5)
  assert_eq!(add(-1, 1), 0)
}

test "multiply function" {
  assert_eq!(multiply(2, 3), 6)
  assert_eq!(multiply(-2, 3), -6)
}
```

### Running Tests

```bash
moon test
```

Output:
```
test math::add function ... ok
test math::multiply function ... ok

test result: ok. 2 passed; 0 failed
```

## Building a Simple CLI Tool

Create `src/main.mbt`:

```moonbit
fn main {
  let args = @args.get_args()

  if args.length() == 0 {
    println("Usage: calculator <operation> <a> <b>")
    println("Operations: add, sub, mul, div")
    return
  }

  let operation = args[0]
  let a = parse_number(args[1])
  let b = parse_number(args[2])

  let result = match operation {
    "add" => a + b
    "sub" => a - b
    "mul" => a * b
    "div" => {
      if b == 0.0 {
        println("Error: Division by zero")
        return
      }
      a / b
    }
    _ => {
      println("Unknown operation: \{operation}")
      return
    }
  }

  println("Result: \{result}")
}

fn parse_number(s : String) -> Double {
  // Simplified - use proper parsing in real code
  0.0
}
```

## Working with JSON

```moonbit
fn main {
  // Create JSON object
  let person = {
    "name": "Alice",
    "age": 30,
    "email": "alice@example.com",
    "hobbies": ["reading", "coding", "hiking"]
  }

  // Access fields
  match person["name"] {
    String(name) => println("Name: \{name}")
    _ => println("Name not found")
  }

  // Pattern matching
  match person {
    {
      "name": String(name),
      "age": Number(age),
      "email": String(email)
    } => {
      println("Person: \{name}, \{age}, \{email}")
    }
    _ => println("Invalid format")
  }

  // Convert to string
  let json_str = person.to_string()
  println("JSON: \{json_str}")
}
```

## Package Management

### Adding Dependencies

Edit `moon.mod.json`:

```json
{
  "name": "myapp",
  "version": "0.1.0",
  "deps": {
    "username/package-name": "0.2.0"
  }
}
```

Or use the CLI:

```bash
moon add username/package-name
```

### Using Dependencies

```moonbit
use package_name::function_name

fn main {
  let result = function_name()
  println(result)
}
```

## Common Commands Reference

```bash
# Project management
moon new <name>          # Create new project
moon build               # Build project
moon run <target>        # Run project
moon clean               # Clean build artifacts

# Testing
moon test                # Run all tests
moon test <file>         # Run specific test file

# Code quality
moon check               # Type check without building
moon fmt                 # Format code
moon doc                 # Generate documentation

# Package management
moon add <package>       # Add dependency
moon remove <package>    # Remove dependency
moon update              # Update dependencies
moon install             # Install dependencies

# Publishing
moon publish             # Publish to mooncakes.io
```

## Project Templates

### Web API Server

```moonbit
struct User {
  id : Int
  name : String
  email : String
}

let mut users : Map[Int, User] = Map::new()

fn get_user(id : Int) -> User? {
  users.get(id)
}

fn create_user(name : String, email : String) -> User {
  let id = users.size() + 1
  let user = { id, name, email }
  users.set(id, user)
  user
}

fn main {
  let user = create_user("Alice", "alice@example.com")
  println("Created user: \{user.name}")

  match get_user(user.id) {
    Some(u) => println("Found user: \{u.name}")
    None => println("User not found")
  }
}
```

### Data Processing Pipeline

```moonbit
fn process_data(data : Array[Int]) -> Array[Int] {
  data.iter()
    .filter(fn(x) { x > 0 })           // Remove negatives
    .map(fn(x) { x * 2 })               // Double values
    .filter(fn(x) { x < 100 })          // Keep values < 100
    .collect()
}

fn main {
  let raw_data = [5, -3, 10, 200, 25, -8, 45]
  let processed = process_data(raw_data)

  println("Raw data: \{raw_data}")
  println("Processed: \{processed}")
}
```

## Next Steps

1. **Explore the [snippets file](./snippets-moonbit.md)** for more code examples
2. **Read the [complete course](./moonbit-course.md)** for in-depth learning
3. **Visit [docs.moonbitlang.com](https://docs.moonbitlang.com/)** for official documentation
4. **Browse [mooncakes.io](https://mooncakes.io/)** for available packages
5. **Join the [Discord community](https://discord.gg/moonbitlang)** for help and discussions

## Troubleshooting

### Common Issues

**"moon: command not found"**
- Make sure you've added moon to your PATH
- Restart your terminal after installation

**Build errors**
- Run `moon clean` and try building again
- Check for syntax errors in your code
- Ensure all dependencies are installed

**Type errors**
- Enable type checking: `moon check`
- Add explicit type annotations where needed

**IDE not working**
- Restart VS Code
- Reinstall the MoonBit extension
- Check that `moon` is in your PATH

## Getting Help

- **Documentation**: [docs.moonbitlang.com](https://docs.moonbitlang.com/)
- **Discord**: Join the MoonBit community
- **GitHub**: [github.com/moonbitlang](https://github.com/moonbitlang)
- **Issues**: Report bugs on GitHub

---

**You're ready to start coding with MoonBit!** 🚀

Try modifying the examples above, then move on to [snippets-moonbit.md](./snippets-moonbit.md) for more patterns and [moonbit-course.md](./moonbit-course.md) for comprehensive learning.
