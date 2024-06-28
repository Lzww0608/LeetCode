func diStringMatch(s string) []int {
    n := len(s)
    ans := make([]int, n + 1)
    l, r := 0, n
    for i := range s {
        if s[i] == 'D' {
            ans[i] = r
            r--
        } else {
            ans[i] = l
            l++
        }
    }
    ans[n] = r
    return ans
}
