<!-- METADATA
{
  "title": "Python Command Line Arguments",
  "tags": [
    "python",
    "cli"
  ],
  "language": "python"
}
-->

## Command Line Arguments
Parsing command line arguments with argparse
```python
import argparse

def main():
    parser = argparse.ArgumentParser(description="Sample CLI app")
    
    parser.add_argument(
        "-n", "--name",
        default="World",
        help="Name to greet"
    )
    
    parser.add_argument(
        "-a", "--age",
        type=int,
        help="Age of the person"
    )
    
    parser.add_argument(
        "-v", "--verbose",
        action="store_true",
        help="Enable verbose output"
    )
    
    parser.add_argument(
        "files",
        nargs="*",
        help="Files to process"
    )
    
    args = parser.parse_args()
    
    if args.verbose:
        print(f"Arguments: {vars(args)}")
    
    greeting = f"Hello, {args.name}!"
    if args.age:
        greeting += f" You are {args.age} years old."
    
    print(greeting)
    
    if args.files:
        print(f"Processing files: {args.files}")

if __name__ == "__main__":
    main()
```