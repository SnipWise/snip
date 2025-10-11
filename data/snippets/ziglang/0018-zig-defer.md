<!-- METADATA
{
  "title": "Zig Defer",
  "tags": [
    "zig",
    "defer",
    "control-flow"
  ],
  "language": "zig"
}
-->

## Defer
Using defer for cleanup
```zig
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // defer executes at end of scope
    {
        defer try stdout.print("Third\n", .{});
        defer try stdout.print("Second\n", .{});
        try stdout.print("First\n", .{});
    }

    // Practical use with allocator
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const data = try allocator.alloc(i32, 5);
    defer allocator.free(data);

    try stdout.print("Allocated array of size: {d}\n", .{data.len});
}
```
