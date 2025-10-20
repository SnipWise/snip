<!-- METADATA
{
  "title": "Crystal UUID",
  "tags": [
    "crystal",
    "uuid",
    "identifiers"
  ],
  "language": "crystal"
}
-->

## UUID
Generating unique identifiers
```crystal
require "uuid"

# Generate random UUID (v4)
id = UUID.random
puts id  # => "550e8400-e29b-41d4-a716-446655440000"

# Convert to string
id_string = id.to_s

# Parse UUID from string
parsed = UUID.new("550e8400-e29b-41d4-a716-446655440000")

# Check validity
begin
  UUID.new("invalid-uuid")
rescue ArgumentError
  puts "Invalid UUID format"
end

# Compare UUIDs
uuid1 = UUID.random
uuid2 = UUID.random
puts uuid1 == uuid2  # => false

# Convert to bytes
bytes = id.bytes

# Nil UUID (all zeros)
nil_uuid = UUID.empty
puts nil_uuid  # => "00000000-0000-0000-0000-000000000000"

# Use in classes
class User
  property id : UUID
  property name : String

  def initialize(@name)
    @id = UUID.random
  end
end

user = User.new("Alice")
puts user.id
```
