<!-- METADATA
{
  "title": "Zig Packed Structs",
  "tags": [
    "zig",
    "packed",
    "structs"
  ],
  "language": "zig"
}
-->

## Packed Structs
Using packed structs for bit manipulation
```zig
const std = @import("std");

const Flags = packed struct {
    read: bool,
    write: bool,
    execute: bool,
    _padding: u5 = 0,
};

const Color = packed struct {
    r: u8,
    g: u8,
    b: u8,
    a: u8,
};

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    const flags = Flags{
        .read = true,
        .write = false,
        .execute = true,
    };

    const flags_byte: u8 = @bitCast(flags);
    try stdout.print("Flags as byte: 0b{b:0>8}\n", .{flags_byte});

    const color = Color{ .r = 255, .g = 128, .b = 64, .a = 255 };
    const color_int: u32 = @bitCast(color);
    try stdout.print("Color as int: 0x{X:0>8}\n", .{color_int});
}
```
