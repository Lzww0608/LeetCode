func pondSizes(land [][]int) (ans []int) {
    n, m := len(land), len(land[0])
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

    var dfs func(int, int) int
    dfs = func(i, j int) (res int) {
        if land[i][j] != 0 {
            return 0
        }
        land[i][j] = 1
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < n && y >= 0 && y < m {
                res += dfs(x, y)
            }
        }
        return res + 1
    }

    for i := range land {
        for j := range land[i] {
            if land[i][j] == 0 {
                ans = append(ans, dfs(i, j))
            }
        }
    }
    sort.Ints(ans)

    return
}