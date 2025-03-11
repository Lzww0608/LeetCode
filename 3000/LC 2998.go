func minimumOperationsToMakeEqual(x int, y int) int { 
    memo := make([]int, x + 1)
    
    var dfs func(int, int) int 
    dfs = func(x, y int) (res int) {
        if x <= y {
            return y - x
        } 

        if memo[x] != 0 {
            return memo[x]
        }

        res = x - y
        res = min(res, dfs(x / 11, y) + 1 + x % 11)
        if x % 11 != 0 {
            res = min(res, dfs(x / 11 + 1, y) + 1 + (11 - x % 11))
        }
        res = min(res, dfs(x / 5, y) + 1 + x % 5)
        if x % 5 != 0 {
            res = min(res, dfs(x / 5 + 1, y) + 1 + (5 - x % 5))
        }
        memo[x] = res
        return res
    }

    return dfs(x, y)
}

