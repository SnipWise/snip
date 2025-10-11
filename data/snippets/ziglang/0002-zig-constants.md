<!-- METADATA
{
  "title": "Zig Constants",
  "tags": [
    "zig",
    "constants"
  ],
  "language": "zig"
}
-->

## Constants
Compile-time constants and runtime constants
```zig
const std = @import("std");

const PI: f64 = 3.14159;              // Compile-time constant
const MAX_SIZE: comptime_int = 100;   // Compile-time integer

pub fn main() !void {
    const message = "Hello";          // Runtime constant

    const stdout = std.io.getStdOut().writer();
    try stdout.print("PI: {d}, Max: {d}, Message: {s}\n", .{PI, MAX_SIZE, message});
}
```
