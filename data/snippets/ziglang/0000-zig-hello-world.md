<!-- METADATA
{
  "title": "Zig Hello World",
  "tags": [
    "zig",
    "basics"
  ],
  "language": "zig"
}
-->

## Hello World
Basic program structure and main function
```zig
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    try stdout.print("Hello, World!\n", .{});
}
```
