func permute(nums []int) [][]int {
    n := len(nums)
    ans := [][]int{}
    path := []int{}
    color := make([]int, n)
    var dfs func()
    dfs = func() {
        f := true
        for i, x := range nums {
            if color[i] == 0 {
                f = false
                color[i] = 1
                path = append(path, x)
                dfs()
                path = path[:len(path)-1]
                color[i] = 0
            }
        }
        if f {
            ans = append(ans, append([]int(nil), path...))
        }
    }
    dfs()
    return ans
}
