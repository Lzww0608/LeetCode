func maxValue(events [][]int, k int) int {
    sort.Slice(events, func(i, j int) bool {
        return events[i][1] < events[j][1]
    })

    n := len(events)
    f := make([][]int, n + 1)
    for i := range f {
        f[i] = make([]int, k + 1)
    }
    for i, v := range events {
        a, c := v[0], v[2]
        
        mx := sort.Search(i, func(j int) bool {
            return events[j][1] >= a
        })
        for j := 1; j <= k; j++ {
            f[i+1][j] = max(f[i][j], f[mx][j-1] + c)
        }
    }
    
    return f[n][k]
}