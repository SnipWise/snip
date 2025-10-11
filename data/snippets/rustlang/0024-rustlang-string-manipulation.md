## String Manipulation
Common string operations
```rust
fn main() {
    let text = "  Hello, World!  ";
    
    println!("Original: '{}'", text);
    println!("Trimmed: '{}'", text.trim());
    println!("Uppercase: {}", text.to_uppercase());
    println!("Contains 'World': {}", text.contains("World"));
    
    let words: Vec<&str> = text.trim().split_whitespace().collect();
    println!("Words: {:?}", words);
    
    // String manipulation
    let mut s = String::from("Hello");
    s.push_str(", World!");
    s.push('!');
    println!("Built string: {}", s);
    
    // Parse string to number
    let num: i32 = "42".parse().unwrap_or(0);
    println!("Parsed number: {}", num);
}
```