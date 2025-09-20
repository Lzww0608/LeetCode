func stringHash(s string, k int) string {
    ans := strings.Builder{}

    cur := 0
    for i := range s {
        cur += int(s[i] - 'a')
        cur %= 26
        if i % k == k - 1 {
            ans.WriteByte(byte('a' + cur))
            cur = 0
        }
    }

    return ans.String()
}