## Data Processing with Pandas
Basic data manipulation with pandas
```python
import pandas as pd

# Create sample data
data = {
    'name': ['Alice', 'Bob', 'Charlie', 'Diana'],
    'age': [25, 30, 35, 28],
    'salary': [50000, 60000, 70000, 55000]
}

df = pd.DataFrame(data)
print(df)

# Filtering
young = df[df['age'] < 30]
high_earners = df[df['salary'] > 55000]

# Basic stats
avg_age = df['age'].mean()
total_salary = df['salary'].sum()

print(f"Average age: {avg_age}")
print(f"Total salary: {total_salary}")

# Add new column
df['salary_k'] = df['salary'] / 1000
df['category'] = df['age'].apply(lambda x: 'Young' if x < 30 else 'Senior')

# Sort and display
df_sorted = df.sort_values('age')
print(f"\nSorted by age:\n{df_sorted[['name', 'age', 'category']]}")
```