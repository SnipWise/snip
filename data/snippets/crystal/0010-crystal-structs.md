<!-- METADATA
{
  "title": "Crystal Structs",
  "tags": [
    "crystal",
    "structs",
    "data-structures"
  ],
  "language": "crystal"
}
-->

## Structs
Value types with struct
```crystal
# Struct definition (passed by value)
struct Point
  property x : Float64
  property y : Float64

  def initialize(@x, @y)
  end

  def distance_from_origin
    Math.sqrt(@x ** 2 + @y ** 2)
  end
end

# Usage
p1 = Point.new(3.0, 4.0)
p2 = p1  # Copy, not reference

p2.x = 10.0
puts p1.x  # => 3.0 (p1 unchanged)
puts p2.x  # => 10.0

puts p1.distance_from_origin  # => 5.0
```
