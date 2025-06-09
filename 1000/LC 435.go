func eraseOverlapIntervals(intervals [][]int) (ans int) {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    cur := intervals[0][1]
    for _, v := range intervals[1:] {
        if v[0] < cur {
            ans++
            cur = min(v[1], cur)
        } else {
            cur = v[1]
        }
    }

    return
}