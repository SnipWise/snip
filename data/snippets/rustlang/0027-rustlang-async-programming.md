<!-- METADATA
{
  "title": "Rustlang Async Programming",
  "tags": [
    "rust",
    "async"
  ],
  "language": "rust"
}
-->

## Async Programming
Asynchronous programming with async/await
```rust
use tokio::time::{sleep, Duration};

async fn fetch_data(id: u32) -> String {
    sleep(Duration::from_millis(100)).await;
    format!("Data for ID: {}", id)
}

async fn process_multiple() {
    let mut handles = vec![];
    
    for i in 1..=3 {
        let handle = tokio::spawn(async move {
            fetch_data(i).await
        });
        handles.push(handle);
    }
    
    for handle in handles {
        match handle.await {
            Ok(result) => println!("{}", result),
            Err(e) => println!("Error: {}", e),
        }
    }
}

#[tokio::main]
async fn main() {
    println!("Starting async operations...");
    process_multiple().await;
    println!("All operations completed");
}
```