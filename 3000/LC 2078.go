func maxDistance(colors []int) int {
    n := len(colors)
    if colors[0] != colors[n-1] {
        return n - 1
    }

    for i := 0; i < n; i++ {
        if colors[i] != colors[n - 1] || colors[n - 1 - i] != colors[0] {
            return n - 1 - i
        }
    }

    return 1
}