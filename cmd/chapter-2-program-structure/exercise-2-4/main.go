package pop_count_v2

func PopCountBitShift(x uint64) int {
	var count int

	for i := 0; i < 64; i++ {

		/*
			Bitwise AND is used to check whether an particular bit is set or clear
			Left-shift is equivalent to multiply by 2^n
			To find how many bits are set on given number, we do a Bitwise AND between number and (1 left-shift *number of bit*)
			Example:
				13 = (1101)_2; number of bits = 64

				i = 0:
					13 [1101] & (1 << i) [1 * 2^0 = 1] = 1101 & 0001 = 0001 = 1 (Bit is set)
				i = 1:
					13 [1101] & (1 << i) [1 * 2^1 = 2] = 1101 & 0010 = 0000 = 0 (Bit is clear)
				i = 2:
					13 [1101] & (1 << i) [1 * 2^2 = 4] = 1101 & 0100 = 0100 = 4 (Bit is set)
				i = 3:
					13 [1101] & (1 << i) [1 * 2^3 = 8] = 1101 & 1000 = 1000 = 8 (Bit is set)
				i = 4:
					13 [1101] & (1 << i) [1 * 2^4 = 8] = 01101 & 10000 = 00000 = 0 (Bit is clear)
				...
				...
		*/

		if x&(1<<i) != 0 {
			count++
		}

	}

	return count
}
