<!-- METADATA
{
  "title": "Python File Io",
  "tags": [
    "python",
    "file-io",
    "io"
  ],
  "language": "python"
}
-->

## File I/O
Reading and writing files
```python
import os

# Writing to file
content = "Hello, Python!\nSecond line"
with open("test.txt", "w") as file:
    file.write(content)

# Reading from file
with open("test.txt", "r") as file:
    file_content = file.read()
    print(f"File content:\n{file_content}")

# Reading lines
with open("test.txt", "r") as file:
    lines = file.readlines()
    for i, line in enumerate(lines, 1):
        print(f"Line {i}: {line.strip()}")

# Appending to file
with open("test.txt", "a") as file:
    file.write("\nAppended line")

# JSON file handling
import json

data = {"name": "Alice", "age": 30, "skills": ["Python", "Java"]}
with open("data.json", "w") as file:
    json.dump(data, file, indent=2)

with open("data.json", "r") as file:
    loaded_data = json.load(file)
    print(f"Loaded data: {loaded_data}")

# Cleanup
os.remove("test.txt")
os.remove("data.json")
```