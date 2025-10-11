<!-- METADATA
{
  "title": "Python Dictionaries",
  "tags": [
    "python",
    "io"
  ],
  "language": "python"
}
-->

## Dictionaries
Creating and manipulating dictionaries
```python
person = {"name": "Alice", "age": 30, "city": "New York"}

# Dictionary operations
person["email"] = "alice@example.com"
age = person.get("age", 0)
keys = list(person.keys())
values = list(person.values())

# Dictionary comprehension
squares_dict = {x: x**2 for x in range(5)}

# Iterating
for key, value in person.items():
    print(f"{key}: {value}")
```