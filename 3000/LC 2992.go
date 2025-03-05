func selfDivisiblePermutationCount(n int) int {
    memo := make([][]int, n + 1)
    for i := range memo {
        memo[i] = make([]int, n + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }func selfDivisiblePermutationCount(n int) int {
    memo := make([][]int, n + 1)
    for i := range memo {
        memo[i] = make([]int, 1 << n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    gcd := func(x, y int) int {
        for y != 0 {
            x, y = y, x % y 
        }
        return x 
    }
    
    var dfs func(int, int) int
    dfs = func(cur, mask int) (ans int) {
        if mask == 0 {
            return 1
        }
        if memo[cur][mask] != -1 {
            return memo[cur][mask]
        }
        for j := 0; j < n; j++ {
            if mask & (1 << j) != 0 && gcd(cur, j + 1) == 1 {
                ans += dfs(cur + 1, mask ^ (1 << j))
            }
        }
        memo[cur][mask] = ans
        return 
    }

    
    return dfs(1, (1 << n) - 1)
}



    gcd := func(x, y int) int {
        for y != 0 {
            x, y = y, x % y 
        }
        return x 
    }
    
    var dfs func(int, int) int
    dfs = func(cur, mask int) (ans int) {
        if mask == 0 {
            return 1
        }
        for j := 0; j < n; j++ {
            if mask & (1 << j) != 0 && gcd(cur, j + 1) == 1 {
                ans += dfs(cur + 1, mask ^ (1 << j))
            }
        }
        return 
    }

    
    return dfs(1, (1 << n) - 1)
}

