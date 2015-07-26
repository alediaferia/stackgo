# Stackgo [![Build Status](https://secure.travis-ci.org/alediaferia/stackgo.svg)](http://travis-ci.org/alediaferia/stackgo) [![Coverage Status](https://coveralls.io/repos/alediaferia/stackgo/badge.svg?branch=master)](https://coveralls.io/r/alediaferia/stackgo?branch=master)

**Stackgo** is a *slice-based* implementation of a simple stack in Go.
It uses a pre-alloc pagination strategy which adds little memory overhead to the stack allocation
but makes push operations faster than with a classic Stack implementation.

Please **NOTE** that the current implementation is NOT thread-safe.

- [Usage](#Usage)
- [Performance](#Performance)
- [Contribute](#Contribute)
- [License](#License)

## Getting started

### Import
You can either import this package directly:

```go
import "github.com/alediaferia/stackgo"
```

or through [gopkg.in](http://gopkg.in)

```go
import "gopkg.in/alediaferia/stackgo.v1"
```

Currently only **version 1** has been released.

### Usage

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
Check the implementation details [here](docs/IMPLEMENTATION.md).

## Contribute
I'd really appreciate contributions, otherwise I wouldn't have made this open :smiley:.
Also, if you have suggestions on how to make this perform even faster I'd be really happy to hear about them.

## License
This code is released under the MIT License term, included in this project tree.
Copyright © 2015, Alessandro Diaferia <alediaferia@gmail.com>
