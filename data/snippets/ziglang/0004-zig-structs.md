<!-- METADATA
{
  "title": "Zig Structs",
  "tags": [
    "zig",
    "structs",
    "data-structures"
  ],
  "language": "zig"
}
-->

## Structs
Defining and using structs
```zig
const std = @import("std");

const Person = struct {
    name: []const u8,
    age: u32,

    pub fn greet(self: Person) void {
        std.debug.print("Hello, I'm {s} and I'm {d} years old\n", .{self.name, self.age});
    }
};

pub fn main() !void {
    const person = Person{
        .name = "John",
        .age = 30,
    };

    person.greet();
}
```
