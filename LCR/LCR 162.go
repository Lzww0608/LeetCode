func digitOneInNumber(num int) int {
    s := strconv.Itoa(num)
    n := len(s)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int, bool) int 
    dfs = func(i, cnt int, isLimit bool) (res int) {
        if i == n {
            return cnt
        }
        if !isLimit && memo[i][cnt] != -1 {
            return memo[i][cnt]
        }

        up := 9 
        if isLimit {
            up = int(s[i] - '0')
        }

        for d := 0; d <= up; d++ {
            if d == 1 {
                res += dfs(i + 1, cnt + 1, isLimit && d == up)
            } else {
                res += dfs(i + 1, cnt, isLimit && d == up)
            }
        }

        if !isLimit {
            memo[i][cnt] = res
        }

        return 
    }

    return dfs(0, 0, true)
}