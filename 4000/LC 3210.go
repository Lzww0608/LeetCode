func getEncryptedString(s string, k int) string {
    ans := make([]byte, len(s))
    for i := range s {
        ans[i] = s[(i + k) % len(s)]
    }

    return string(ans)
}