func minimumSubarrayLength(nums []int, k int) int {
    n := len(nums)
	if n == 0 {
		return -1
	}

	bitCount := make(map[int]int)
	minLength := n + 1
	currentOr := 0
	left := 0

	for right := 0; right < n; right++ {
		for i := 0; i < 32; i++ {
			if (nums[right]>>i)&1 == 1 {
				bitCount[i]++
				currentOr |= (1 << i)
			}
		}

		for left <= right && currentOr >= k {
			if right - left + 1 < minLength {
				minLength = right - left + 1
			}
			for i := 0; i < 32; i++ {
				if (nums[left] >> i) & 1 == 1 {
					bitCount[i]--
					if bitCount[i] == 0 { 
						currentOr &= ^(1 << i)
					}
				}
			}
			left++
		}
	}

	if minLength <= n {
		return minLength
	}
	return -1
}