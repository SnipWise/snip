<!-- METADATA
{
  "title": "Python Collections And Data Structures",
  "tags": [
    "python",
    "structs",
    "io",
    "collections"
  ],
  "language": "python"
}
-->

## Collections and Data Structures
Advanced data structures
```python
from collections import defaultdict, Counter, deque, namedtuple
import heapq

# defaultdict
dd = defaultdict(list)
dd['fruits'].append('apple')
dd['fruits'].append('banana')
print(f"Default dict: {dict(dd)}")

# Counter
text = "hello world"
char_count = Counter(text)
print(f"Character count: {char_count}")
print(f"Most common: {char_count.most_common(3)}")

# deque (double-ended queue)
queue = deque([1, 2, 3])
queue.appendleft(0)
queue.append(4)
left = queue.popleft()
right = queue.pop()
print(f"Queue: {queue}, popped: {left}, {right}")

# namedtuple
Person = namedtuple('Person', ['name', 'age', 'city'])
person = Person('Alice', 30, 'New York')
print(f"Person: {person.name}, {person.age}")

# heapq (priority queue)
heap = [3, 1, 4, 1, 5, 9, 2, 6]
heapq.heapify(heap)
smallest = heapq.heappop(heap)
heapq.heappush(heap, 0)
print(f"Smallest: {smallest}, heap: {heap}")
```