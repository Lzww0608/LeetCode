func findReplaceString(s string, indices []int, sources []string, targets []string) string {
    n := len(s)
    replaceStr := make([]string, n)
    replaceLen := make([]int, n)
    for i, id := range indices {
        if strings.HasPrefix(s[id:], sources[i]) {
            replaceStr[id] = targets[i]
            replaceLen[id] = len(sources[i])
        }
    }

    ans := &strings.Builder{}
    for i := 0; i < n; i++ {
        if replaceStr[i] == "" {
            ans.WriteByte(s[i])
        } else {
            ans.WriteString(replaceStr[i])
            i += replaceLen[i]-1
        }
    }
    return ans.String()
}