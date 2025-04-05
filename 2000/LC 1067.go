func digitsCount(d int, low int, high int) int {
    var dfs func(int, int) int
    dfs = func(x, y int) int {
        if y < 10 {
            if x <= y {
                return 1
            }
            return 0
        }

        t := y / 10
        k := y % 10
        ans := t + dfs(x, k)

        ans += (dfs(x, t - 1) - dfs(x, 0)) * 10 + (dfs(x, t) - dfs(x, t - 1)) * (k + 1)
        return ans
    }

    return dfs(d, high) - dfs(d, low - 1)
}