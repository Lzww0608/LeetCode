func numDupDigitsAtMostN(n int) int {
    s := strconv.Itoa(n)
    m := len(s)
    memo := make([][1<<10]int, m)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int, bool, bool) int 
    dfs = func(i, mask int, isLimit, isNum bool) (res int) {
        if i == m {
            if isNum {
                return 1
            }
            return
        }

        if !isLimit && isNum {
            p := &memo[i][mask]
            if *p >= 0 {
                return *p
            }
            defer func() {*p = res} ()
        }

        if !isNum {
            res += dfs(i + 1, mask, false, false)
        }

        d := 0
        if !isNum {
            d = 1
        }
        up := 9
        if isLimit {
            up = int(s[i] - '0')
        }
        for ; d <= up; d++ {
            if mask >> d & 1 == 0 {
                res += dfs(i + 1, mask | 1 << d, isLimit && d == up, true)
            }
        }
        return
    }

    return n - dfs (0, 0, true, false)
}