func sortByAbsoluteValue(nums []int) []int {
    sort.Slice(nums, func(i, j int) bool {
        return abs(nums[i]) < abs(nums[j])
    })

    return nums
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}