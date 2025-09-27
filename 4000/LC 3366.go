func minArraySum(nums []int, k int, op1 int, op2 int) int {
    n := len(nums)
    memo := make([][][]int, n)
    for i := range memo {
        memo[i] = make([][]int, op1 + 1)
        for j := range memo[i] {
            memo[i][j] = make([]int, op2 + 1)
            for p := range memo[i][j] {
                memo[i][j][p] = -1
            }
        }
    }

    var dfs func(int, int, int) int 
    dfs = func(i, a, b int) (res int) {
        if i == n {
            return 0
        }

        p := &memo[i][a][b]
        if *p != -1 {
            return *p
        }

        res = nums[i] + dfs(i + 1, a, b)
        if a > 0 {
            res = min(res, (nums[i] + 1) / 2 + dfs(i + 1, a - 1, b))
        }
        
        if nums[i] >= k && b > 0 {
            res = min(res, nums[i] - k + dfs(i + 1, a, b - 1))
            if a > 0 {
                res = min(res, (nums[i] - k + 1) / 2 + dfs(i + 1, a - 1, b - 1))
                if (nums[i] + 1) / 2 >= k {
                    res = min(res, (nums[i] + 1) / 2 - k + dfs(i + 1, a - 1, b - 1))
                }
            }
            
        }

        *p = res 
        return
    }

    return dfs(0, op1, op2)
}