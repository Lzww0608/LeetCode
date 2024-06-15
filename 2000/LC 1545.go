func findKthBit(n int, k int) byte {

    var dfs func(n, k, d int) byte 
    dfs = func(n, k, d int) byte {
        if k == 1 {
            return byte('0' + d & 1)
        } else if k == (1 << (n - 1)) || k == (1 << n) - 1 {
            return byte('1' - d & 1)
        } else if k < (1 << (n - 1)) {
            return dfs(n - 1, k, d)
        } 

        return dfs(n - 1, (1 << n) - k, d + 1)
    }
    
    return dfs(n, k, 0)
}