func maxAltitude(heights []int, k int) []int {
    n := len(heights)
    if n == 0 {
        return nil
    }
    ans := make([]int, n - k + 1)
    q := []int{}

    for i, x := range heights {
        for len(q) > 0 && heights[q[len(q)-1]] < x {
            q = q[:len(q)-1]
        }
        q = append(q, i)
        if i - q[0] >= k {
            q = q[1:]
        }

        if i >= k - 1 {
            ans[i-k+1] = heights[q[0]]
        }
    }

    return ans
}
