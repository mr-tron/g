package g

func MinInt(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinIntWithDefault(slice []int, d int) int {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinInt8(slice []int8) (int8, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinInt8WithDefault(slice []int8, d int8) int8 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinInt16(slice []int16) (int16, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinInt16WithDefault(slice []int16, d int16) int16 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinInt32(slice []int32) (int32, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinInt32WithDefault(slice []int32, d int32) int32 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinInt64(slice []int64) (int64, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinInt64WithDefault(slice []int64, d int64) int64 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinUint(slice []uint) (uint, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinUintWithDefault(slice []uint, d uint) uint {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinUint8(slice []uint8) (uint8, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinUint8WithDefault(slice []uint8, d uint8) uint8 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinUint16(slice []uint16) (uint16, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinUint16WithDefault(slice []uint16, d uint16) uint16 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinUint32(slice []uint32) (uint32, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinUint32WithDefault(slice []uint32, d uint32) uint32 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinUint64(slice []uint64) (uint64, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinUint64WithDefault(slice []uint64, d uint64) uint64 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinFloat32(slice []float32) (float32, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinFloat32WithDefault(slice []float32, d float32) float32 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}

func MinFloat64(slice []float64) (float64, error) {
	if len(slice) == 0 {
		return 0, EmptySlice
	}
	min := slice[0]
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min, nil
}

func MinFloat64WithDefault(slice []float64, d float64) float64 {
	min := d
	if len(slice) != 0 {
		min = slice[0]
	}
	for i := range slice {
		if slice[i] < min {
			min = slice[i]
		}
	}
	return min
}
