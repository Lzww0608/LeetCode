func minimumPairRemoval(nums []int) int {
    n := len(nums)
    for i := 0; i < n - 1; i++ {
        if check(nums) {
            return i 
        }

        j, mn := 0, math.MaxInt32 
        for k := range len(nums) - 1 {
            if nums[k] + nums[k + 1] < mn {
                mn = nums[k] + nums[k + 1]
                j = k 
            }
        }
        nums[j] += nums[j + 1]
        nums = append(nums[:j + 1], nums[j + 2:]...)
    }

    return n - 1
}

func check(a []int) bool {
    n := len(a)
    for i := range n - 1 {
        if a[i] > a[i + 1] {
            return false
        }
    }

    return true
}