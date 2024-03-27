func minimumDeletions(s string) int {
    ans, cnt := 0, 0
    for _, c := range s {
        if c == 'b' {
            cnt++
        } else {
            ans = min(ans + 1, cnt)
        }
    }
    return ans
}