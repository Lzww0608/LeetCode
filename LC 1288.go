func removeCoveredIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })

    ans := len(intervals)
    r := intervals[0][1]
    for i := 1; i < len(intervals); i++ {
        b := intervals[i][1]
        if r >= b {
            ans--
        } else if b > r {
            r = b
        }
    }

    return ans
}