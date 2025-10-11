<!-- METADATA
{
  "title": "Rustlang Configuration",
  "tags": [
    "rust",
    "io",
    "configuration"
  ],
  "language": "rust"
}
-->

## Configuration
Managing configuration with config crate
```rust
use serde::Deserialize;

#[derive(Debug, Deserialize)]
struct ServerConfig {
    host: String,
    port: u16,
}

#[derive(Debug, Deserialize)]
struct DatabaseConfig {
    url: String,
    max_connections: u32,
}

#[derive(Debug, Deserialize)]
struct AppConfig {
    server: ServerConfig,
    database: DatabaseConfig,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut settings = config::Config::default();
    
    // Add configuration sources
    settings
        .merge(config::File::with_name("config"))?
        .merge(config::Environment::with_prefix("APP"))?;
    
    // Deserialize configuration
    let app_config: AppConfig = settings.try_into()?;
    
    println!("Config: {:?}", app_config);
    println!("Server will run on {}:{}", app_config.server.host, app_config.server.port);
    
    Ok(())
}
```