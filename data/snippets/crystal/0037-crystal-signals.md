<!-- METADATA
{
  "title": "Crystal Signal Handling",
  "tags": [
    "crystal",
    "system",
    "signals"
  ],
  "language": "crystal"
}
-->

## Signal Handling
Handling OS signals
```crystal
# Handle SIGINT (Ctrl+C)
Signal::INT.trap do
  puts "\nCaught interrupt signal"
  exit
end

# Handle SIGTERM
Signal::TERM.trap do
  puts "Received termination signal"
  # Cleanup code
  exit
end

# Multiple signals
[Signal::INT, Signal::TERM].each do |sig|
  sig.trap do
    puts "Shutting down gracefully..."
    exit
  end
end

# Reset signal handler
Signal::INT.reset

# Ignore signal
Signal::INT.ignore

# Example: graceful shutdown
running = true

Signal::INT.trap do
  puts "\nShutting down..."
  running = false
end

while running
  puts "Working..."
  sleep 1
end

puts "Cleanup complete"
```
