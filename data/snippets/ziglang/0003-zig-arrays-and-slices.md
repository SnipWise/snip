<!-- METADATA
{
  "title": "Zig Arrays and Slices",
  "tags": [
    "zig",
    "arrays",
    "slices"
  ],
  "language": "zig"
}
-->

## Arrays and Slices
Working with arrays and slices in Zig
```zig
const std = @import("std");

pub fn main() !void {
    // Fixed-size array
    var arr = [_]i32{1, 2, 3, 4, 5};

    // Slice (view into array)
    var slice: []i32 = arr[1..4];

    // String is a slice
    const message: []const u8 = "Hello";

    const stdout = std.io.getStdOut().writer();
    try stdout.print("Array length: {d}, Slice: {any}\n", .{arr.len, slice});
}
```
