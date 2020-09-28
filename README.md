# G - general framework for go

[![GoDoc](https://godoc.org/github.com/mr-tron/g?status.svg)](https://godoc.org/github.com/mr-tron/g)  [![Go Report Card](https://goreportcard.com/badge/github.com/mr-tron/g)](https://goreportcard.com/report/github.com/mr-tron/g)
[![Used By](https://sourcegraph.com/github.com/mr-tron/g/-/badge.svg)](https://sourcegraph.com/github.com/mr-tron/g?badge)


All built-in golang types:
```
bool
string
int8, int16, int32, int64, int
uint8, uint16, uint32, uint64, uint
float32, float64
complex64, complex128
rune, byte # aliases for int32 and uint8
```

Also library supports `time.Time` type

All supported functions:
```
XXXInSlice(value T, slice []T) bool      # all types
MaxXXX(a, b T) T                         # intX, uintX, floatX, time.Time types
MaxXXXInSlice(slice []T) (T, error)      # intX, uintX, floatX, time.Time types
MaxWithDeafult(slice []T, default T) T   # intX, uintX, floatX, time.Time types
MinXXX(a, b T) T                         # intX, uintX, floatX, time.Time types
MinXXXInSlice(slice []T) (T, error)      # intX, uintX, floatX, time.Time types
MinWithDeafult(slice []T, default T) T   # intX, uintX, floatX, time.Time types
UniqXXX(slice []T) []T                   # all types
Pxxx(v T) *T                             # all types
```

## Example

```go
package main
import (
    "github.com/mr-tron/g"  
    "fmt"
)

func main() {
    prices := []uint64{15999, 12000, 13490, 99999}
    outTransactions := []uint64{500, 2900, 12000, 980, 500}
    lowestPrice, err  := g.MinUint64InSlice(prices)
    if err != nil {
        fmt.Println("Nobody sells giant dildos. Sorry")
        return
    }
    fmt.Printf("- Did we buy giant dildo on lowest price?\n- %v", g.Uint64InSlice(lowestPrice, outTransactions))
}
```