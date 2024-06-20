func findIntegers(n int) int {
    s := strconv.FormatInt(int64(n), 2)
    m := len(s)
    f := make([][2]int, m)
    for i := range f {
        f[i] = [2]int{-1, -1}
    }
    
    var dfs func(int, int, bool) int
    dfs = func(i, pre int, isLimit bool) (res int) {
        if i == m {
            return 1
        }
        if !isLimit {
            p := &f[i][pre]
            if *p >= 0 {
                return *p
            }
            defer func() {*p = res}()
        }
        up := 1
        if isLimit {
            up = int(s[i] & 1)
        }
        res = dfs(i + 1, 0, isLimit && up == 0)
        if pre == 0 && up == 1 {
            res += dfs(i + 1, 1, isLimit)
        }
        return
    }

    return dfs(0, 0, true)
}
