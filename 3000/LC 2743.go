const N int = 26
func numberOfSpecialSubstrings(s string) (ans int) {
    cnt := make([]int, 26)
    n := len(s)
    for l, r := 0, 0; r < n; r++ {
        x := int(s[r] - 'a')
        cnt[x]++
        for cnt[x] > 1 {
            y := int(s[l] - 'a')
            cnt[y]--
            l++
        }
        ans += r - l + 1
    }

    return
}