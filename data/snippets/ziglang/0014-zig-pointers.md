<!-- METADATA
{
  "title": "Zig Pointers",
  "tags": [
    "zig",
    "pointers",
    "memory"
  ],
  "language": "zig"
}
-->

## Pointers
Working with pointers
```zig
const std = @import("std");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    // Single-item pointer
    var x: i32 = 42;
    const ptr: *i32 = &x;
    try stdout.print("Value: {d}\n", .{ptr.*});

    // Modify through pointer
    ptr.* = 100;
    try stdout.print("Modified: {d}\n", .{x});

    // Const pointer
    const y: i32 = 50;
    const const_ptr: *const i32 = &y;
    try stdout.print("Const value: {d}\n", .{const_ptr.*});

    // Multi-item pointer (slice)
    var arr = [_]i32{1, 2, 3};
    const slice: []i32 = &arr;
    try stdout.print("Slice: {any}\n", .{slice});
}
```
