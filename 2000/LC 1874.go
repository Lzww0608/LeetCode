func minProductSum(a []int, b []int) (ans int) {
    sort.Ints(a)
    sort.Ints(b)

    n := len(a)
    for i := 0; i < n; i++ {
        ans += a[i] * b[n-i-1]
    }

    return
}