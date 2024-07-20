func combinationSum4(nums []int, target int) int {
    memo := make([]int, target + 1)
    for i := range memo {
        memo[i] = -1
    }

    var dfs func(int) int
    dfs = func(idx int) int {
        if idx == 0 {
            return 1
        }

        p := &memo[idx]
        if *p != -1 {
            return *p
        }

        ans := 0
        for _, x := range nums {
            if idx >= x {
                ans += dfs(idx - x)
            }
        }

        *p = ans
        return ans
    }

    return dfs(target)
}