package numbers

const (
	IntSize  = 32 << (^uint(0) >> 63)
	MaxInt   = 1<<(IntSize-1) - 1
	MinInt   = -1 << (IntSize - 1)
	MinInt64 = -1 << 63
	MaxInt64 = 1<<63 - 1
	MinInt32 = -1 << 31
	MaxInt32 = 1<<31 - 1
	MinInt16 = -1 << 15
	MaxInt16 = 1<<15 - 1
	MinInt8  = -1 << 7
	MaxInt8  = 1<<7 - 1
)

func FromString(s string) (res int) {
	for i, j := len(s)-1, 1; i >= 0; i, j = i-1, j*10 {
		if s[i] == '-' {
			res = -res
		} else {
			res = res + int(s[i]-'0')*j
		}
	}

	return
}

func IsDigit(char byte) bool {
	switch char {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

func Abs(x int) int {
	return (x + (x >> (IntSize - 1))) ^ (x >> (IntSize - 1))
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
