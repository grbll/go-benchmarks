package multiplicationbytwovsbitshiftonint32

import "testing"

func Benchmark_Shift(b *testing.B) {
	var num int32
	b.N = 100000000000
	for i := 0; i < b.N; i++ {
		num = int32(1)
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
		num <<= 1
	}
}
func Benchmark_MultiplicationByTwo(b *testing.B) {
	var num int32
	b.N = 100000000000

	for i := 0; i < b.N; i++ {
		num = int32(1)
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
		num *= 2
	}
}
