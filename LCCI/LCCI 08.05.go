func multiply(a int, b int) int {
    res := 0
	for b > 0 {
		if b & 1 == 1 {
			res += a
		}
		a <<= 1
		b >>= 1
	}

	return res
}