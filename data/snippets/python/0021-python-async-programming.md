<!-- METADATA
{
  "title": "Python Async Programming",
  "tags": [
    "python",
    "async"
  ],
  "language": "python"
}
-->

## Async Programming
Asynchronous programming with asyncio
```python
import asyncio
import time

async def say_hello(name, delay):
    await asyncio.sleep(delay)
    return f"Hello, {name}!"

async def fetch_data(data_id, delay):
    print(f"Fetching data {data_id}...")
    await asyncio.sleep(delay)
    return f"Data {data_id} fetched"

async def main():
    # Run tasks concurrently
    greet_tasks = [
        say_hello("Alice", 1),
        say_hello("Bob", 2),
        say_hello("Charlie", 1)
    ]
    
    # Sequential vs concurrent execution
    start = time.time()
    results = await asyncio.gather(*greet_tasks)
    concurrent_time = time.time() - start
    
    for result in results:
        print(result)
    
    print(f"Concurrent execution: {concurrent_time:.2f}s")
    
    # Fetch data concurrently
    data_tasks = [fetch_data(i, 0.5) for i in range(3)]
    data_results = await asyncio.gather(*data_tasks)
    print("Data results:", data_results)

if __name__ == "__main__":
    asyncio.run(main())
```