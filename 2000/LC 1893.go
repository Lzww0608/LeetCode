func isCovered(ranges [][]int, left int, right int) bool {
    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i][0] < ranges[j][0]
    })

    for _, v := range ranges {
        a, b := v[0], v[1]
        if a <= left {
            left = max(left, b + 1)
        }
    }

    return left > right
}