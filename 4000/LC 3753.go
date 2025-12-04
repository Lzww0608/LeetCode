func totalWaviness(num1 int64, num2 int64) int64 {
    s1 := strconv.Itoa(int(num1) - 1)
    s2 := strconv.Itoa(int(num2))

    return int64(solve(s2) - solve(s1))
}

func solve(s string) int {
    n := len(s)
    memo := make([][11][11][20]int, n)

    var dfs func(i int, isNum, isLimit bool, last1, last2, sum int) int 
    dfs = func(i int, isNum, isLimit bool, last1, last2, sum int) int {
        if i == n {
            return sum
        }
        p := &memo[i][last1 + 1][last2 + 1][sum]
        if *p != 0 && isNum && !isLimit {
            return *p
        }

        h := int(s[i] - '0')
        if !isLimit {
            h = 9
        }

        res := 0
        for x := range h + 1 {
            if !isNum {
                t := x 
                if x == 0 {
                    t = -1
                }
                res += dfs(i + 1, x != 0, isLimit && x == h, t, -1, 0)
            } else {
                cur := sum 
                if last2 > -1 && ((last1 > x && last1 > last2) || (last1 < x && last1 < last2)) {
                    cur++
                }
                res += dfs(i + 1, true, isLimit && x == h, x, last1, cur)
            }
        }
        if !isLimit && isNum {
            *p = res 
        }
        
        return res
    }

    return dfs(0, false, true, -1, -1, 0)
}