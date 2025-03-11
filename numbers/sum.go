package numbers

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Sum[T Number](nums []T) (ans T) {
	for _, num := range nums {
		ans += num
	}

	return
}
