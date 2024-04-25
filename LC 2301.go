func matchReplacement(s string, sub string, mappings [][]byte) bool {
    mp := [255][255]bool{}
    for _, v := range mappings {
        mp[v[0]][v[1]] = true
    }
    

    n, m := len(s), len(sub) 
next:
    for i := 0; i + m <= n; i++ {
        for j := 0; j < m; j++ {
            if s[i+j] != sub[j] && !mp[sub[j]][s[i+j]] {
                continue next
            }
        }
        return true
    }

    return false
}