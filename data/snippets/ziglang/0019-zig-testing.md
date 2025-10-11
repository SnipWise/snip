<!-- METADATA
{
  "title": "Zig Testing",
  "tags": [
    "zig",
    "testing",
    "test"
  ],
  "language": "zig"
}
-->

## Testing
Writing unit tests
```zig
const std = @import("std");
const expect = std.testing.expect;
const expectEqual = std.testing.expectEqual;

fn add(a: i32, b: i32) i32 {
    return a + b;
}

test "addition" {
    try expectEqual(@as(i32, 5), add(2, 3));
    try expectEqual(@as(i32, 0), add(-1, 1));
}

test "basic assertions" {
    try expect(true);
    try expect(5 > 3);
    try expect(add(1, 1) == 2);
}

test "string equality" {
    const str1 = "hello";
    const str2 = "hello";
    try expect(std.mem.eql(u8, str1, str2));
}
```
