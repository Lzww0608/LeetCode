func countBalanced(low int64, high int64) int64 {
    return int64(solve(int(high)) - solve(int(low) - 1))
}

func solve(x int) int {
    if x < 10 {
        return 0
    }
    s := strconv.Itoa(x)
    n := len(s)
    m := make(map[int]int)

    var dfs func(i, sum int, odd, isLimit, isNum bool) int 
    dfs = func(i, sum int, odd, isLimit, isNum bool) int {
        if i == n {
            if isNum && sum == 0 {
                return 1
            }
            return 0
        } 
        if abs(sum) > (n - i + 1) / 2 * 9 {
            return 0
        }
        
        mask := (sum << 6) + (i << 1)
        if !odd {
            mask += 1
        }
        if v, ok := m[mask]; ok && !isLimit && isNum {
            return v
        }

        res := 0
        if !isNum {
            res += dfs(i + 1, 0, true, false, false)
            for d := 1; d < 10; d++ {
                if i == 0 && d > int(s[0] - '0') {
                    break
                }
                res += dfs(i + 1, d, false, i == 0 && d == int(s[0] - '0'), true)
            }
        } else {
            l, r := 0, 9 
            if isLimit {
                r = int(s[i] - '0')
            }
            for d := l; d <= r; d++ {
                cur := sum 
                if odd {
                    cur += d 
                } else {
                    cur -= d
                }
                res += dfs(i + 1, cur, !odd, isLimit && d == r, true)
            }
        }
        

        if !isLimit && isNum {
            m[mask] = res 
        }
        return res
    }

    return dfs(0, 0, true, true, false)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}