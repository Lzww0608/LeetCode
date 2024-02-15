func combinationSum2(candidates []int, target int) (ans [][]int) {
    path := []int{}
    var dfs func(start int, sum int)
    dfs = func(start int, sum int) {
        if sum == target {
            ans = append(ans, append([]int(nil), path...))
        }
        for i := start; i < len(candidates); i++ {
            if sum + candidates[i] > target {
                break
            }
            if i > start && candidates[i] == candidates[i-1] {
                continue
            }
            path = append(path, candidates[i])
            dfs(i + 1, sum + candidates[i])
            path = path[:len(path)-1]
        }
    }
    sort.Ints(candidates)
    dfs(0, 0)
    return
}