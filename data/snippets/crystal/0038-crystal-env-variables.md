<!-- METADATA
{
  "title": "Crystal Environment Variables",
  "tags": [
    "crystal",
    "system",
    "environment"
  ],
  "language": "crystal"
}
-->

## Environment Variables
Advanced environment variable handling
```crystal
# Get environment variable
home = ENV["HOME"]
path = ENV["PATH"]

# Get with default
port = ENV.fetch("PORT", "3000")
debug = ENV.fetch("DEBUG", "false")

# Check existence
if ENV.has_key?("DATABASE_URL")
  puts "Database configured"
end

# Set environment variable
ENV["MY_VAR"] = "my_value"

# Delete variable
ENV.delete("MY_VAR")

# Iterate all variables
ENV.each do |key, value|
  puts "#{key} = #{value}"
end

# Convert to boolean
def env_bool(key, default = false)
  value = ENV[key]?
  return default unless value
  ["true", "1", "yes"].includes?(value.downcase)
end

is_debug = env_bool("DEBUG")

# Parse integer
def env_int(key, default)
  ENV[key]?.try(&.to_i) || default
end

max_connections = env_int("MAX_CONN", 10)
```
