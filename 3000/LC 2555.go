import "sort"
func maximizeWin(a []int, k int) (ans int) {
    n := len(a)
    if k >= (a[n-1] - a[0]) / 2 {
        return n
    }

    f := make([]int, n)
    for i, x := range a {
        if i > 0 && x == a[i-1] {
            f[i] = f[i-1]
        } else {
            pos := sort.SearchInts(a, x + k + 1)
            f[i] = pos - i
        }
    }

    suf := make([]int, n + 100)
    for i := n - 1; i >= 0; i-- {
        suf[i] = max(suf[i+1], f[i])
    }

    for i, x := range f {
        ans = max(ans, x + suf[i+x])
    }

    return
}