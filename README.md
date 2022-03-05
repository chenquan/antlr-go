# antlr-go

## [calculate](/calculate)

A simple calculator implementation.

### How to use

```go
package main

import (
	"fmt"
	"github.com/chenquan/antrl-go/calculate"
)

func main() {
	s := calculate.Calculate(`1+1`)
	fmt.Println(s) // output 2

	s = calculate.Calculate(`a=1
b=3
a+b
`)
	fmt.Println(s) // output 4

}
```