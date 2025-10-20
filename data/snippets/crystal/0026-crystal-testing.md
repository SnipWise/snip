<!-- METADATA
{
  "title": "Crystal Testing",
  "tags": [
    "crystal",
    "testing",
    "spec"
  ],
  "language": "crystal"
}
-->

## Testing
Unit testing with Spec
```crystal
require "spec"

# Simple test
describe "Math" do
  it "adds two numbers" do
    result = 2 + 2
    result.should eq 4
  end

  it "multiplies numbers" do
    (3 * 4).should eq 12
  end
end

# Testing a class
class Calculator
  def add(a, b)
    a + b
  end

  def divide(a, b)
    raise "Division by zero" if b == 0
    a / b
  end
end

describe Calculator do
  calc = Calculator.new

  it "adds numbers" do
    calc.add(5, 3).should eq 8
  end

  it "raises on division by zero" do
    expect_raises(Exception, "Division by zero") do
      calc.divide(10, 0)
    end
  end
end

# Run with: crystal spec
```
