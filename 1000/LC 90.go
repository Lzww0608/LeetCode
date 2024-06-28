func subsetsWithDup(nums []int) (ans [][]int) {
    sort.Ints(nums)
    n := len(nums)
    path := []int{}

    var dfs func(int)
    dfs = func(idx int) {
        ans = append(ans, append([]int(nil), path...))
        for i := idx; i < n; i++ {
            if i > idx && nums[i] == nums[i-1] {
                continue
            }
            path = append(path, nums[i])
            dfs(i + 1)
            path = path[:len(path)-1]
        }
    }
    dfs(0)

    return
}
