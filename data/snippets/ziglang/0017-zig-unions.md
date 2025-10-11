<!-- METADATA
{
  "title": "Zig Unions",
  "tags": [
    "zig",
    "unions",
    "data-structures"
  ],
  "language": "zig"
}
-->

## Unions
Tagged and untagged unions
```zig
const std = @import("std");

const Value = union(enum) {
    int: i32,
    float: f64,
    bool: bool,
};

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    const v1 = Value{ .int = 42 };
    const v2 = Value{ .float = 3.14 };

    switch (v1) {
        .int => |val| try stdout.print("Integer: {d}\n", .{val}),
        .float => |val| try stdout.print("Float: {d}\n", .{val}),
        .bool => |val| try stdout.print("Bool: {}\n", .{val}),
    }

    switch (v2) {
        .int => |val| try stdout.print("Integer: {d}\n", .{val}),
        .float => |val| try stdout.print("Float: {d}\n", .{val}),
        .bool => |val| try stdout.print("Bool: {}\n", .{val}),
    }
}
```
