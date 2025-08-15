func arrayNesting(nums []int) (ans int) {
    n := len(nums)
    vis := make([]int, n)
    for i := range vis {
        vis[i] = -1
    }

    var dfs func(x int) int 
    dfs = func(x int) int {
        if vis[x] != -1 {
            return vis[x]
        }

        vis[x] = 0
        vis[x] = 1 + dfs(nums[x])
        return vis[x]
    }

    for i := 0; i < n; i++ {
        ans = max(ans, dfs(i))
    }

    return 
}