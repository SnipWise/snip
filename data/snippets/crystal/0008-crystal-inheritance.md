<!-- METADATA
{
  "title": "Crystal Inheritance",
  "tags": [
    "crystal",
    "oop",
    "inheritance"
  ],
  "language": "crystal"
}
-->

## Inheritance
Class inheritance and method overriding
```crystal
class Animal
  property name : String

  def initialize(@name)
  end

  def speak
    "Some sound"
  end
end

class Dog < Animal
  def speak
    "Woof!"
  end

  def fetch
    "#{@name} is fetching the ball"
  end
end

class Cat < Animal
  def speak
    "Meow!"
  end
end

# Usage
dog = Dog.new("Rex")
cat = Cat.new("Whiskers")

puts dog.speak    # => "Woof!"
puts cat.speak    # => "Meow!"
puts dog.fetch    # => "Rex is fetching the ball"
```
