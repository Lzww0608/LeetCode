func combinationSum3(k int, n int) (ans [][]int) {
    if n > 45 {
        return 
    }

    path := []int{}
    var dfs func(int, int)
    dfs = func(idx, sum int) {
        if sum > n {
            return
        } else if len(path) == k {
            if sum == n {
                ans = append(ans, append([]int(nil), path...))
            }
            return
        }

        for i := idx; i < 10; i++ {
            path = append(path, i)
            sum += i 
            dfs(i + 1, sum)
            sum -= i 
            path = path[:len(path)-1]
        }
    }   
    dfs(1, 0)
    
    return 
}
