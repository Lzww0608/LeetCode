func findLUSlength(strs []string) int {
    sort.Slice(strs, func(i, j int) bool {
        if len(strs[i]) == len(strs[j]) {
            return strs[i] < strs[j]
        }
        return len(strs[i]) < len(strs[j])
    })
    ans := -1
    n := len(strs)
    for i := 0; i < n; i++ {
        if i > 0 && strs[i] == strs[i-1] {
            continue
        } 
        if i + 1 < n && strs[i] == strs[i+1] {
            continue
        }  
        f := true
        for j := i + 1; j < n; j++ {
            if contain(strs[j], strs[i]) {
                f = false
                break
            }
        }
        if f {
            ans = max(ans, len(strs[i]))
        }
    } 
    return ans
}

func contain(a, b string) bool {
    i, j := 0, 0
    n, m := len(a), len(b)
    for i < n && j < m {
        if a[i] == b[j] {
            i++
            j++
        } else {
            i++
        }
    }
    return j == m
}
