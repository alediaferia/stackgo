# Stackgo

**Stackgo** is a *slice-based* implementation of a simple stack in Go.
It uses a pre-alloc pagination strategy which adds little memory overhead to the stack allocation
but makes push operations faster than with a classic Stack implementation.

Please **NOTE** that the current implementation is NOT thread-safe.

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

This implementation performs better than a traditional stack implementation.
What follows is the results of the benchmarks of both **stackgo** and default implementation of a stack, running on an Early 2011 Mac powered by Intel i5 2.3GHz with 8 GB of 1333MHz DDR3 RAM.

```bash
$ go test -bench=".*" github.com/alediaferia/stackgo
PASS
Benchmark_PushStackgo	10000000	       138 ns/op	      24 B/op	       1 allocs/op
--- BENCH: Benchmark_PushStackgo
	stackgo_test.go:147: Testing push speed of 1 integer values on the default Stack implementation
	stackgo_test.go:147: Testing push speed of 100 integer values on the default Stack implementation
	stackgo_test.go:147: Testing push speed of 10000 integer values on the default Stack implementation
	stackgo_test.go:147: Testing push speed of 1000000 integer values on the default Stack implementation
	stackgo_test.go:147: Testing push speed of 10000000 integer values on the default Stack implementation
Benchmark_PushStandardStack	 5000000	       274 ns/op	      40 B/op	       2 allocs/op
--- BENCH: Benchmark_PushStandardStack
	stackgo_test.go:158: Testing push speed of 1 integer values on the alternate Stack implementation
	stackgo_test.go:158: Testing push speed of 100 integer values on the alternate Stack implementation
	stackgo_test.go:158: Testing push speed of 10000 integer values on the alternate Stack implementation
	stackgo_test.go:158: Testing push speed of 1000000 integer values on the alternate Stack implementation
	stackgo_test.go:158: Testing push speed of 5000000 integer values on the alternate Stack implementation
Benchmark_PopStackgo	100000000	        12.0 ns/op	       0 B/op	       0 allocs/op
--- BENCH: Benchmark_PopStackgo
	stackgo_test.go:168: Testing pop speed of 1 integer values on the default Stack implementation
	stackgo_test.go:168: Testing pop speed of 100 integer values on the default Stack implementation
	stackgo_test.go:168: Testing pop speed of 10000 integer values on the default Stack implementation
	stackgo_test.go:168: Testing pop speed of 1000000 integer values on the default Stack implementation
	stackgo_test.go:168: Testing pop speed of 100000000 integer values on the default Stack implementation
Benchmark_PopStandardStack	200000000	        55.3 ns/op	       0 B/op	       0 allocs/op
--- BENCH: Benchmark_PopStandardStack
	stackgo_test.go:181: Testing pop speed of 1 integer values on the alternate Stack implementation
	stackgo_test.go:181: Testing pop speed of 100 integer values on the alternate Stack implementation
	stackgo_test.go:181: Testing pop speed of 10000 integer values on the alternate Stack implementation
	stackgo_test.go:181: Testing pop speed of 1000000 integer values on the alternate Stack implementation
	stackgo_test.go:181: Testing pop speed of 100000000 integer values on the alternate Stack implementation
	stackgo_test.go:181: Testing pop speed of 200000000 integer values on the alternate Stack implementation
ok  	github.com/alediaferia/stackgo	119.856s
```

## Contribute
I'd really appreciate contributions, otherwise I wouldn't have made this open :p
Also, if you have suggestions on how to make this perform even faster I'd be really happy to hear about them.

## License
This code is released under the MIT License term, included in this project tree.
Copyright © 2015, Alessandro Diaferia <alediaferia@gmail.com>
