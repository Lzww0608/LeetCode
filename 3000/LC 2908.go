func minimumSum(nums []int) int {
    ans := math.MaxInt32
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] < nums[j] && nums[k] < nums[j] {
                    ans = min(ans, nums[i] + nums[k] + nums[j])
                }
            }
        }
    }
    if ans == math.MaxInt32 {
        return -1
    }

    return ans
}
