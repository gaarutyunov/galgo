package numbers

func GeneratePrimes(max int) (res []int) {
	for i, f := range SieveOfEratosthenes(max) {
		if f {
			res = append(res, i)
		}
	}

	return
}

func SieveOfEratosthenes(max int) (flags []bool) {
	flags = make([]bool, max+1)

	for i := range flags {
		flags[i] = true
	}

	flags[0] = false
	flags[1] = false
	prime := 2

	sqrt := Sqrt(max)

	for prime <= sqrt {
		crossOffMultiples(flags, prime)
		prime = getNextPrime(flags, prime)
	}

	return
}

func getNextPrime(flags []bool, prime int) int {
	next := prime + 1
	for next < len(flags) && !flags[next] {
		next++
	}
	return next
}

func crossOffMultiples(flags []bool, n int) {
	for i := n * n; i < len(flags); i += n {
		flags[i] = false
	}
}
