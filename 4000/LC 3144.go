func minimumSubstringsInPartition(s string) int {
    n := len(s)
    f := make([]int, n + 1)
    for i := range f {
        f[i] = n
    }
    f[0] = 0
    for i := range s {
        cnt := make([]int, 26)
        mx, k := 0, 0
        for j := i; j >= 0; j-- {
            x := int(s[j] - 'a')
            cnt[x]++
            if cnt[x] == 1 {
                k++
            }
            mx = max(mx, cnt[x])
            if i - j + 1 == mx * k {
                f[i+1] = min(f[i+1], f[j] + 1)
            }
        }
    }

    return f[n]
}