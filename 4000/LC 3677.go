func countBinaryPalindromes(n int64) int {
    if n <= 1 {
        return int(n) + 1
    } 
    s := fmt.Sprintf("%b", n)
    m := len(s)

    memo := make(map[int]int)
    var dfs func(int, bool, bool) int 
    dfs = func(i int, limit, f bool) int {
        if i >= (m + 1) / 2 {
            if f {
                return 0
            }
            return 1
        }

        if v, ok := memo[i]; ok && !limit && !f {
            return v 
        }

        res := 0
        if limit {
            if s[i] == '0'{
                if s[m - i - 1] == '0' {
                    res += dfs(i + 1, true, f)
                } else {
                    res += dfs(i + 1, true, false)
                }
            } else {
                // 0
                res += dfs(i + 1, false, false)
                // 1
                res += dfs(i + 1, true, f || s[m - i - 1] == '0')
            }
        } else {
            res = 2 * dfs(i + 1, limit, false)
            memo[i] = res
        }

        return res
    } 

    ans := dfs(1, true, s[m - 1] == '0')
    for i := 1; i < m; i++ {
        ans += 1 << ((m - i + 1) / 2 - 1)
    }

    return ans + 1
}