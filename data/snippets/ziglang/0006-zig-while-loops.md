<!-- METADATA
{
  "title": "Zig While Loops",
  "tags": [
    "zig",
    "loops",
    "control-flow"
  ],
  "language": "zig"
}
-->

## While Loops
Using while loops and continue expressions
```zig
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    var i: u32 = 0;
    while (i < 5) : (i += 1) {
        try stdout.print("{d} ", .{i});
    }
    try stdout.print("\n", .{});

    // While with condition
    var x: u32 = 0;
    while (x < 10) {
        x += 1;
        if (x == 5) break;
    }
    try stdout.print("x = {d}\n", .{x});
}
```
