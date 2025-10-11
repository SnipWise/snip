## Iterators
Working with iterators
```rust
fn main() {
    let vec = vec![1, 2, 3, 4, 5];
    
    // Iterator methods
    let sum: i32 = vec.iter().sum();
    let max = vec.iter().max();
    println!("Sum: {}, Max: {:?}", sum, max);
    
    // Chain operations
    let result: Vec<i32> = vec
        .iter()
        .filter(|&&x| x > 2)
        .map(|x| x * 2)
        .collect();
    println!("Filtered and doubled: {:?}", result);
    
    // Custom iterator
    for (i, value) in vec.iter().enumerate() {
        println!("Index {}: {}", i, value);
    }
}
```