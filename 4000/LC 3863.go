func minOperations(s string) int {
    n := len(s)
    if n == 2 && s[1] < s[0] {
        return -1
    } else if slices.IsSorted([]byte(s)) || n == 2 {
        return 0
    }
    
    t := []byte(s)
    if s[0] <= slices.Min(t[1:]) || s[n - 1] >= slices.Max(t[:n - 1]) {
        return 1
    } 
    sort.Slice(t, func(i, j int) bool {
        return t[i] < t[j]
    })
    if s[0] == t[n - 1] && s[n - 1] == t[0] && s[0] != t[n - 2] && s[n - 1] != t[1] {
        return 3
    }

    return 2
}