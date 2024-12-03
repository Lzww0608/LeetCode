const MOD int = 1_000_000_007
func totalStrength(a []int) (ans int) {
    n := len(a)
    pre := make([]int, n + 2)
    s := 0
    for i, x := range a {
        s += x
        pre[i+2] = (pre[i+1] + s) % MOD 
    }

    left := make([]int, n)
    right := make([]int, n)
    for i := range right {
        right[i] = n 
    }

    st := []int{-1}
    for i, x := range a {
        for len(st) > 1 && a[st[len(st)-1]] >= x {
            right[st[len(st)-1]] = i 
            st = st[:len(st)-1]
        }
        left[i] = st[len(st)-1]
        st = append(st, i)
    }

    for i, x := range a {
        l, r := left[i] + 1, right[i] - 1
        sum := ((i - l + 1) * (pre[r+2] - pre[i+1]) - (r - i + 1) * (pre[i+1] - pre[l])) % MOD
        ans = (ans + x * sum) % MOD
    }

    return (ans + MOD) % MOD
}