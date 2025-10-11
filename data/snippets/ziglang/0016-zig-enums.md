<!-- METADATA
{
  "title": "Zig Enums",
  "tags": [
    "zig",
    "enums",
    "data-structures"
  ],
  "language": "zig"
}
-->

## Enums
Defining and using enumerations
```zig
const std = @import("std");

const Color = enum {
    red,
    green,
    blue,

    pub fn toString(self: Color) []const u8 {
        return switch (self) {
            .red => "Red",
            .green => "Green",
            .blue => "Blue",
        };
    }
};

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const color = Color.red;

    try stdout.print("Color: {s}\n", .{color.toString()});

    switch (color) {
        .red => try stdout.print("Stop!\n", .{}),
        .green => try stdout.print("Go!\n", .{}),
        .blue => try stdout.print("Blue!\n", .{}),
    }
}
```
