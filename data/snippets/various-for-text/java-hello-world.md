<!-- METADATA
{
  "title": "Hello World",
  "tags": [
    "java",
    "basics"
  ],
  "language": "java"
}
-->

# Java Hello World

## Basic Hello World Program

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}
```

## Steps to Create and Run

1. **Create the file**: Save the code as `HelloWorld.java` (the filename must match the class name)

2. **Compile the program**:
   ```bash
   javac HelloWorld.java
   ```
   This creates a `HelloWorld.class` file

3. **Run the program**:
   ```bash
   java HelloWorld
   ```

## Explanation

- `public class HelloWorld`: Defines a public class named HelloWorld
- `public static void main(String[] args)`: The main method, entry point of the program
- `System.out.println()`: Prints text to the console with a newline

## Modern Java (Java 11+)

You can run Java files directly without explicit compilation:

```bash
java HelloWorld.java
```
