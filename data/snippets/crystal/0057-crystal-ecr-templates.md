<!-- METADATA
{
  "title": "Crystal ECR Templates",
  "tags": [
    "crystal",
    "templates",
    "ecr"
  ],
  "language": "crystal"
}
-->

## ECR Templates
Embedded Crystal templating
```crystal
require "ecr"

# Inline template
name = "Alice"
ECR.embed "template.ecr", io
# template.ecr: Hello, <%= name %>!

# Template in string
template = <<-ECR
  <h1>Welcome, <%= name %>!</h1>
  <p>Age: <%= age %></p>
ECR

# Define method with template
class Page
  ECR.def_to_s "views/page.ecr"
end

# With loops
users = ["Alice", "Bob", "Charlie"]
ECR.embed "users.ecr", io
# users.ecr:
# <ul>
# <% users.each do |user| %>
#   <li><%= user %></li>
# <% end %>
# </ul>

# Conditionals
show_admin = true
ECR.embed "dashboard.ecr", io
# dashboard.ecr:
# <% if show_admin %>
#   <a href="/admin">Admin Panel</a>
# <% end %>

# Render to string
output = ECR.render "template.ecr"

# Class-based templates
class UserPage
  def initialize(@name : String, @age : Int32)
  end

  ECR.def_to_s "#{__DIR__}/templates/user.ecr"
end

page = UserPage.new("Alice", 30)
puts page.to_s
```
