<!-- METADATA
{
  "title": "Crystal Annotations",
  "tags": [
    "crystal",
    "annotations",
    "metadata"
  ],
  "language": "crystal"
}
-->

## Annotations
Metadata with annotations
```crystal
# Define custom annotation
annotation MyAnnotation
end

# JSON annotations
class User
  include JSON::Serializable

  property name : String

  @[JSON::Field(key: "user_age")]
  property age : Int32

  @[JSON::Field(ignore: true)]
  property password : String = ""

  def initialize(@name, @age)
  end
end

json = %({"name": "Alice", "user_age": 30})
user = User.from_json(json)
puts user.age  # => 30

# Custom annotation usage
@[MyAnnotation]
class MyClass
  @[MyAnnotation]
  def my_method
  end
end

# Check annotations at compile time
{% if MyClass.annotation(MyAnnotation) %}
  puts "MyClass has MyAnnotation"
{% end %}
```
