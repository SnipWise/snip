<!-- METADATA
{
  "title": "Zig Allocators",
  "tags": [
    "zig",
    "memory",
    "allocators"
  ],
  "language": "zig"
}
-->

## Allocators
Memory allocation using allocators
```zig
const std = @import("std");

pub fn main() !void {
    // Create a general purpose allocator
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Allocate memory for a single item
    const value = try allocator.create(i32);
    defer allocator.destroy(value);
    value.* = 42;

    // Allocate array
    const array = try allocator.alloc(i32, 5);
    defer allocator.free(array);

    std.debug.print("Value: {d}\n", .{value.*});
}
```
