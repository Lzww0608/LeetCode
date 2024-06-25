func maximum(a int, b int) int {
	diff := int64(a) - int64(b)

	sign := (diff >> 63) & 1

	return int((1-sign)*int64(a) + sign*int64(b))
}