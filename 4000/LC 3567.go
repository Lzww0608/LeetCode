func minAbsDiff(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m - k + 1)
    for i := range ans {
        ans[i] = make([]int, n - k + 1)
    }

    for i := range len(ans) {
        for j := range len(ans[0]) {
            cnt := make(map[int]int)
            for p := i; p < i + k; p++ {
                for q := j; q < j + k; q++ {
                    cnt[grid[p][q]]++
                }
            }
            a := make([]int, 0, len(cnt))
            for k := range cnt {
                a = append(a, k)
            }
            sort.Ints(a)
            mx := a[len(a) - 1] - a[0]
            for k := 1; k < len(a); k++ {
                mx = min(mx, a[k] - a[k - 1])
            }
            ans[i][j] = mx
        }
    }

    return ans
}