const MOD int = 1_000_000_007
func countSteppingNumbers(low string, high string) int {
    n := len(high)
    s := high
    memo := make([][10]int, n)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int, bool, bool) int
    dfs = func(i, pre int, isLimit, isNum bool) (ans int) {
        if i == n {
            if isNum {
                return 1
            }
            return 0
        }
        if isNum && !isLimit && memo[i][pre] != -1 {
            return memo[i][pre]
        }

        if !isNum {
            ans = dfs(i + 1, pre, false, false)
        }

        up, low := 9, 0
        if isLimit {
            up = int(s[i] - '0')
        }
        if !isNum {
            low = 1
        }
        for d := low; d <= up; d++ {
            if !isNum || abs(d - pre) == 1 {
                ans = (ans + dfs(i + 1, d, isLimit && d == up, true)) % MOD
            }
        }

        if !isLimit && isNum {
            memo[i][pre] = ans
        }

        return 
    }
    r := dfs(0, 0, true, false)
    s, n = low, len(low)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return (r - dfs(0, 0, true, false) + valid(low) + MOD) % MOD
}

func valid(s string) int {
    for i := 1; i < len(s); i++ {
        if abs(int(s[i] - '0') - int(s[i-1] - '0')) != 1 {
            return 0
        }
    }

    return 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x 
}