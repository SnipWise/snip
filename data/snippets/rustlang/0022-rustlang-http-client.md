## HTTP Client
Making HTTP requests with reqwest
```rust
use reqwest;
use tokio;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let response = reqwest::get("https://httpbin.org/json").await?;
    let status = response.status();
    let body = response.text().await?;
    
    println!("Status: {}", status);
    println!("Body: {}", body);
    
    // POST request with JSON
    let client = reqwest::Client::new();
    let json_body = serde_json::json!({
        "name": "John",
        "age": 30
    });
    
    let post_response = client
        .post("https://httpbin.org/post")
        .json(&json_body)
        .send()
        .await?;
    
    println!("POST Status: {}", post_response.status());
    Ok(())
}
```