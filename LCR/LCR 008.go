func minSubArrayLen(target int, nums []int) int {
    n := len(nums)
    ans := n + 1

    sum := 0
    for l, r := 0, 0; r < n; r++ {
        sum += nums[r]
        for sum >= target {
            ans = min(ans, r - l + 1)
            sum -= nums[l]
            l++
        }
    }

    if ans > n {
        return 0
    }

    return ans
}