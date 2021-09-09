package pow

import (
	"math/rand"
	"testing"
)

var x int64

// Benchmark_PowerOfTwo-16    	173414430	         7.488 ns/op	       0 B/op	       0 allocs/op
func Benchmark_PowerOfTwoBitwise(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x = ClosestPowerOfTwoBitwiseFor(int64(i))
	}
}

// Benchmark_PowerOfTwoBruteForce-16    	134810326	         8.502 ns/op	       0 B/op	       0 allocs/op
func Benchmark_PowerOfTwoBruteForce(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x = ClosestPowerOfTwoBruteForce(int64(i))
	}
}

func Test_PowerOfTwo(t *testing.T) {
	for i := 0; i < 1000; i++ {
		input := rand.Intn(1000000000)
		bruteForce := ClosestPowerOfTwoBruteForce(int64(input))
		bitwisefor := ClosestPowerOfTwoBitwiseFor(int64(input))
		if bruteForce != bitwisefor {
			t.Errorf("bruteforce and bitwisefor are not equal for input %d. bruteforce got=%d, bitwisefor got=%d.", input, bruteForce, bitwisefor)
		}
	}
}
