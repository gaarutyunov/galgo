package functional

func Pipe2[T1 any, T2 any, T3 any](f1 func(T1) T2, f2 func(T2) T3) func(T1) T3 {
	return func(t1 T1) T3 {
		return f2(f1(t1))
	}
}

func Pipe3[T1 any, T2 any, T3 any, T4 any](f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4) func(T1) T4 {
	return func(t1 T1) T4 {
		return f3(Pipe2(f1, f2)(t1))
	}
}

func Pipe4[T1 any, T2 any, T3 any, T4 any, T5 any](f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4, f4 func(T4) T5) func(T1) T5 {
	return func(t1 T1) T5 {
		return f4(Pipe3(f1, f2, f3)(t1))
	}
}

func Pipe5[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4, f4 func(T4) T5, f5 func(T5) T6) func(T1) T6 {
	return func(t1 T1) T6 {
		return f5(Pipe4(f1, f2, f3, f4)(t1))
	}
}
