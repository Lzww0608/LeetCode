func findMaxValueOfEquation(points [][]int, k int) (ans int) {
    q := []int{}
    ans = math.MinInt32
    for i, v := range points {
        for len(q) > 0 && v[0] - points[q[0]][0] > k {
            q = q[1:]
        }
        if len(q) > 0 {
            ans = max(ans, points[q[0]][1] - points[q[0]][0] + v[0] + v[1])
        }
        for len(q) > 0 && points[q[len(q)-1]][1] - points[q[len(q)-1]][0] <= v[1] - v[0] {
            q = q[:len(q)-1]
        }
        q = append(q, i)
    }


    return
}