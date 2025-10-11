<!-- METADATA
{
  "title": "Zig Variable Declaration",
  "tags": [
    "zig",
    "variables",
    "io"
  ],
  "language": "zig"
}
-->

## Variable Declaration
Different ways to declare and initialize variables
```zig
const std = @import("std");

pub fn main() !void {
    const name: []const u8 = "John";  // Immutable (const)
    var age: i32 = 30;                // Mutable (var)
    var city = "New York";            // Type inference

    const stdout = std.io.getStdOut().writer();
    try stdout.print("Name: {s}, Age: {d}\n", .{name, age});
}
```
