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

func MaxIntWithDefault(slice []int, d int) int {
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

func MaxInt8WithDefault(slice []int8, d int8) int8 {
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

func MaxInt16WithDefault(slice []int16, d int16) int16 {
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

func MaxInt32WithDefault(slice []int32, d int32) int32 {
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

func MaxInt64WithDefault(slice []int64, d int64) int64 {
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

func MaxUintWithDefault(slice []uint, d uint) uint {
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

func MaxUint8WithDefault(slice []uint8, d uint8) uint8 {
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

func MaxUint16WithDefault(slice []uint16, d uint16) uint16 {
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

func MaxUint32WithDefault(slice []uint32, d uint32) uint32 {
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

func MaxUint64WithDefault(slice []uint64, d uint64) uint64 {
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

func MaxFloat32WithDefault(slice []float32, d float32) float32 {
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

func MaxFloat64WithDefault(slice []float64, d float64) float64 {
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

func MaxTimeWithDefault(slice []time.Time, t time.Time) time.Time {
	max := t
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
