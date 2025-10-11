<!-- METADATA
{
  "title": "Python Classes And Objects",
  "tags": [
    "python",
    "classes"
  ],
  "language": "python"
}
-->

## Classes and Objects
Object-oriented programming
```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
        self._id = id(self)  # Private-ish attribute

    def greet(self):
        return f"Hello, I'm {self.name}"

    def __str__(self):
        return f"Person(name={self.name}, age={self.age})"

    @classmethod
    def from_string(cls, person_str):
        name, age = person_str.split('-')
        return cls(name, int(age))

    @staticmethod
    def is_adult(age):
        return age >= 18

# Usage
person = Person("Alice", 30)
print(person.greet())
print(f"Is adult: {Person.is_adult(person.age)}")

person2 = Person.from_string("Bob-25")
print(person2)
```