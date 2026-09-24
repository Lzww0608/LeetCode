func countIntersectingIntervals(a [][]int) int64 {
    ans := 0
    sort.Slice(a, func(i, j int) bool {
        if a[i][1] == a[j][1] {
            return a[i][0] < a[j][0]
        }

        return a[i][1] < a[j][1]
    })

    n := len(a)
    for i, v := range a {
        j := sort.Search(i, func(k int) bool {
            return a[k][1] >= v[0]
        })

        ans -= j
    }
    ans += n * (n - 1) / 2

    return int64(ans)
}