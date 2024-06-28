func subsets(nums []int) (ans [][]int) {
    path := []int{}
    sort.Ints(nums)
    n := len(nums)

    var dfs func(int)
    dfs = func(i int) {
        
        if i == n {
            ans = append(ans, append([]int(nil), path...))
            return
        }

        dfs(i + 1)

        path = append(path, nums[i])
        dfs(i + 1)
        path = path[:len(path)-1]
    }
    dfs(0)

    return
}