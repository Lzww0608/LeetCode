const N int = 26
func numberOfSubstrings(s string, k int) (ans int) {
    cnt := make([]int, N)
    n := len(s)
    t := 0

    for l, r := 0, 0; r < n; r++ {
        x := int(s[r] - 'a')
        if cnt[x]++; cnt[x] == k {
            t++
        }

        for t >= 1 {
            y := int(s[l] - 'a')
            if cnt[y]--; cnt[y] == k - 1 {
                t--
            }
            ans += n - r
            l++
        }
    }

    return
}