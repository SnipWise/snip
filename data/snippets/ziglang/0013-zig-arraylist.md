<!-- METADATA
{
  "title": "Zig ArrayList",
  "tags": [
    "zig",
    "arraylist",
    "data-structures"
  ],
  "language": "zig"
}
-->

## ArrayList
Dynamic array using ArrayList
```zig
const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var list = std.ArrayList(i32).init(allocator);
    defer list.deinit();

    // Add items
    try list.append(1);
    try list.append(2);
    try list.append(3);

    // Access items
    for (list.items) |item| {
        std.debug.print("{d} ", .{item});
    }
    std.debug.print("\n", .{});

    // Pop item
    _ = list.pop();
    std.debug.print("After pop: {any}\n", .{list.items});
}
```
