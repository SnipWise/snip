<!-- METADATA
{
  "title": "Zig Functions",
  "tags": [
    "zig",
    "functions"
  ],
  "language": "zig"
}
-->

## Functions
Function declaration and return values
```zig
const std = @import("std");

fn add(a: i32, b: i32) i32 {
    return a + b;
}

fn greet(name: []const u8) !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Hello, {s}!\n", .{name});
}

pub fn main() !void {
    const sum = add(5, 3);
    std.debug.print("Sum: {d}\n", .{sum});

    try greet("World");
}
```
