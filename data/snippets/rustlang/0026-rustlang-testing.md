## Testing
Unit tests and integration tests
```rust
fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err("Division by zero".to_string())
    } else {
        Ok(a / b)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add() {
        assert_eq!(add(2, 3), 5);
        assert_eq!(add(-1, 1), 0);
    }

    #[test]
    fn test_divide() {
        assert_eq!(divide(10.0, 2.0), Ok(5.0));
        assert!(divide(10.0, 0.0).is_err());
    }

    #[test]
    #[should_panic]
    fn test_panic() {
        panic!("This test should panic");
    }
}

fn main() {
    println!("Run with: cargo test");
}
```