func minSwaps(nums []int) int {
    sum := 0
    n := len(nums)
    for _, x := range nums {
        sum += x
    }
    cur := 0
    for i := 0; i < sum; i++ {
        cur += nums[i]
    }
    ans := sum - cur
    for i := 0; i < n; i++ {
        cur = cur - nums[i] + nums[(i+sum) % n]
        ans = min(ans, sum - cur)
    }

    return ans
}