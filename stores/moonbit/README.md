# MoonBit Learning Resources

This directory contains comprehensive learning materials for MoonBit, a next-generation programming language designed for cloud and edge computing with WebAssembly.

## What is MoonBit?

MoonBit is an AI-native programming language that combines functional, imperative, and object-oriented paradigms. It compiles primarily to WebAssembly with excellent performance and small output size.

### Key Features

- **WebAssembly First**: Native compilation to Wasm with best-in-class performance
- **Multi-Backend**: Compiles to WebAssembly, JavaScript, and native code
- **Built-in Async**: No explicit async/await keywords needed
- **Checked Exceptions**: Compile-time verification of error handling
- **Strong Type System**: Static typing with powerful inference
- **Fast Compilation**: Extremely fast build times

## Learning Materials

### 1. [snippets-moonbit.md](./snippets-moonbit.md)

Quick reference snippets covering:
- Hello World
- Variables and data types
- Control flow (if/else, loops, match)
- Functions and closures
- Structs and enums
- Arrays, lists, and maps
- Pattern matching
- Error handling
- Traits and generics
- Iterators and functional programming
- Async programming
- Testing
- Advanced features (pipe operator, cascade, defer, etc.)

**Best for**: Quick lookups and code examples

### 2. [moonbit-course.md](./moonbit-course.md)

Complete programming course covering:
- Introduction and setup
- Basic syntax and concepts
- Type system and data structures
- Control flow and pattern matching
- Functions and closures
- Structs, enums, and traits
- Generics and type parameters
- Error handling with checked exceptions
- Collections and iterators
- Module system and packages
- Async programming
- Testing and best practices
- Advanced patterns

**Best for**: Systematic learning from beginner to advanced

## Getting Started

### Installation

```bash
# On Unix-like systems (macOS, Linux)
curl -fsSL https://cli.moonbitlang.com/install/unix.sh | bash

# On Windows
powershell -Command "iwr https://cli.moonbitlang.com/install/windows.ps1 -useb | iex"
```

### Your First Program

Create a new project:

```bash
moon new hello_world
cd hello_world
```

Edit `src/main.mbt`:

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

## Quick Examples

### Hello World
```moonbit
fn main {
  println("Hello, World!")
}
```

### Variables
```moonbit
fn main {
  let name = "Alice"
  let mut age = 30
  age = age + 1
  println("Name: \{name}, Age: \{age}")
}
```

### Functions
```moonbit
fn add(a : Int, b : Int) -> Int {
  a + b
}

fn main {
  let sum = add(5, 3)
  println("Sum: \{sum}")
}
```

### Pattern Matching
```moonbit
fn classify(n : Int) -> String {
  match n {
    0 => "zero"
    x if x < 0 => "negative"
    x if x > 0 => "positive"
    _ => "unknown"
  }
}
```

### Error Handling
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

## Unique Features

### 1. Checked Exceptions

MoonBit uses checked exceptions for type-safe error handling:

```moonbit
fn risky_operation() -> String!Error {
  if something_wrong {
    raise MyError("Something went wrong")
  }
  "Success"
}
```

### 2. Built-in Async

No async/await keywords - async is handled automatically:

```moonbit
fn fetch_data(url : String) -> String {
  // This is async - IDE shows it in italics
  "Data from \{url}"
}
```

### 3. Pipe Operator

Clean function composition:

```moonbit
let result = value
  |> double
  |> add_ten
  |> square
```

### 4. Overloaded Literals

Literals adapt to the expected type:

```moonbit
let x : Int = 42
let y : BigInt = 42  // Same literal, different type
```

## Essential Commands

```bash
moon new <name>       # Create new project
moon build            # Build project
moon run main         # Run project
moon test             # Run tests
moon check            # Type check
moon fmt              # Format code
moon doc              # Generate documentation
moon add <package>    # Add dependency
moon update           # Update dependencies
```

## Official Resources

- **Website**: [moonbitlang.com](https://www.moonbitlang.com/)
- **Documentation**: [docs.moonbitlang.com](https://docs.moonbitlang.com/)
- **GitHub**: [github.com/moonbitlang](https://github.com/moonbitlang)
- **Package Registry**: [mooncakes.io](https://mooncakes.io/)
- **Discord**: Join the MoonBit community

## Comparison with Other Languages

| Feature | MoonBit | Rust | TypeScript | Go |
|---------|---------|------|------------|-----|
| Type System | Strong, inferred | Strong, explicit | Structural | Simple |
| Error Handling | Checked exceptions | Result type | Try/catch | Error values |
| Async | Built-in | async/await | async/await | Goroutines |
| Memory Safety | Yes | Yes | Runtime | Runtime |
| WebAssembly | Primary target | Supported | Via compilation | Supported |
| Learning Curve | Moderate | Steep | Easy | Easy |

## Use Cases

- **WebAssembly Applications**: MoonBit's primary target
- **Cloud Computing**: High-performance cloud services
- **Edge Computing**: Lightweight edge applications
- **Data Processing**: Efficient data transformation
- **Web Backends**: Fast, type-safe APIs
- **Smart Contracts**: Blockchain applications
- **System Tools**: CLI tools and utilities

## Project Structure

```
project/
├── moon.mod.json     # Package manifest
├── src/
│   ├── main.mbt      # Main source file
│   ├── lib.mbt       # Library code
│   └── utils.mbt     # Utilities
├── test/
│   └── main_test.mbt # Tests
└── target/           # Build output
```

## Learning Path

1. **Start with snippets-moonbit.md** for quick examples
2. **Read moonbit-course.md** for comprehensive understanding
3. **Build small projects** to practice
4. **Explore the standard library** in the official docs
5. **Join the community** on Discord
6. **Contribute to packages** on mooncakes.io

## Tips for Success

### Coming from Rust
- No explicit lifetimes in most cases
- Simpler syntax
- Checked exceptions instead of Result everywhere
- Built-in async without keywords

### Coming from TypeScript
- Stronger type system
- No null/undefined (use Option)
- Compiles to WebAssembly
- Better performance

### Coming from Python
- Static typing (but inferred)
- Compiled language
- Much faster execution
- Similar functional features

## Common Patterns

### Builder Pattern
```moonbit
struct Config { mut host : String; mut port : Int }

fn Config::new() -> Config {
  { host: "localhost", port: 8080 }
}

let config = Config::new()
  ..set_host("example.com")
  ..set_port(9000)
```

### Newtype Pattern
```moonbit
struct UserId(Int)
struct Email(String)

// Type safety prevents mixing these up
fn send_email(user_id : UserId, email : Email) -> Unit {
  // ...
}
```

### Option Handling
```moonbit
fn find_user(id : Int) -> User? {
  if id == 1 {
    Some({ name: "Alice", age: 30 })
  } else {
    None
  }
}

match find_user(1) {
  Some(user) => println("Found: \{user.name}")
  None => println("Not found")
}
```

## Contributing

Feel free to contribute improvements to these learning materials:
1. Fork the repository
2. Make your changes
3. Submit a pull request

## License

These learning materials are provided as-is for educational purposes.

---

**Ready to start?** Open [snippets-moonbit.md](./snippets-moonbit.md) for quick examples or [moonbit-course.md](./moonbit-course.md) for the complete course!
