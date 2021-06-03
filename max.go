package g

import "time"

func MaxIntInSlice(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxIntWithDefault(d int, slice ...int) int {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxInt8InSlice(slice []int8) (int8, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxInt8WithDefault(d int8, slice ...int8) int8 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxInt16InSlice(slice []int16) (int16, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxInt16WithDefault(d int16, slice ...int16) int16 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxInt32InSlice(slice []int32) (int32, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxInt32WithDefault(d int32, slice ...int32) int32 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxInt64InSlice(slice []int64) (int64, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxInt64WithDefault(d int64, slice ...int64) int64 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxUintInSlice(slice []uint) (uint, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxUintWithDefault(d uint, slice ...uint) uint {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxUint8InSlice(slice []uint8) (uint8, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxUint8WithDefault(d uint8, slice ...uint8) uint8 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxUint16InSlice(slice []uint16) (uint16, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxUint16WithDefault(d uint16, slice ...uint16) uint16 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxUint32InSlice(slice []uint32) (uint32, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxUint32WithDefault(d uint32, slice ...uint32) uint32 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxUint64InSlice(slice []uint64) (uint64, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxUint64WithDefault(d uint64, slice ...uint64) uint64 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxFloat32InSlice(slice []float32) (float32, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxFloat32WithDefault(d float32, slice ...float32) float32 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxFloat64InSlice(slice []float64) (float64, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxFloat64WithDefault(d float64, slice ...float64) float64 {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func MaxTimeWithDefault(d time.Time, slice ...time.Time) time.Time {
	max := d
	if len(slice) != 0 {
		max = slice[0]
	}
	for i := range slice {
		if slice[i].After(max) {
			max = slice[i]
		}
	}
	return max
}

func MaxTimeInSlice(slice []time.Time) (time.Time, error) {
	if len(slice) == 0 {
		return time.Time{}, EmptySlice
	}
	max := slice[0]
	for i := range slice {
		if slice[i].After(max) {
			max = slice[i]
		}
	}
	return max, nil
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}
func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}
func MaxUint8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}
func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}
func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
func MaxTime(a, b time.Time) time.Time {
	if a.After(b) {
		return a
	}
	return b
}
