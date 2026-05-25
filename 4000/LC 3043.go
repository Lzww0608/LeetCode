func longestCommonPrefix(a []int, b []int) (ans int) {
    vis := make(map[string]bool)
    for _, x := range a {
        s := strconv.Itoa(x)
        for j := range s {
            vis[s[:j + 1]] = true
        }
    }
    
    for _, x := range b {
        s := strconv.Itoa(x)
        for j := ans; j < len(s); j++ {
            if vis[s[:j + 1]] {
                ans = max(ans, j + 1)
            }
        }
    }
    
    return 
} 