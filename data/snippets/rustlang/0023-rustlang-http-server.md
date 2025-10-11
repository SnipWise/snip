<!-- METADATA
{
  "title": "Rustlang Http Server",
  "tags": [
    "rust",
    "http",
    "networking"
  ],
  "language": "rust"
}
-->

## HTTP Server
Creating a simple HTTP server with warp
```rust
use warp::Filter;

#[tokio::main]
async fn main() {
    let hello = warp::path!("hello" / String)
        .map(|name| format!("Hello, {}!", name));
    
    let json_route = warp::path("json")
        .map(|| warp::reply::json(&serde_json::json!({
            "message": "Hello, JSON!"
        })));
    
    let routes = hello.or(json_route);
    
    println!("Server starting on localhost:3030");
    warp::serve(routes)
        .run(([127, 0, 0, 1], 3030))
        .await;
}
```