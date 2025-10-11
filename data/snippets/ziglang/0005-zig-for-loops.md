<!-- METADATA
{
  "title": "Zig For Loops",
  "tags": [
    "zig",
    "loops",
    "control-flow"
  ],
  "language": "zig"
}
-->

## For Loops
Different ways to use for loops
```zig
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // Iterate over array
    const items = [_]i32{1, 2, 3, 4, 5};
    for (items) |item| {
        try stdout.print("{d} ", .{item});
    }
    try stdout.print("\n", .{});

    // Iterate with index
    for (items, 0..) |item, i| {
        try stdout.print("Index {d}: {d}\n", .{i, item});
    }

    // Range loop
    for (0..5) |i| {
        try stdout.print("{d} ", .{i});
    }
}
```
