# Stackgo

**Stackgo** is a *slice-based* implementation of a simple stack in Go.
It uses a pre-alloc strategy which adds little memory overhead to the stack allocation
but makes push operations about ~+3x times faster than with a classic Stack implementation.

## Usage

Using it is pretty straightforward

```go
package main

import (
  "github.com/alediaferia/stackgo"
  "fmt"
)

func main() {
  stack := stackgo.NewStack()

  // Stack supports any type
  // so we just push whatever
  // we want here
  stack.Push(75)
  stack.Push(124)
  stack.Push("Hello World!")

  for stack.Size() > 0 {
    fmt.Printf("Just popped %v\n", stack.Pop())
  }

}
```

## Performance

This tiny package comes with a couple of benchmarks that compare the slice-based
implementation with a *traditional* one.

I'm pasting here the benchmark results of how it performs on my machine, an Intel Core i5,
2.3 GHz, 8 GB Ram:

```
Benchmark_PushDefaultStack	20000000	        90.6 ns/op	      24 B/op	       1 allocs/op
--- BENCH: Benchmark_PushDefaultStack
	stackgo_test.go:104: Testing push speed of 1 integer values on the default Stack implementation
	stackgo_test.go:104: Testing push speed of 100 integer values on the default Stack implementation
	stackgo_test.go:104: Testing push speed of 10000 integer values on the default Stack implementation
	stackgo_test.go:104: Testing push speed of 1000000 integer values on the default Stack implementation
	stackgo_test.go:104: Testing push speed of 20000000 integer values on the default Stack implementation
Benchmark_PushAltStack	 5000000	       275 ns/op	      40 B/op	       2 allocs/op
--- BENCH: Benchmark_PushAltStack
	stackgo_test.go:115: Testing push speed of 1 integer values on the alternate Stack implementation
	stackgo_test.go:115: Testing push speed of 100 integer values on the alternate Stack implementation
	stackgo_test.go:115: Testing push speed of 10000 integer values on the alternate Stack implementation
	stackgo_test.go:115: Testing push speed of 1000000 integer values on the alternate Stack implementation
	stackgo_test.go:115: Testing push speed of 5000000 integer values on the alternate Stack implementation
Benchmark_PopDefaultStack	2000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
--- BENCH: Benchmark_PopDefaultStack
	stackgo_test.go:125: Testing pop speed of 1 integer values on the default Stack implementation
	stackgo_test.go:125: Testing pop speed of 100 integer values on the default Stack implementation
	stackgo_test.go:125: Testing pop speed of 10000 integer values on the default Stack implementation
	stackgo_test.go:125: Testing pop speed of 1000000 integer values on the default Stack implementation
	stackgo_test.go:125: Testing pop speed of 100000000 integer values on the default Stack implementation
	stackgo_test.go:125: Testing pop speed of 2000000000 integer values on the default Stack implementation
Benchmark_PopAltStack	200000000	        61.3 ns/op	       0 B/op	       0 allocs/op
--- BENCH: Benchmark_PopAltStack
	stackgo_test.go:138: Testing pop speed of 1 integer values on the alternate Stack implementation
	stackgo_test.go:138: Testing pop speed of 100 integer values on the alternate Stack implementation
	stackgo_test.go:138: Testing pop speed of 10000 integer values on the alternate Stack implementation
	stackgo_test.go:138: Testing pop speed of 1000000 integer values on the alternate Stack implementation
	stackgo_test.go:138: Testing pop speed of 100000000 integer values on the alternate Stack implementation
	stackgo_test.go:138: Testing pop speed of 200000000 integer values on the alternate Stack implementation
ok  	github.com/alediaferia/stackgo	313.428s
```

## Contribute
I'd really appreciate contributions, otherwise I wouldn't have made this open :p
Also, if you have suggestions on how to make this perform even faster I'd be really happy to hear about them.

## License
This code is released under the MIT License term, included in this project tree.
Copyright © 2015, Alessandro Diaferia <alediaferia@gmail.com>
