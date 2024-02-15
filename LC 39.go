func combinationSum(candidates []int, target int) (ans [][]int) {
    path := []int{}
    var dfs func(int, int)
    dfs = func(start int, sum int) {
        if sum == target {
            ans = append(ans, append([]int(nil), path...))
        }
        for i := start; i < len(candidates); i++ {
            if sum + candidates[i] > target {
                break
            }
            sum += candidates[i]
            path = append(path, candidates[i])
            dfs(i, sum)
            sum -= candidates[i]
            path = path[:len(path)-1]
        }
    }
    sort.Ints(candidates)
    dfs(0, 0)
    return
}
