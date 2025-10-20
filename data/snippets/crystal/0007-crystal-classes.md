<!-- METADATA
{
  "title": "Crystal Classes",
  "tags": [
    "crystal",
    "oop",
    "classes"
  ],
  "language": "crystal"
}
-->

## Classes
Class definition and instantiation
```crystal
class Person
  # Properties with getter/setter
  property name : String
  property age : Int32

  # Read-only property
  getter id : Int32

  def initialize(@name, @age)
    @id = Random.rand(1000..9999)
  end

  def greet
    "Hi, I'm #{@name}, #{@age} years old"
  end

  def birthday
    @age += 1
  end
end

# Usage
person = Person.new("Alice", 30)
puts person.name      # => "Alice"
puts person.greet     # => "Hi, I'm Alice, 30 years old"
person.birthday
puts person.age       # => 31
```
