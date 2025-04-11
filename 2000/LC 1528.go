func restoreString(s string, indices []int) string {
    ans := make([]byte, len(s))
    for i := 0; i < len(s); i++ {
        ans[indices[i]] = s[i]
    }

    return string(ans)
}