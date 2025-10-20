<!-- METADATA
{
  "title": "Crystal Named Tuples",
  "tags": [
    "crystal",
    "tuples",
    "data-structures"
  ],
  "language": "crystal"
}
-->

## Named Tuples
Immutable named tuples
```crystal
# Create named tuple
person = {name: "Alice", age: 30, city: "Paris"}

# Access values
puts person[:name]  # => "Alice"
puts person[:age]   # => 30

# Type safety
person2 : NamedTuple(name: String, age: Int32) = {
  name: "Bob",
  age: 25
}

# Methods
puts person.keys    # => [:name, :age, :city]
puts person.values  # => ["Alice", 30, "Paris"]
puts person.size    # => 3

# Merge tuples
merged = person.merge({country: "France"})
puts merged[:country]  # => "France"

# Convert to hash
hash = person.to_h
puts hash["name"]  # => "Alice"

# Function return
def get_user
  {id: 1, name: "Alice", role: "admin"}
end

user = get_user
puts user[:role]  # => "admin"
```
