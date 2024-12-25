func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
    sort.Slice(horizontalCut, func(i, j int) bool {
        return horizontalCut[i] > horizontalCut[j]
    })
    sort.Slice(verticalCut, func(i, j int) bool {
        return verticalCut[i] > verticalCut[j]
    })
    i, j := 0, 0
    ans := 0
    for i < m - 1 || j < n - 1 {
        if j >= n - 1 || i < m - 1 && horizontalCut[i] >= verticalCut[j] {
            ans += horizontalCut[i] * (j + 1)
            i++
        } else {
            ans += verticalCut[j] * (i + 1)
            j++
        }
    }

    return int64(ans)
}