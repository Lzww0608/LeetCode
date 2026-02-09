func beautifulNumbers(left int, right int) int {
    l := strconv.Itoa(left)
    r := strconv.Itoa(right)
    n := len(r)
    dif := n - len(l)
    memo := make(map[int]int)

    var dfs func(int, int, int, bool, bool) int 
    dfs = func(i, m, s int, isLow, isHigh bool) int {
        if i == n {
            if s != 0 && m % s == 0 {
                return 1
            }
            return 0
        }        
        mask := (m << 32) | (i << 16) | s 
        if _, ok := memo[mask]; !isLow && !isHigh && ok {
            return memo[mask]
        }

        low, high := 0, 9
        if isLow && i >= dif {
            low = int(l[i - dif] - '0')
        }
        if isHigh {
            high = int(r[i] - '0')
        }

        d := low
        res := 0
        if isLow && i < dif {
            res = dfs(i + 1, 1, 0, true, false)
            d = 1
        }

        for ; d <= high; d++ {
            res += dfs(i + 1, m * d, s + d, isLow && d == low, isHigh && d == high) 
        }
        if !isLow && !isHigh {
            memo[mask] = res
        }

        return res
    }

    return dfs(0, 1, 0, true, true)
}