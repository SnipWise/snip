## Arrays and Vectors
Working with arrays and vectors
```rust
fn main() {
    let arr: [i32; 3] = [1, 2, 3];
    let mut vec = vec![1, 2, 3, 4];
    
    vec.push(5);
    vec.pop();
    
    println!("Array: {:?}", arr);
    println!("Vector: {:?}, len: {}", vec, vec.len());
    
    for (i, value) in vec.iter().enumerate() {
        println!("Index {}: {}", i, value);
    }
}
```