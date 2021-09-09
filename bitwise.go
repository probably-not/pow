package pow

func NextPowerOfTwo(target int64) int64 {
	target--

	target |= target >> 1
	target |= target >> 2
	target |= target >> 4
	target |= target >> 8
	target |= target >> 16
	target |= target >> 32

	target++
	return target
}

func ClosestPowerOfTwoBitwise(input int64) int64 {
	// We can't calculate for negative numbers
	if input < 0 {
		panic("input cannot be negative")
	}

	// Less than or equal to 1 should be 1 (2^0=1)
	if input <= 1 {
		return 1
	}

	// Exactly 2 should be 2 (2^1)
	if input == 2 {
		return 2
	}

	nextPowerOfTwo := NextPowerOfTwo(input)

	if prev := nextPowerOfTwo / 2; (input - prev) < (nextPowerOfTwo - input) {
		return prev
	}

	return nextPowerOfTwo
}
