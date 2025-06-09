func findMinArrowShots(points [][]int) (ans int) {
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })

    cur := math.MinInt32 - 1
    for _, v := range points {
        if v[0] > cur {
            ans++
            cur = v[1]
        } 
    }

    return
}