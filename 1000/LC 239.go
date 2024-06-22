func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    ans := make([]int, 0, n - k + 1)
    q := []int{}

    for i, x := range nums {
        for len(q) > 0 && nums[q[len(q)-1]] <= x {
            q = q[:len(q)-1]
        }
        q = append(q, i)
        if i - q[0] >= k {
            q = q[1:]
        }
        if i >= k - 1 {
            ans = append(ans, nums[q[0]])
        }
    }

    return ans
}
