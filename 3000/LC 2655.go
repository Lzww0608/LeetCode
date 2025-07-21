func findMaximalUncoveredRanges(n int, ranges [][]int) (ans [][]int) {
    sort.Slice(ranges, func(i, j int) bool {
        if ranges[i][0] == ranges[j][0] {
            return ranges[i][1] < ranges[j][1]
        }
        return ranges[i][0] < ranges[j][0]
    })

    r := -1
    for _, v := range ranges {
        if v[0] > r + 1 {
            ans = append(ans, []int{r + 1, v[0] - 1})
        }
        r = max(r, v[1])
    }

    if r < n - 1 {
        ans = append(ans, []int{r + 1, n - 1})
    }

    return
}