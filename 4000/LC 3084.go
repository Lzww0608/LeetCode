func countSubstrings(s string, c byte) int64 {
    ans := 0
    cnt := 0
    for i := range s {
        if s[i] == c {
            cnt++
            ans += cnt
        }
    }

    return int64(ans)
}