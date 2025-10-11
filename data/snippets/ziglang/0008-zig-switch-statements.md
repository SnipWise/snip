<!-- METADATA
{
  "title": "Zig Switch Statements",
  "tags": [
    "zig",
    "switch",
    "control-flow"
  ],
  "language": "zig"
}
-->

## Switch Statements
Using switch for pattern matching
```zig
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const value: i32 = 2;

    // Switch statement
    switch (value) {
        1 => try stdout.print("One\n", .{}),
        2 => try stdout.print("Two\n", .{}),
        3...5 => try stdout.print("Three to Five\n", .{}),
        else => try stdout.print("Other\n", .{}),
    }

    // Switch as expression
    const result = switch (value) {
        1 => "one",
        2 => "two",
        else => "other",
    };
    try stdout.print("Result: {s}\n", .{result});
}
```
