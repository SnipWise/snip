## Logging
Structured logging with log and env_logger
```rust
use log::{info, warn, error, debug};

fn main() {
    env_logger::init();
    
    debug!("This is a debug message");
    info!("Application started");
    warn!("This is a warning");
    error!("This is an error message");
    
    // Structured logging
    info!("User logged in"; "user_id" => 123, "ip" => "192.168.1.1");
    
    // Log with formatting
    let user_name = "Alice";
    let login_count = 5;
    info!("User {} has logged in {} times", user_name, login_count);
}
```