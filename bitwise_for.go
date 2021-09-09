package pow

func NextPowerOfTwoFor(target int64) int64 {
	target--

	for i := 1; target>>i != 0; i *= 2 {
		target |= target >> i
	}

	target++
	return target
}

func ClosestPowerOfTwoBitwiseFor(input int64) int64 {
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

	nextPowerOfTwo := NextPowerOfTwoFor(input)

	if prev := nextPowerOfTwo / 2; (input - prev) < (nextPowerOfTwo - input) {
		return prev
	}

	return nextPowerOfTwo
}
