<!-- METADATA
{
  "title": "Zig Generics",
  "tags": [
    "zig",
    "generics",
    "types"
  ],
  "language": "zig"
}
-->

## Generics
Generic functions and data structures
```zig
const std = @import("std");

fn GenericList(comptime T: type) type {
    return struct {
        items: std.ArrayList(T),

        pub fn init(allocator: std.mem.Allocator) @This() {
            return .{ .items = std.ArrayList(T).init(allocator) };
        }

        pub fn deinit(self: *@This()) void {
            self.items.deinit();
        }

        pub fn add(self: *@This(), item: T) !void {
            try self.items.append(item);
        }
    };
}

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var int_list = GenericList(i32).init(allocator);
    defer int_list.deinit();

    try int_list.add(42);
    std.debug.print("List: {any}\n", .{int_list.items.items});
}
```
