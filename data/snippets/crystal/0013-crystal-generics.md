<!-- METADATA
{
  "title": "Crystal Generics",
  "tags": [
    "crystal",
    "generics",
    "types"
  ],
  "language": "crystal"
}
-->

## Generics
Generic types and methods
```crystal
# Generic class
class Box(T)
  property value : T

  def initialize(@value)
  end

  def update(new_value : T)
    @value = new_value
  end
end

# Usage
int_box = Box(Int32).new(42)
str_box = Box(String).new("Hello")

# Generic method
def first(array : Array(T)) : T forall T
  array[0]
end

numbers = [1, 2, 3]
puts first(numbers)  # => 1

# Multiple type parameters
class Pair(K, V)
  property key : K
  property value : V

  def initialize(@key, @value)
  end
end

pair = Pair(String, Int32).new("age", 30)
```
