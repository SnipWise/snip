<!-- METADATA
{
  "title": "Rustlang Data Types",
  "tags": [
    "rust"
  ],
  "language": "rust"
}
-->

## Data Types
Working with basic data types
```rust
fn main() {
    let integer: i32 = 42;
    let float: f64 = 3.14;
    let boolean: bool = true;
    let character: char = '🦀';
    let string: String = String::from("Rust");
    let str_slice: &str = "slice";
    
    println!("int: {}, float: {}, bool: {}, char: {}", integer, float, boolean, character);
    println!("String: {}, &str: {}", string, str_slice);
}
```