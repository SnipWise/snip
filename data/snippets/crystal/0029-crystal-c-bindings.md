<!-- METADATA
{
  "title": "Crystal C Bindings",
  "tags": [
    "crystal",
    "c",
    "ffi"
  ],
  "language": "crystal"
}
-->

## C Bindings
Interfacing with C libraries
```crystal
# Basic C binding
@[Link("m")]  # Link with math library
lib LibM
  fun sqrt(x : Float64) : Float64
  fun pow(x : Float64, y : Float64) : Float64
end

puts LibM.sqrt(16.0)  # => 4.0
puts LibM.pow(2.0, 3.0)  # => 8.0

# C struct binding
lib LibC
  struct TimeVal
    tv_sec : Int64
    tv_usec : Int64
  end

  fun gettimeofday(tv : TimeVal*, tz : Void*) : Int32
end

tv = LibC::TimeVal.new
LibC.gettimeofday(pointerof(tv), nil)
puts tv.tv_sec

# Callback from C
lib LibCallback
  alias Callback = (Int32 -> Int32)
  fun register_callback(cb : Callback)
end

callback = ->(x : Int32) { x * 2 }
LibCallback.register_callback(callback)
```
