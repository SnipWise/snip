<!-- METADATA
{
  "title": "Zig JSON",
  "tags": [
    "zig",
    "json",
    "serialization"
  ],
  "language": "zig"
}
-->

## JSON
Parsing and stringifying JSON
```zig
const std = @import("std");

const Person = struct {
    name: []const u8,
    age: u32,
};

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Parse JSON
    const json_string = "{\"name\":\"John\",\"age\":30}";
    const parsed = try std.json.parseFromSlice(Person, allocator, json_string, .{});
    defer parsed.deinit();

    std.debug.print("Name: {s}, Age: {d}\n", .{parsed.value.name, parsed.value.age});

    // Stringify to JSON
    const person = Person{ .name = "Alice", .age = 25 };
    var string = std.ArrayList(u8).init(allocator);
    defer string.deinit();

    try std.json.stringify(person, .{}, string.writer());
    std.debug.print("JSON: {s}\n", .{string.items});
}
```
