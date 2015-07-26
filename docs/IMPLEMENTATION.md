# Implementation
**Stackgo** takes advantage of a simple page-based implementation that pre-allocates memory in advance using a default page size of 4096 bytes.

Although there is a rare pre-allocation overhead, on average the *push* time complexity is O(1).

Memory-shrinking is supported when completely freeing a page upon *pop*. This means that there may be an additional overhead when popping, but on average *pop* time complexity is still O(1).

Benchmarks
----------
Here the results of benchmarking **stackgo**

 Implementation | Processor             | Operation | ns/op | B/op | allocs/op
----------------|-----------------------|-----------|-------|------|----------
  **stackgo**   | 2.3 GHz Intel Core i5 |   PUSH    |  133  |  24  |    1
  **classic**   | 2.3 GHz Intel Core i5 |   PUSH    |  255  |  40  |    2
  **stackgo**   | 2.3 GHz Intel Core i5 |   POP     |  15.6 |   0  |    0
  **classic**   | 2.3 GHz Intel Core i5 |   POP     |  12.6 |   0  |    0
  
As you can see, **stackgo** outperforms on push operation, giving roughly the same performances on pop.