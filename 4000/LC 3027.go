func numberOfPairs(points [][]int) (ans int) {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] > points[j][1])
    })

    n := len(points)
    for i, v := range points {
        y := math.MinInt32
        for j := i + 1; j < n; j++ {
            cur_y := points[j][1]
            if cur_y > y && cur_y <= v[1] {
                ans++
                y = max(y, cur_y)
            } 
        }
    }

    return
}