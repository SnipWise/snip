<!-- METADATA
{
  "title": "Crystal Pointers",
  "tags": [
    "crystal",
    "pointers",
    "unsafe"
  ],
  "language": "crystal"
}
-->

## Pointers
Low-level pointer operations
```crystal
# Get pointer to variable
x = 42
ptr = pointerof(x)
puts ptr.value  # => 42

# Modify through pointer
ptr.value = 100
puts x  # => 100

# Pointer arithmetic
numbers = [1, 2, 3, 4, 5]
ptr = numbers.to_unsafe
puts ptr[0]  # => 1
puts ptr[1]  # => 2

# Allocate memory
buffer = Pointer(Int32).malloc(10)
buffer[0] = 42
puts buffer[0]  # => 42

# C bindings
lib LibC
  fun malloc(size : UInt32) : Void*
  fun free(ptr : Void*)
end

ptr = LibC.malloc(100)
LibC.free(ptr)

# Slice from pointer
slice = Slice.new(ptr.as(Int32*), 10)
```
