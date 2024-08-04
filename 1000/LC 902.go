func atMostNGivenDigitSet(digits []string, n int) int {
    s := strconv.Itoa(n)
    m := len(s)

    memo := make([]int, m)
    for i := range memo {
        memo[i] = -1
    }

    var dfs func(int, bool, bool) int
    dfs = func(i int, isLimit, isNum bool) (res int) {
        if i == m {
            if isNum {
                return 1
            }
            return
        }

        if !isLimit && isNum {
            p := &memo[i]
            if *p != -1 {
                return *p
            }
            defer func(){*p = res}()
        }

        if !isNum {
            res += dfs(i + 1, false, false)
        }

        up := byte('9')
        if isLimit {
            up = s[i]
        }
        for _, v := range digits {
            if v[0] <= up {
                res += dfs(i + 1, isLimit && v[0] == up, true)
            }
        }

        return
    }

    return dfs(0, true, false)
}