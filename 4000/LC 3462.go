func maxSum(grid [][]int, limits []int, k int) int64 {
    n, m := len(grid), len(grid[0])
    a := make([]int, 0, m * n)
    for idx, v := range grid {
        sort.Slice(v, func(i, j int) bool {
            return v[i] > v[j]
        })
        t := limits[idx]
        for i := 0; i < m && i < t; i++ {
            a = append(a, v[i])
        }
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i] > a[j]
    })
    ans := 0
    for i := 0; i < k; i++ {
        ans += a[i]
    }

    return int64(ans)
}