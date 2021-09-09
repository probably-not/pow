# pow

Just a small experimental scratchpad to test out some performance theories in the classic Closest Power of Two problem.

## The Problem

One of the classic leetcode problems of this day and age is to calculate the closest power of two to a given input integer. The generally accepted "correct solution" is to do some bit shifting magic and calculate the closest power by shifting bits until your binary is a one followed by some zeroes, which tells you that the number is a power of two. While this is a good solution, it's also not an easy one for people who don't read binary or memorize how all bitwise operations work (like myself), so I decided to benchmark doing the bitwise operations vs just simply bruteforcing the solution, by looping until we get the correct number (which is my way of solving the problem).

## Results

The results were pretty interesting.

I assumed that the bitwise magic would be way faster than the bruteforce solution, but it looks like this only holds true when we have hardcoded our set of bit shifting operations.

Since this would theoretically do extra shift attempts when we don't need them, I wanted to try out adding a for loop so we could cut off the number of shifts, but apparently the overhead of a for loop adds too much of a performance hit.

On the other hand, from what I can tell, the difference between the Bitwise version with the For loop and the Bruteforce is pretty minimal, (approximately 1ns) which was a pretty surprising revelation! I guess that when bruteforcing with raw math and a for loop vs. bitwise operations and a for loop, the real performance penalty is the for loop itself. From what I could tell, the bitwise operations and math operations have essentially the same overhead (1ns per op), so the for loop adds the major hit.

```
goos: darwin
goarch: amd64
pkg: github.com/probably-not/pow
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_PowerOfTwo_Bitwise-16         346069328                3.521 ns/op           0 B/op          0 allocs/op
Benchmark_PowerOfTwo_BitwiseFor-16      153441846                7.544 ns/op           0 B/op          0 allocs/op
Benchmark_PowerOfTwo_BruteForce-16      142918903                8.557 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/probably-not/pow     6.008s
```