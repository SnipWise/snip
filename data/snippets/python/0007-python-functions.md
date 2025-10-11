## Functions
Function definition and parameters
```python
def greet(name, greeting="Hello"):
    return f"{greeting}, {name}!"

def add_numbers(*args):
    return sum(args)

def create_profile(**kwargs):
    return {k: v for k, v in kwargs.items()}

# Lambda functions
square = lambda x: x**2
add = lambda x, y: x + y

# Function calls
message = greet("Alice")
total = add_numbers(1, 2, 3, 4, 5)
profile = create_profile(name="Bob", age=25, city="Boston")

print(f"Message: {message}")
print(f"Total: {total}")
print(f"Square of 5: {square(5)}")
```