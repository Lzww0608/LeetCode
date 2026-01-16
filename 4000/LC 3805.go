func countPairs(words []string) int64 {
    ans := 0
    cnt := make(map[string]int)
    for _, s := range words {
        t := solve(s)
        ans += cnt[t]
        cnt[t]++
    }

    return int64(ans)
}

func solve(s string) string {
    n := len(s)
    t := make([]byte, n)
    t[0] = 'a'
    for i := 1; i < n; i++ {
        x := int(s[i] + 26 - s[i - 1]) % 26
        t[i] = byte('a' + (int(t[i - 1] - 'a') + x) % 26)
    }

    return string(t)
}