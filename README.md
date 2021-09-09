# pow

Just a small experimental scratchpad to test out some performance theories in the classic Closest Power of Two problem.

## The Problem

One of the classic leetcode problems of this day and age is to calculate the closest power of two to a given input integer. The generally accepted "correct solution" is to do some bit shifting magic and calculate the closest power by shifting bits until your binary is a one followed by some zeroes, which tells you that the number is a power of two. While this is a good solution, it's also not an easy one for people who don't read binary or memorize how all bitwise operations work (like myself), so I decided to benchmark doing the bitwise operations vs just simply bruteforcing the solution, by looping until we get the correct number (which is my way of solving the problem).

## Results

The results were pretty interesting. I assumed that the bitwise magic would be way faster than the bruteforce solution, but from what I can tell, the difference is pretty minimal, approximately 1ns!

```
goos: darwin
goarch: amd64
pkg: github.com/probably-not/pow
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PowerOfTwoBitwise-16          161628186                7.228 ns/op           0 B/op          0 allocs/op
Benchmark_PowerOfTwoBruteForce-16       142655391                8.412 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/probably-not/pow     4.597s
```