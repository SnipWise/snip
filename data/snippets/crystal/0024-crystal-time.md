<!-- METADATA
{
  "title": "Crystal Time and Date",
  "tags": [
    "crystal",
    "time",
    "date"
  ],
  "language": "crystal"
}
-->

## Time and Date
Working with time and dates
```crystal
require "time"

# Current time
now = Time.local
puts now

# UTC time
utc = Time.utc
puts utc

# Create specific time
birthday = Time.local(1990, 6, 15, 10, 30, 0)
puts birthday

# Format time
formatted = now.to_s("%Y-%m-%d %H:%M:%S")
puts formatted  # => "2025-10-20 14:30:45"

# Parse time
parsed = Time.parse("2025-12-25 10:00:00", "%Y-%m-%d %H:%M:%S", Time::Location.local)

# Time arithmetic
tomorrow = now + 1.day
next_week = now + 7.days
one_hour_ago = now - 1.hour

# Compare times
if tomorrow > now
  puts "Tomorrow is in the future"
end

# Unix timestamp
timestamp = now.to_unix
puts timestamp
```
