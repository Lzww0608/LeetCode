func beautySum(s string) (ans int) {
    n := len(s)

    for i := range s {
        m := make([]int, 26)
        cnt := make([]int, 500)
        mx, mn := 0, 0
        for j := i; j < n; j++ {
            x := int(s[j] - 'a')
            if m[x] == 0 {
                mn = 1
            } else if m[x] == mn && cnt[mn] == 1 {
                mn++
            }
            cnt[m[x]]--
            m[x]++
            cnt[m[x]]++
            mx = max(mx, m[x])
            ans += mx - mn
        }
    }

    return
}
