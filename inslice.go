package g

import "time"

func IntInSlice(v int, slice []int) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Int8InSlice(v int8, slice []int8) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Int16InSlice(v int16, slice []int16) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Int32InSlice(v int32, slice []int32) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Int64InSlice(v int64, slice []int64) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func RuneInSlice(v rune, slice []rune) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func UintInSlice(v uint, slice []uint) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Uint8InSlice(v uint8, slice []uint8) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Uint16InSlice(v uint16, slice []uint16) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Uint32InSlice(v uint32, slice []uint32) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Uint64InSlice(v uint64, slice []uint64) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func ByteInSlice(v byte, slice []byte) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Float32InSlice(v float32, slice []float32) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Float64InSlice(v float64, slice []float64) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Complex64InSlice(v complex64, slice []complex64) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func Complex128InSlice(v complex128, slice []complex128) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func BoolInSlice(v bool, slice []bool) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func StringInSlice(v string, slice []string) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

func TimeInSlice(v time.Time, slice []time.Time) bool {
	for i := range slice {
		if slice[i].Equal(v) {
			return true
		}
	}
	return false
}
