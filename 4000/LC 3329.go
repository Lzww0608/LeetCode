const N int = 26
func numberOfSubstrings(s string, k int) int64 {
    ans := 0
    n := len(s)

    cnt := [N]int{}
    sum := 0
    for l, r := 0, 0; r < n; r++ {
        x := int(s[r] - 'a')
        if cnt[x]++; cnt[x] == k {
            sum++
        }

        for sum > 0 {
            y := int(s[l] - 'a')
            if cnt[y]--; cnt[y] == k - 1 {
                sum--
            }
            l++
        }

        ans += l
    }

    return int64(ans)
}