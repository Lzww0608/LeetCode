func maxKDistinct(nums []int, k int) []int {
    sort.Ints(nums)
    ans := make([]int, 0, k)
    ans = append(ans, nums[len(nums) - 1])
    for i := len(nums) - 2; i >= 0 && len(ans) < k; i-- {
        if nums[i] != nums[i + 1] {
            ans = append(ans, nums[i])
        }
    }

    return ans
}