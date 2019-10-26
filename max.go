package g

func MaxInt(slice []int) (int, error) {
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

func MaxInt8(slice []int8) (int8, error) {
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

func MaxInt16(slice []int16) (int16, error) {
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

func MaxInt32(slice []int32) (int32, error) {
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

func MaxInt64(slice []int64) (int64, error) {
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

func MaxUint(slice []uint) (uint, error) {
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

func MaxUint8(slice []uint8) (uint8, error) {
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

func MaxUint16(slice []uint16) (uint16, error) {
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

func MaxUint32(slice []uint32) (uint32, error) {
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

func MaxUint64(slice []uint64) (uint64, error) {
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

func MaxFloat32(slice []float32) (float32, error) {
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

func MaxFloat64(slice []float64) (float64, error) {
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
