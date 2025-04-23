func countLargestGroup(n int) (ans int) {
    s := strconv.Itoa(n)
    m := len(s)
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, m * 9 + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int, bool) int 
    dfs = func(i, left int, isLimit bool) (res int) {
        if i == m {
            if left == 0 {
                return 1
            }
            return 0
        }

        if !isLimit {
            p := &memo[i][left]
            if *p != -1 {
                return *p 
            }
            defer func() {*p = res}()
        }

        up := 9 
        if isLimit {
            up = int(s[i] - '0')
        }

        for d := 0; d <= min(up, left); d++ {
            res += dfs(i + 1, left - d, isLimit && d == up)
        }
        return
    }

    mx := -1
    for i := 1; i <= m * 9; i++ {
        cur := dfs(0, i, true)
        if cur > mx {
            mx = cur
            ans = 1
        } else if cur == mx {
            ans++
        }
    }

    return
}