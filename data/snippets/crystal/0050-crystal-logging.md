<!-- METADATA
{
  "title": "Crystal Logging",
  "tags": [
    "crystal",
    "logging",
    "debugging"
  ],
  "language": "crystal"
}
-->

## Logging
Application logging
```crystal
require "log"

# Basic logging
Log.info { "Application started" }
Log.debug { "Debug information" }
Log.warn { "Warning message" }
Log.error { "Error occurred" }
Log.fatal { "Fatal error" }

# Create named logger
log = Log.for("my_app")
log.info { "My app started" }

# Configure log level
Log.setup(:debug)  # Show all messages
Log.setup(:info)   # Show info and above
Log.setup(:error)  # Show only errors

# Custom backend
backend = Log::IOBackend.new
backend.formatter = Log::Formatter.new do |entry, io|
  io << "[#{entry.severity}] "
  io << entry.timestamp.to_s("%Y-%m-%d %H:%M:%S")
  io << " - " << entry.message
end

Log.setup do |config|
  config.bind("*", :debug, backend)
end

# Log to file
file_backend = Log::IOBackend.new(File.new("app.log", "a"))
Log.setup do |config|
  config.bind("*", :info, file_backend)
end

# Structured logging
log.info(user_id: 123, action: "login") { "User logged in" }

# Context
Log.context.set(request_id: "abc123")
log.info { "Processing request" }
```
