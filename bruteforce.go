package pow

func ClosestPowerOfTwoBruteForce(input int64) int64 {
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

	prev := int64(2)
	for {
		nextPowerOfTwo := prev * 2

		if nextPowerOfTwo >= input {
			if (input - prev) < (nextPowerOfTwo - input) {
				return prev
			}

			return nextPowerOfTwo
		}

		prev = nextPowerOfTwo
	}
}
