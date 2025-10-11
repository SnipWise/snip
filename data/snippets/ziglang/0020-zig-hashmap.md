<!-- METADATA
{
  "title": "Zig HashMap",
  "tags": [
    "zig",
    "hashmap",
    "data-structures"
  ],
  "language": "zig"
}
-->

## HashMap
Using HashMap for key-value storage
```zig
const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var map = std.StringHashMap(i32).init(allocator);
    defer map.deinit();

    // Insert key-value pairs
    try map.put("age", 30);
    try map.put("year", 2024);

    // Get value
    if (map.get("age")) |value| {
        std.debug.print("Age: {d}\n", .{value});
    }

    // Iterate
    var iter = map.iterator();
    while (iter.next()) |entry| {
        std.debug.print("{s} = {d}\n", .{entry.key_ptr.*, entry.value_ptr.*});
    }
}
```
