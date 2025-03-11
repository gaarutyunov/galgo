package numbers

func Sqrt(n int) int {
	x := n

	for {
		y := (x + n/x) / 2
		if y >= x {
			return x
		}
		x = y
	}
}
