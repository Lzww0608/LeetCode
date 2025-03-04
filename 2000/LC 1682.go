const N int = 27
func longestPalindromeSubseq(s string) int {
    n := len(s)
    memo := make([][][N]int, n + 1)
    for i := range memo {
        memo[i] = make([][N]int, n + 1)
        for j := range memo[i] {
            for k := 0; k < N; k++ {
                memo[i][j][k] = -1
            }
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, j, last int) int {
        if i >= j {
            return 0
        } 
        p := &memo[i][j][last]
        if *p != -1 {
            return *p 
        }
        res := 0 
        if s[i] == s[j] && int(s[i] - 'a') != last {
            res = max(res, dfs(i + 1, j - 1, int(s[i] - 'a')) + 2)
        } 
        res = max(res, dfs(i + 1, j, last), dfs(i, j - 1, last))
        *p = res
        return res
    }

    return dfs(0, n - 1, N - 1)
}