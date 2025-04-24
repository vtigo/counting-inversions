# Benchmark Results Analysis

The benchmark results reveal interesting performance characteristics of both algorithms:

## Quadratic Approach (O(n²))
- For small inputs (n=10), it's extremely fast (~28 ns/op) with no memory allocations
- Performance scales predictably with input size:
  - n=100: ~1,980 ns/op (about 70× slower than n=10)
  - n=1000: ~167,600 ns/op (about 85× slower than n=100)
- Shows consistent performance across all scenarios (sorted, reverse, random)
- Zero memory allocations across all test cases

## Divide and Conquer Approach (O(n log n))
- For small inputs (n=10), it's slower (~280 ns/op) than the quadratic approach
- Shows superior scaling with increasing input size:
  - n=100: ~4,300 ns/op (~15× slower than n=10)
  - n=1000: ~66,700 ns/op (~15× slower than n=100)
  - n=10000: ~825,000 ns/op (~12× slower than n=1000)
- Small variations in performance between scenarios, with sorted arrays performing slightly better
- Significant memory usage: allocates arrays during the merge steps (~9 allocations for n=10, ~999 for n=1000)

## Key Insights
1. **Crossover point**: The divide and conquer approach becomes faster than the quadratic approach between n=100 and n=1000
2. **Memory trade-off**: The divide and conquer approach uses significant memory (O(n log n) allocations) to achieve its speed advantage
3. **Scaling behavior**: The results clearly demonstrate the theoretical complexity:
   - Quadratic: increasing n by 10× makes it ~85× slower
   - Divide & Conquer: increasing n by 10× makes it only ~15× slower
4. **Algorithm selection guidance**: 
   - For n < ~500: Use the quadratic approach
   - For n > ~500: Use the divide and conquer approach
   - For memory-constrained environments: Consider the quadratic approach even for larger inputs

These results validate the theoretical time complexity of both algorithms and provide practical guidance for algorithm selection based on input size and memory constraints.
