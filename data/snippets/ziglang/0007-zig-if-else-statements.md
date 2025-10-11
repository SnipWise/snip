<!-- METADATA
{
  "title": "Zig If-Else Statements",
  "tags": [
    "zig",
    "conditionals",
    "control-flow"
  ],
  "language": "zig"
}
-->

## If-Else Statements
Conditional statements in Zig
```zig
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const age: u32 = 25;

    if (age >= 18) {
        try stdout.print("Adult\n", .{});
    } else {
        try stdout.print("Minor\n", .{});
    }

    // If as expression
    const status = if (age >= 18) "adult" else "minor";
    try stdout.print("Status: {s}\n", .{status});

    // Multiple conditions
    if (age < 13) {
        try stdout.print("Child\n", .{});
    } else if (age < 18) {
        try stdout.print("Teenager\n", .{});
    } else {
        try stdout.print("Adult\n", .{});
    }
}
```
