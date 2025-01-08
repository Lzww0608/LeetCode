func maximumLengthSubstring(s string) (ans int) {
    cnt := make(map[byte]int)
    n := len(s)
    for l, r := 0, 0; r < n; r++ {
        cnt[s[r]]++
        for cnt[s[r]] > 2 {
            cnt[s[l]]--
            l++
        }
        ans = max(ans, r - l + 1)
    }

    return 
}