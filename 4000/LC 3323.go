func minConnectedGroups(intervals [][]int, k int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1]
    })

    a := [][]int{intervals[0]}
    for _, v := range intervals {
        l, r := v[0], v[1]
        pre := a[len(a)-1]
        if l == pre[0] {
            continue
        }

        if l <= pre[1] {
            pre[1] = max(pre[1], r)
        } else {
            a = append(a, v)
        }
    }

    ans := 0
    n := len(a)
    for l, r := 0, 0; l < n; l++ {
        for r < n && a[r][0] - a[l][1] <= k {
            r++
        }
        ans = max(ans, r - l)
    } 

    return n - ans + 1
}