func findMaxVal(n int, restrictions [][]int, diff []int) int {
    a := make([]int, n)
    rest := make([]int, n)
    for _, v := range restrictions {
        rest[v[0]] = v[1]
    }
    for i := 1; i < n; i++ {
        a[i] = a[i - 1] + diff[i - 1]
        if rest[i] != 0 {
            a[i] = min(a[i], rest[i])
        }
    }

    for i := n - 2; i > 0; i-- {
        a[i] = min(a[i], a[i + 1] + diff[i])
    }

    return slices.Max(a)
}
