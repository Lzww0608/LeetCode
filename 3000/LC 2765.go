func alternatingSubarray(nums []int) int {
    ans := -1
    for i := range nums {
        k := 1
        for j := i + 1; j < len(nums); j++ {
            if nums[j] - nums[j-1] == k {
                k = -k
                ans = max(ans, j - i + 1)
            } else {
                break
            }
        }
        if len(nums) - i + 1 <= ans {
            break
        }
    }
    return ans
}
