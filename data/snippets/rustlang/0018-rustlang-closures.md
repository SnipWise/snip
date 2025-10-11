## Closures
Using closures and functional programming
```rust
fn main() {
    let numbers = vec![1, 2, 3, 4, 5];
    
    // Closure that captures environment
    let multiplier = 2;
    let doubled: Vec<i32> = numbers.iter().map(|x| x * multiplier).collect();
    println!("Doubled: {:?}", doubled);
    
    // Filter and collect
    let evens: Vec<i32> = numbers.into_iter().filter(|&x| x % 2 == 0).collect();
    println!("Evens: {:?}", evens);
    
    // Using closure as parameter
    let add_one = |x| x + 1;
    println!("5 + 1 = {}", add_one(5));
}
```