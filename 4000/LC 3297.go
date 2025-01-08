const N int = 26
func validSubstringCount(s string, t string) int64 {
    ans := 0
    cnt := make([]int, N)
    for i := range t {
        x := int(t[i] - 'a')
        cnt[x]++
    }

    n := len(s)
    cur := make([]int, N)

    check := func() bool {
        for j := 0; j < N; j++ {
            if cur[j] < cnt[j] {
                return false
            }
        }

        return true
    }

    for l, r := 0, 0; r < n; r++ {
        x := int(s[r] - 'a')
        cur[x]++

        for check() {
            ans += n - r
            y := int(s[l] - 'a')
            cur[y]--
            l++
        }
    }

    return int64(ans)
}