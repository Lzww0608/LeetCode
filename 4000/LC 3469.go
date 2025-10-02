func minCost(nums []int) int {
    n := len(nums)
    mx := slices.Max(nums)
    if n <= 2 {
        return mx
    } else if n == 3 {
        return mx + min(nums[0], nums[1], nums[2])
    }
    memo := make([]map[int]int, n)
    for i := range n {
        memo[i] = make(map[int]int)
    }

    var dfs func(int, int) int 
    dfs = func(x, i int) int {
        if i >= n {
            return 0
        }

        if v, ok := memo[i][x]; ok {
            return v
        } 

        if i == n - 1 {
            return max(x, nums[i])
        }

        res := math.MaxInt32
        y, z := nums[i], nums[i + 1]
        if i + 1 == n - 1 {
            memo[i][x] = max(x, y, z) + min(x, y, z)
            return memo[i][x]
        }
        // x, y 
        t := max(x, y)
        res = min(res, t + dfs(z, i + 2))
        // x, z 
        t = max(x, z)
        res = min(res, t + dfs(y, i + 2))
        // y, z
        t = max(y, z)
        res = min(res, t + dfs(x, i + 2))
        memo[i][x] = res 
        return res
    }
    
    x, y, z := nums[0], nums[1], nums[2]
    return min(max(x, y) + dfs(z, 3), max(x, z) + dfs(y, 3), max(y, z) + dfs(x, 3))
}