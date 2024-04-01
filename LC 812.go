func largestTriangleArea(points [][]int) (ans float64) {
    n := len(points)

    for i := range points {
        for j := i + 1; j < n; j++ {
            for k := j + 1; k < n; k++ {
                x1, y1 := float64(points[i][0]), float64(points[i][1])
                x2, y2 := float64(points[j][0]) - x1, float64(points[j][1]) - y1    
                x3, y3 := float64(points[k][0]) - x1, float64(points[k][1]) - y1 
                ans = max(ans, abs(x2 * y3 - x3 * y2) / 2)   
            }
        }
    }

    return 
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}