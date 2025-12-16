func lexSmallest(s string) string {
    n := len(s)
    ans := s 

    t := []byte(s)
    for k := 2; k <= n; k++ {
        slices.Reverse(t[:k])
        ans = min(ans, string(t))
        slices.Reverse(t[:k])

        slices.Reverse(t[n-k:])
        ans = min(ans, string(t))
        slices.Reverse(t[n-k:])
    }

    return ans
}