func maxResult(nums []int, k int) int {
    n := len(nums)
    f := make([]int, n)
    f[0] = nums[0]
    q := []int{0}
    for i := 1; i < n; i++ {
        for len(q) > 0 && i - q[0] > k {
            q = q[1:]
        }
        f[i] = f[q[0]] + nums[i]
        for len(q) > 0 && f[q[len(q)-1]] < f[i] {
            q = q[:len(q)-1]
        }
        q = append(q, i)
    }

    return f[n-1]
}