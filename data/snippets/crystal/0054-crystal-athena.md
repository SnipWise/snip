<!-- METADATA
{
  "title": "Crystal Athena Framework",
  "tags": [
    "crystal",
    "web",
    "athena"
  ],
  "language": "crystal"
}
-->

## Athena Framework
Component-based web framework
```crystal
require "athena"

# Controller
@[ADI::Register]
class ArticleController < ART::Controller

  @[ARTA::Get("/articles")]
  def index : Array(String)
    ["Article 1", "Article 2", "Article 3"]
  end

  @[ARTA::Get("/articles/:id")]
  def show(id : Int32) : String
    "Article #{id}"
  end

  @[ARTA::Post("/articles")]
  def create(
    @[ARTA::RequestBody] body : ArticleCreate
  ) : ArticleCreate
    body
  end

  @[ARTA::QueryParam("page")]
  @[ARTA::Get("/paginated")]
  def paginated(page : Int32 = 1) : String
    "Page #{page}"
  end
end

# DTO (Data Transfer Object)
struct ArticleCreate
  include JSON::Serializable

  property title : String
  property content : String
end

# Run server
ART.run

# Custom service
@[ADI::Register]
class ArticleService
  def find_all : Array(String)
    ["Article 1", "Article 2"]
  end
end
```
