<!-- METADATA
{
  "title": "Crystal JSON",
  "tags": [
    "crystal",
    "json",
    "serialization"
  ],
  "language": "crystal"
}
-->

## JSON
JSON serialization and deserialization
```crystal
require "json"

# Define a serializable class
class User
  include JSON::Serializable

  property name : String
  property age : Int32
  property email : String

  def initialize(@name, @age, @email)
  end
end

# Serialize to JSON
user = User.new("Alice", 30, "alice@example.com")
json_str = user.to_json
puts json_str
# => {"name":"Alice","age":30,"email":"alice@example.com"}

# Deserialize from JSON
json_data = %({"name":"Bob","age":25,"email":"bob@example.com"})
bob = User.from_json(json_data)
puts bob.name  # => "Bob"

# Parse raw JSON
data = JSON.parse(%({"key": "value"}))
puts data["key"]  # => "value"
```
