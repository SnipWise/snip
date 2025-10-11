## Context Managers
Using and creating context managers
```python
import contextlib
import os

# File context manager (built-in)
with open("temp.txt", "w") as file:
    file.write("Hello, context manager!")

# Custom context manager class
class DatabaseConnection:
    def __init__(self, db_name):
        self.db_name = db_name
        self.connection = None

    def __enter__(self):
        print(f"Connecting to {self.db_name}")
        self.connection = f"connection_to_{self.db_name}"
        return self.connection

    def __exit__(self, exc_type, exc_val, exc_tb):
        print(f"Closing connection to {self.db_name}")
        if exc_type:
            print(f"Exception occurred: {exc_val}")
        return False

# Context manager using contextlib
@contextlib.contextmanager
def temporary_file(filename):
    print(f"Creating {filename}")
    try:
        with open(filename, "w") as f:
            yield f
    finally:
        print(f"Cleaning up {filename}")
        if os.path.exists(filename):
            os.remove(filename)

# Usage
with DatabaseConnection("users_db") as conn:
    print(f"Using {conn}")

with temporary_file("temp_data.txt") as f:
    f.write("Temporary data")
    print("File created and used")
```