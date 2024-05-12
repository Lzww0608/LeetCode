func minDays(n int) int {
    memo := map[int]int{}
    var dfs func(int) int 
    dfs = func(i int) int {
        if i <= 2 {
            return i
        }
        if v, ok := memo[i]; ok {
            return v
        }
        memo[i] = 1 + min(i % 2 + dfs(i / 2), i % 3 + dfs(i / 3))

        return memo[i]
    }
    return dfs(n)
}