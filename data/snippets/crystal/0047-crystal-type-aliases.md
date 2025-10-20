<!-- METADATA
{
  "title": "Crystal Type Aliases",
  "tags": [
    "crystal",
    "types",
    "aliases"
  ],
  "language": "crystal"
}
-->

## Type Aliases
Creating type aliases for readability
```crystal
# Simple alias
alias UserId = Int32
alias UserName = String

def create_user(id : UserId, name : UserName)
  {id: id, name: name}
end

user = create_user(1, "Alice")

# Complex type alias
alias JsonValue = String | Int64 | Float64 | Bool | Nil |
                  Array(JsonValue) | Hash(String, JsonValue)

# Generic alias
alias StringHash = Hash(String, String)
alias IntArray = Array(Int32)

config : StringHash = {"host" => "localhost"}
numbers : IntArray = [1, 2, 3]

# Proc alias
alias Handler = Proc(String, String)

def register(handler : Handler)
  result = handler.call("input")
  puts result
end

my_handler : Handler = ->(s : String) { s.upcase }
register(my_handler)

# Tuple alias
alias Point2D = Tuple(Float64, Float64)
alias Point3D = Tuple(Float64, Float64, Float64)

point : Point2D = {10.0, 20.0}

# Named tuple alias
alias UserData = NamedTuple(id: Int32, name: String, active: Bool)

user_data : UserData = {id: 1, name: "Bob", active: true}
```
