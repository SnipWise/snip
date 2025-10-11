## Inheritance
Class inheritance and method overriding
```python
class Animal:
    def __init__(self, name):
        self.name = name

    def speak(self):
        pass

    def info(self):
        return f"This is {self.name}"

class Dog(Animal):
    def __init__(self, name, breed):
        super().__init__(name)
        self.breed = breed

    def speak(self):
        return "Woof!"

    def info(self):
        return f"{super().info()}, a {self.breed} dog"

class Cat(Animal):
    def speak(self):
        return "Meow!"

# Usage
dog = Dog("Buddy", "Golden Retriever")
cat = Cat("Whiskers")

print(f"{dog.name} says: {dog.speak()}")
print(dog.info())
print(f"{cat.name} says: {cat.speak()}")
```