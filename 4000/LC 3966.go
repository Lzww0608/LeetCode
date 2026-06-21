func goodIntegers(l int64, r int64, k int) int64 {
    return int64(solve(int(r), k) - solve(int(l - 1), k))
}

func solve(x, k int) int {
    if x < 10 {
        return x + 1
    }

    s := strconv.Itoa(x)
    n := len(s)
    memo := make([][10]int, n)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(i int, isNum, isLimit bool, pre int) int 
    dfs = func(i int, isNum, isLimit bool, pre int) int {
        if i == n {
            return 1
        }

        if isNum && !isLimit && memo[i][pre] != -1 {
            return memo[i][pre]
        }

        res := 0
        t := int(s[i] - '0')
        if pre == -1 {
            if i != 0 {
                t = 9
            } 
            for x := range t + 1 {
                if x == 0 {
                    res += dfs(i + 1, false, false, -1)
                } else if isLimit && x == t {
                    res += dfs(i + 1, true, true, x)
                } else {
                    res += dfs(i + 1, true, false, x)
                }
            }
            return res
        }

        l := max(0, pre - k)
        r := min(9, pre + k)
        if isLimit {
            r = min(r, t)
        }
        for x := l; x <= r; x++ {
            res += dfs(i + 1, isNum || x != 0, isLimit && x == t, x)
        }

        if isNum && !isLimit {
            memo[i][pre] = res
        }
        return res 
    }

    return dfs(0, false, true, -1)
}