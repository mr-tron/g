package g

func UniqBool(input []bool) []bool {
	output := make([]bool, 0, 2)
	counter := 0
	for i := range input {
		if counter == 3 {
			return output
		}
		if input[i] {
			counter = counter | 1
		} else {
			counter = counter | 2
		}
		if !BoolInSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqString(input []string) []string {
	output := make([]string, 0, len(input))
	for i := range input {
		if !StringInSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqInt(input []int) []int {
	output := make([]int, 0, len(input))
	for i := range input {
		if !IntInSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}
func UniqInt8(input []int8) []int8 {
	output := make([]int8, 0, len(input))
	for i := range input {
		if !Int8InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqInt16(input []int16) []int16 {
	output := make([]int16, 0, len(input))
	for i := range input {
		if !Int16InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqInt32(input []int32) []int32 {
	output := make([]int32, 0, len(input))
	for i := range input {
		if !Int32InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqInt64(input []int64) []int64 {
	output := make([]int64, 0, len(input))
	for i := range input {
		if !Int64InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqUint(input []uint) []uint {
	output := make([]uint, 0, len(input))
	for i := range input {
		if !UintInSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqUint8(input []uint8) []uint8 {
	output := make([]uint8, 0, len(input))
	for i := range input {
		if !Uint8InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqUint16(input []uint16) []uint16 {
	output := make([]uint16, 0, len(input))
	for i := range input {
		if !Uint16InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqUint32(input []uint32) []uint32 {
	output := make([]uint32, 0, len(input))
	for i := range input {
		if !Uint32InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqUint64(input []uint64) []uint64 {
	output := make([]uint64, 0, len(input))
	for i := range input {
		if !Uint64InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqFloat32(input []float32) []float32 {
	output := make([]float32, 0, len(input))
	for i := range input {
		if !Float32InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}

func UniqFloat64(input []float64) []float64 {
	output := make([]float64, 0, len(input))
	for i := range input {
		if !Float64InSlice(input[i], output) {
			output = append(output, input[i])
		}
	}
	return output
}
