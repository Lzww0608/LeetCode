func minTimeToVisitAllPoints(points [][]int) (ans int) {
    for i := 1; i < len(points); i++ {
        a, b := points[i-1][0], points[i-1][1]
        c, d := points[i][0], points[i][1]
        ans += max(abs(a - c), abs(b - d))
    }

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}