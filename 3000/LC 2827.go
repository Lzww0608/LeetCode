func numberOfBeautifulIntegers(low int, high int, k int) int {
    b := strconv.Itoa(high)
    a := strconv.Itoa(low - 1)
    return solve(b, k) - solve(a, k)   
}


func solve(s string, k int) int {
    n := len(s)
    memo := make([][][]int, n + 1)
    for i := range memo {
        memo[i] = make([][]int, k)
        for j := range memo[i] {
            memo[i][j] = make([]int, n * 2 + 1)
            for t := range memo[i][j] {
                memo[i][j][t] = -1
            }
        }
    }

    var dfs func(int, int, int, bool, bool) int
    dfs = func(i, val, dif int, isNum, isLimit bool) (res int) {
        if i == n {
            if isNum && val == 0 && dif == n {
                return 1
            } else {
                return 0
            }
        }

        if !isLimit && isNum {
            p := &memo[i][val][dif]
            if *p >= 0 {
                return *p
            }
            defer func() {*p = res} ()
        }

        if !isNum {
            res += dfs(i + 1, val, dif, false, false)
        }
        up := 9
        if isLimit {
            up = int(s[i] - '0')
        }
        d := 0
        if !isNum {
            d = 1
        }
        for ; d <= up; d++ {
            res += dfs(i + 1, (val * 10 + d) % k, dif + d % 2 * 2 - 1, true, isLimit && d == up)
        }
        return
    }

    return dfs(0, 0, n, false, true)
}