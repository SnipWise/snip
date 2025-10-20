<!-- METADATA
{
  "title": "Crystal Modules",
  "tags": [
    "crystal",
    "modules",
    "mixins"
  ],
  "language": "crystal"
}
-->

## Modules
Modules as mixins and namespaces
```crystal
# Module as mixin
module Swimmable
  def swim
    "Swimming!"
  end
end

module Flyable
  def fly
    "Flying!"
  end
end

class Duck
  include Swimmable
  include Flyable

  def quack
    "Quack!"
  end
end

duck = Duck.new
puts duck.swim   # => "Swimming!"
puts duck.fly    # => "Flying!"

# Module as namespace
module Math
  PI = 3.14159

  def self.circle_area(radius)
    PI * radius ** 2
  end
end

area = Math.circle_area(5)
```
