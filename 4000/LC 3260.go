func largestPalindrome(n int, k int) string {
    pow := make([]int, n)
    pow[0] = 1
    for i := 1; i < n; i++ {
        pow[i] = pow[i - 1] * 10 % k
    }

    ans := make([]byte, n)
    m := (n + 1) >> 1
    vis := make([][]bool, m + 1)
    for i := range vis {
        vis[i] = make([]bool, k)
    }

    var dfs func(int, int) bool
    dfs = func(i, j int) bool {
        if i == m {
            return j == 0
        }
        vis[i][j] = true 
        for d := 9; d >= 0; d-- {
            p := 0
            if n % 2 == 1 && i == m - 1 {
                p = (j + d * pow[i]) % k 
            } else {
                p = (j + d * (pow[i] + pow[n - i - 1])) % k
            }

            if !vis[i + 1][p] && dfs(i + 1, p) {
                ans[i], ans[n - i - 1] = byte('0' + d), byte('0' + d)
                return true
            }
        }
        return false
    }
    dfs(0, 0)
    return string(ans)
}