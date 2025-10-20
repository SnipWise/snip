<!-- METADATA
{
  "title": "Crystal Macros",
  "tags": [
    "crystal",
    "metaprogramming",
    "macros"
  ],
  "language": "crystal"
}
-->

## Macros
Compile-time metaprogramming
```crystal
# Simple macro
macro define_method(name)
  def {{name}}
    puts "Method {{name}} called"
  end
end

define_method hello
hello  # => "Method hello called"

# Macro with parameters
macro create_property(name, type)
  property {{name}} : {{type}}
end

class User
  create_property name, String
  create_property age, Int32

  def initialize(@name, @age)
  end
end

# Conditional compilation
macro debug(message)
  {% if flag?(:debug) %}
    puts "DEBUG: {{message}}"
  {% end %}
end

debug "This only shows in debug mode"

# Generate methods from array
macro define_getters(*names)
  {% for name in names %}
    def {{name}}
      @{{name}}
    end
  {% end %}
end
```
