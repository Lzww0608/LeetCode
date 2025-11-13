func isGood(nums []int) bool {
    n := len(nums) - 1
    for i := range n {
        if nums[i] == n {
            nums[i], nums[n] = nums[n], nums[i]
            break
        }
    }

    for i := range n {
        if i + 1 == nums[i] {
            continue
        }

        if nums[i] == n && nums[n] != n {
            nums[i], nums[n] = nums[n], nums[i]
        }

        for i != nums[i] - 1 {
            if nums[i] > n || nums[nums[i] - 1] == nums[i] {
                return false
            } 
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        }
    }

    return nums[n] == n
}