func countGoodSubarrays(a []int) int64 {
    n := len(a)
    st := []int{-1}
    left := make([]int, n)
    for i, x := range a {
        for len(st) > 1 && a[st[len(st) - 1]] | x == x {
            st = st[:len(st) - 1]
        }
        left[i] = st[len(st) - 1]
        st = append(st, i)
    }

    st[0] = n 
    st = st[:1]
    ans := 0
    for i := n - 1; i >= 0; i-- {
        x := a[i]
        for len(st) > 1 && a[st[len(st) - 1]] != x && a[st[len(st) - 1]] | x == x {
            st = st[:len(st) - 1]
        }
        l, r := left[i], st[len(st) - 1]
        ans += (i - l) * (r - i)
        st = append(st, i)
    }

    return int64(ans)
}