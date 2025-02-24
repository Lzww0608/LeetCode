func maxBuilding(n int, a [][]int) (ans int) {
    a = append(a, []int{1, 0})
    a = append(a, []int{n, n})
    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0]
    })

    m := len(a)
    for i := 1; i < m; i++ {
        a[i][1] = min(a[i][1], a[i-1][1] + a[i][0] - a[i-1][0])
    }
    for i := m - 2; i >= 0; i-- {
        a[i][1] = min(a[i][1], a[i+1][1] + a[i+1][0] - a[i][0])
    }

    for i := 0; i < m - 1; i++ {
        cur := ((a[i+1][0] - a[i][0]) + a[i][1] + a[i+1][1]) / 2
        ans = max(ans, cur)
    }

    return 
}