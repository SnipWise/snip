<!-- METADATA
{
  "title": "Zig String Operations",
  "tags": [
    "zig",
    "strings",
    "text"
  ],
  "language": "zig"
}
-->

## String Operations
Working with strings and string manipulation
```zig
const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // String concatenation
    const str1 = "Hello, ";
    const str2 = "World!";
    const result = try std.mem.concat(allocator, u8, &[_][]const u8{str1, str2});
    defer allocator.free(result);

    std.debug.print("{s}\n", .{result});

    // String comparison
    const are_equal = std.mem.eql(u8, str1, "Hello, ");
    std.debug.print("Equal: {}\n", .{are_equal});

    // String contains
    const contains = std.mem.indexOf(u8, result, "World") != null;
    std.debug.print("Contains 'World': {}\n", .{contains});
}
```
