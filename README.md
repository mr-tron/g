# G - general framework for go

All built in golang types:
```
bool
int8, int16, int32, int64, int
uint8, uint16, uint32, uint64, uint
float32, float64
complex64, complex128
rune, byte # aliases for int32 and uint8
```

All supported functions:
```
XXXInSlice(value T, slice []T) bool      # all types
MaxXXX(slice []T) (T, error)             # intX, uintX, floatX types
MaxWithDeafult(slice []T, default T) T   # intX, uintX, floatX types
MinXXX(slice []T) (T, error)             # intX, uintX, floatX types
MinWithDeafult(slice []T, default T) T   # intX, uintX, floatX types
```
