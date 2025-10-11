## Flask Web Server
Creating a simple web server with Flask
```python
from flask import Flask, jsonify, request

app = Flask(__name__)

users = ["Alice", "Bob", "Charlie"]

@app.route("/")
def home():
    return "Hello, Flask!"

@app.route("/users", methods=["GET"])
def get_users():
    return jsonify(users)

@app.route("/users", methods=["POST"])
def create_user():
    data = request.json
    name = data.get("name")
    if name:
        users.append(name)
        return jsonify({"message": f"User {name} added"}), 201
    return jsonify({"error": "Name required"}), 400

if __name__ == "__main__":
    app.run(debug=True, port=5000)
```