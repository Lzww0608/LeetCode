func permuteUnique(nums []int) (ans [][]int) {
    sort.Ints(nums)
    n := len(nums)
    path := []int{}
    col := make([]int, n)
    var dfs func()
    dfs = func() {
        if len(path) == n {
            ans = append(ans, append([]int(nil), path...))
            return
        }
        for i := 0; i < n; i++ {
            if col[i] == 1 {
                continue
            }
            if i > 0 && nums[i] == nums[i-1] && col[i-1] == 0 {
                continue
            }
            path = append(path, nums[i])
            col[i] = 1
            dfs()
            col[i] = 0
            path = path[:len(path)-1]
        }
    }
    dfs()
    return
}
