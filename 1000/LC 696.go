func countBinarySubstrings(s string) (ans int) {
    n := len(s)
    pre := 0
    for i := 0; i < n;  {
        c := s[i]
        cnt := 0
        for i < n && s[i] == c {
            i++
            cnt++
        }
        ans += min(pre, cnt)
        pre = cnt
    }

    return 
}
