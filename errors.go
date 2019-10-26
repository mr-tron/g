package g

type Error string

func (e Error) Error() string {
	return string(e)
}

const EmptySlice Error = "empty slice"
