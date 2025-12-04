const MOD int = 1_000_000_007
func countNumbers(l string, r string, b int) int {
    return (solve(r, b) - solve(l, b) + check(l, b) + MOD) % MOD
}

func check(t string, b int) int {
    m := new(big.Int)
    m, _ = m.SetString(t, 10)
    s := m.Text(b)
    n := len(s)

    for i := 1; i < n; i++ {
        if s[i] < s[i - 1] {
            return 0
        }
    }

    return 1
}

func solve(t string, b int) int {
    m := new(big.Int)
    m, _ = m.SetString(t, 10)
    s := m.Text(b)
    n := len(s)
    memo := make([][11]int, n)
    for i := range n {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(i int, isNum, isLimit bool, pre int) int 
    dfs = func(i int, isNum, isLimit bool, pre int) int  {
        if i == n {
            return 1
        }

        p := &memo[i][pre]
        if isNum && !isLimit && *p != -1 {
            return *p
        }

        low, high := pre, b - 1
        if isLimit {
            high = int(s[i] - '0')
        }

        res := 0
        for d := low; d <= high; d++ {
            res += dfs(i + 1, isNum || d != 0, isLimit && d == high, d)
        }
        res %= MOD

        if isNum && !isLimit {
            *p = res 
        }
        return res
    }

    return dfs(0, false, true, 0)
} 