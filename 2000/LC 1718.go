func constructDistancedSequence(n int) []int {
    path := make([]int, n*2-1)
    used := make([]bool, n+1)
    
    var dfs func(int) bool
    dfs = func(pos int) bool {
        if pos == len(path) {
            return true
        }
        
        if path[pos] != 0 {
            return dfs(pos + 1)
        }
        
        for num := n; num >= 1; num-- {
            if used[num] {
                continue
            }
            
            if num == 1 {
                path[pos] = num
                used[num] = true
                if dfs(pos + 1) {
                    return true
                }
                path[pos] = 0
                used[num] = false
            } else if pos + num < len(path) && path[pos+num] == 0 {
                path[pos] = num
                path[pos+num] = num
                used[num] = true
                if dfs(pos + 1) {
                    return true
                }
                path[pos] = 0
                path[pos+num] = 0
                used[num] = false
            }
        }
        return false
    }
    
    dfs(0)
    return path
}