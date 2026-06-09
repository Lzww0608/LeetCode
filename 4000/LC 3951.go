func minEnergy(n int, brightness int, intervals [][]int) int64 {
    ans := 0
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    cost := max(1, (brightness + 2) / 3)

    pre := -1
    for _, v := range intervals {
        ans += max(0, (v[1] - max(pre + 1, v[0]) + 1)) * cost
        pre = max(pre, v[1])
    }

    return int64(ans)
}