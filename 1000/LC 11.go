func maxArea(height []int) (ans int) {
    n := len(height)
    l, r := 0, n - 1
    for l < r {
        ans = max(ans, (r - l) * min(height[l], height[r]))
        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }

    return
}