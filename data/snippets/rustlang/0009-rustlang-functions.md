## Functions
Function declaration and parameters
```rust
fn add(a: i32, b: i32) -> i32 {
    a + b // No semicolon = return value
}

fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err("Division by zero".to_string())
    } else {
        Ok(a / b)
    }
}

fn main() {
    let sum = add(5, 3);
    println!("Sum: {}", sum);
    
    match divide(10.0, 2.0) {
        Ok(result) => println!("Division: {}", result),
        Err(e) => println!("Error: {}", e),
    }
}
```