const MOD int = 1_000_000_007
func numWays(words []string, target string) int {
    m, n := len(words[0]), len(target)
    cnt := make([][26]int, m)
    for _, s := range words {
        for j := range s {
            x := s[j] - 'a'
            cnt[j][x]++
        }
    }

    f := make([][]int, m)
    for i := range f {
        f[i] = make([]int, n)
        for j := range f[i] {
            f[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(i, j int) (res int) {
        if j == n {
            return 1 
        } else if m - i < n - j {
            return 0 
        }

        if f[i][j] != -1 {
            return f[i][j]
        } 
        p := &f[i][j]
        defer func() {*p = res}()

        x := target[j] - 'a'
        res += dfs(i + 1, j) + cnt[i][x] * dfs(i + 1, j + 1)
        res %= MOD

        return 
    }
    
    return dfs(0, 0)
}