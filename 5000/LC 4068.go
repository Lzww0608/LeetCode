func maxEarnings(meetings [][]int) int64 {
    ans := 0
    n := len(meetings)
    sort.Slice(meetings, func(i, j int) bool {
        return meetings[i][1] < meetings[j][1]
    })

    pre := make([]int, n + 1)
    pre[0] = math.MinInt
    pre_end := meetings[0][1]

    for i, v := range meetings {
        f := v[2]
        if v[0] >= pre_end {
            j := sort.Search(i, func(j int) bool {
                return meetings[j][1] > v[0] 
            })

            f += pre[j] + v[0]
        }

        ans = max(ans, f)
        pre[i + 1] = max(pre[i], f - v[1]) 
    }

    return int64(ans)
}