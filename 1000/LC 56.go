func merge(intervals [][]int) (ans [][]int) {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] != intervals[j][0] {
            return intervals[i][0] < intervals[j][0]
        }
        return intervals[i][1] < intervals[j][1]
    })
    l, r := intervals[0][0], intervals[0][1]
    for _, x := range intervals {
        a, b := x[0], x[1]
        if a > r {
            ans = append(ans, []int{l, r})
            l, r = a, b
        } else if a <= r {
            r = max(r, b)
        }
    }
    ans = append(ans, []int{l, r})
    return 
}
