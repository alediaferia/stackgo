# Stackgo

**Stackgo** is a *slice-based* implementation of a simple stack in Go.
It usually outperforms a traditional stack implementation when pushing elements.
It may outperform a traditional stack implementation when popping elements out
of the stack.

## Performance

This tiny package comes with a couple of benchmarks that compare the slice-based
implementation with a *traditional* one.

I'm pasting here the benchmark results of how it performs on my machine, an Intel Core i5,
2.3 GHz, 8 GB Ram:
```
Benchmark_PushDefaultStack	 5000000	       209 ns/op	      92 B/op	       1 allocs/op
--- BENCH: Benchmark_PushDefaultStack
	stack_test.go:74: Testing push speed of 1 integer values on the default Stack implementation
	stack_test.go:74: Testing push speed of 100 integer values on the default Stack implementation
	stack_test.go:74: Testing push speed of 10000 integer values on the default Stack implementation
	stack_test.go:74: Testing push speed of 1000000 integer values on the default Stack implementation
	stack_test.go:74: Testing push speed of 5000000 integer values on the default Stack implementation
Benchmark_PushAltStack	 5000000	       240 ns/op	      40 B/op	       2 allocs/op
--- BENCH: Benchmark_PushAltStack
	stack_test.go:85: Testing push speed of 1 integer values on the alternate Stack implementation
	stack_test.go:85: Testing push speed of 100 integer values on the alternate Stack implementation
	stack_test.go:85: Testing push speed of 10000 integer values on the alternate Stack implementation
	stack_test.go:85: Testing push speed of 1000000 integer values on the alternate Stack implementation
	stack_test.go:85: Testing push speed of 5000000 integer values on the alternate Stack implementation
Benchmark_PopDefaultStack	100000000	        22.7 ns/op	       0 B/op	       0 allocs/op
--- BENCH: Benchmark_PopDefaultStack
	stack_test.go:95: Testing pop speed of 1 integer values on the default Stack implementation
	stack_test.go:95: Testing pop speed of 100 integer values on the default Stack implementation
	stack_test.go:95: Testing pop speed of 10000 integer values on the default Stack implementation
	stack_test.go:95: Testing pop speed of 1000000 integer values on the default Stack implementation
	stack_test.go:95: Testing pop speed of 100000000 integer values on the default Stack implementation
Benchmark_PopAltStack	100000000	        21.4 ns/op	       0 B/op	       0 allocs/op
--- BENCH: Benchmark_PopAltStack
	stack_test.go:108: Testing pop speed of 1 integer values on the alternate Stack implementation
	stack_test.go:108: Testing pop speed of 100 integer values on the alternate Stack implementation
	stack_test.go:108: Testing pop speed of 10000 integer values on the alternate Stack implementation
	stack_test.go:108: Testing pop speed of 1000000 integer values on the alternate Stack implementation
	stack_test.go:108: Testing pop speed of 100000000 integer values on the alternate Stack implementation
ok  	github.com/alediaferia/stackgo	73.121s
```

## Contribute
I'd really appreciate contributions, otherwise I wouldn't have made this open :p
Also, if you have suggestions on how to make this perform even faster I'd be really happy to hear about them.

## License
This code is released under the MIT License term, included in this project tree.
Copyright © 2015, Alessandro Diaferia <alediaferia@gmail.com>
