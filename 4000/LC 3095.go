func minimumSubarrayLength(nums []int, k int) int {
     n := len(nums)
	minLength := n + 1 
	left := 0
	currentOr := 0

	for right := 0; right < n; right++ {
		currentOr |= nums[right] 


		for currentOr >= k {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}

			if left < right {
				newOr := 0 
				for i := left + 1; i <= right; i++ {
					newOr |= nums[i]
				}
				currentOr = newOr
				left++
			} else {
				break 
			}
		}
	}

	if minLength <= n {
		return minLength
	}
	return -1 
    
}
