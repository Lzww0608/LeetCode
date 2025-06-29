const N int = 26
func minimumKeypresses(s string) (ans int) {
    cnt := make([]int, N)
    for _, c := range s {
        cnt[c - 'a']++
    }
    sort.Slice(cnt, func(i, j int) bool {
        return cnt[i] > cnt[j]
    })

    for i := range N {
        if i < 9 {
            ans += cnt[i]
        } else if i < 18 {
            ans += cnt[i] * 2
        } else {
            ans += cnt[i] * 3
        }
    }

    return
}