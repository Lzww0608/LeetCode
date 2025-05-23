func maxWidthOfVerticalArea(points [][]int) (ans int) {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })

    pre := points[0][0]
    for _, v := range points[1:] {
        ans = max(ans, v[0] - pre)
        pre = v[0]
    }

    return
}