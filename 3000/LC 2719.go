const MOD int = 1_000_000_007
func count(num1 string, num2 string, min_sum int, max_sum int) int {
    n := len(num2)
    num1 = strings.Repeat("0", n - len(num1)) + num1

    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, min(n * 9, max_sum) + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int, bool, bool) int 
    dfs = func(i, sum int, low_limit, hight_limit bool) int {
        if sum > max_sum {
            return 0
        } else if i == n {
            if sum >= min_sum {
                return 1
            } else {
                return 0
            }
        }

        if !low_limit && !hight_limit && memo[i][sum] != -1 {
            return memo[i][sum]
        }

        res := 0
        l, r := 0, 9 
        if low_limit {
            l = int(num1[i] - '0')
        }
        if hight_limit {
            r = int(num2[i] - '0')
        }

        for j := l; j <= r; j++ {
            res = (res + dfs(i + 1, sum + j, low_limit && j == l, hight_limit && j == r)) % MOD
        }

        if !low_limit && !hight_limit {
            memo[i][sum] = res
        }

        return res
    }

    return dfs(0, 0, true, true)
}