<!-- METADATA
{
  "title": "Crystal Lucky Framework Basics",
  "tags": [
    "crystal",
    "web",
    "lucky"
  ],
  "language": "crystal"
}
-->

## Lucky Framework Basics
Type-safe web framework concepts
```crystal
# Lucky Action (controller)
class Users::Index < BrowserAction
  get "/users" do
    users = UserQuery.new
    html IndexPage, users: users
  end
end

# Lucky Page (view)
class Users::IndexPage < MainLayout
  needs users : UserQuery

  def content
    h1 "Users"
    render_users
  end

  private def render_users
    ul do
      @users.each do |user|
        li user.name
      end
    end
  end
end

# Lucky Model
class User < BaseModel
  table do
    column name : String
    column email : String
    column age : Int32
  end
end

# Query
users = UserQuery.new.age.gt(18)
user = UserQuery.find(1)

# Forms
class UserForm < User::SaveOperation
  permit_columns name, email, age

  before_save do
    validate_required name, email
  end
end

# Route
Lucky::Router.draw do
  root Home::Index
  get "/users", Users::Index
end
```
