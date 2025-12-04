func subarraySum(nums []int) (ans int) {
    for i, x := range nums {
        j := max(0, i - x)
        if i > 0 {
            nums[i] += nums[i - 1]
        }
        if j == 0 {
            ans += nums[i]
        } else {
            ans += nums[i] - nums[j - 1]
        }
    }

    return
}