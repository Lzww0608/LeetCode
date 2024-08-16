func maximumScore(nums []int, multipliers []int) int {
    n, m := len(nums), len(multipliers)
    memo := make([][]int, m + 1)
    for i := range memo {
        memo[i] = make([]int, m + 1)
        for j := range memo[i] {
            memo[i][j] = math.MaxInt32
        }
    }

    var dfs func(int, int, int) int
    dfs = func(l, r, i int) int {
        if i == m {
            return 0
        }
        p := &memo[i][l]
        if *p != math.MaxInt32 {
            return *p
        }
        *p = max(nums[l] * multipliers[i] + dfs(l + 1, r, i + 1), nums[r] * multipliers[i] + dfs(l, r - 1, i + 1))
        return *p
    }

    return dfs(0, n - 1, 0)
}