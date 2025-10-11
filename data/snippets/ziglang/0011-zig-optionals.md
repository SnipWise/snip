<!-- METADATA
{
  "title": "Zig Optionals",
  "tags": [
    "zig",
    "optionals",
    "null"
  ],
  "language": "zig"
}
-->

## Optionals
Working with optional values
```zig
const std = @import("std");

fn findValue(search: i32) ?i32 {
    if (search > 0) {
        return search * 2;
    }
    return null;
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // Optional value
    var maybe_value: ?i32 = 42;

    // Check if value exists
    if (maybe_value) |value| {
        try stdout.print("Value: {d}\n", .{value});
    } else {
        try stdout.print("No value\n", .{});
    }

    // Provide default with orelse
    const result = findValue(-1) orelse 0;
    try stdout.print("Result: {d}\n", .{result});
}
```
