## String Manipulation
Common string operations
```python
text = "  Hello, World!  "
name = "Alice"
age = 30

# Basic operations
print(f"Original: '{text}'")
print(f"Trimmed: '{text.strip()}'")
print(f"Upper: {text.upper()}")
print(f"Lower: {text.lower()}")
print(f"Replace: {text.replace('World', 'Python')}")

# String formatting
formatted1 = "Name: {}, Age: {}".format(name, age)
formatted2 = f"Name: {name}, Age: {age}"
formatted3 = "Name: %(name)s, Age: %(age)d" % {"name": name, "age": age}

print(f"Formatted: {formatted2}")

# String methods
sentence = "python is awesome"
words = sentence.split()
capitalized = sentence.title()
starts_with = sentence.startswith("python")

print(f"Words: {words}")
print(f"Capitalized: {capitalized}")
print(f"Starts with 'python': {starts_with}")

# Join strings
joined = " | ".join(words)
print(f"Joined: {joined}")
```