func findPermutation(nums []int) []int {
    n := len(nums)
    memo := make([][]int, 1 << n)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs2 func(int, int) int 
    dfs2 = func(s, j int) int {
        if s + 1 == 1 << n {
            return abs(nums[0] - j)
        }

        p := &memo[s][j]
        if *p != -1 {
            return *p
        }

        res := math.MaxInt
        for k := 1; k < n; k++ {
            if (s >> k) & 1 == 0 {
                res = min(res, dfs2(s | (1 << k), k) + abs(j - nums[k]))
            }
        }
        *p = res
        return res
    }

    ans := make([]int, 0, n)
    var dfs func(int, int)
    dfs = func(s, j int) {
        ans = append(ans, j)
        if s + 1 == 1 << n {
            return
        }

        t := dfs2(s, j)
        for k := 1; k < n; k++ {
            if (s >> k) & 1 == 0 && dfs2(s | (1 << k), k) + abs(j - nums[k]) == t {
                dfs(s | (1 << k), k)
                break
            }
        }
    }
    dfs(1, 0)
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}