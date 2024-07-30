func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])
    })

    l := points[0][0]
    for _, v := range points[1:] {
        r := v[0]
        if r - l > w {
            ans++
            l = r
        }
    }

    return ans + 1
}