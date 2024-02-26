type pair struct {
    x, y int
}
func minAreaRect(points [][]int) int {
    m := map[pair]int{}
    for _, point := range points {
        m[pair{point[0], point[1]}] = 1
    }
    ans := math.MaxInt32
    for i := range points {
        a, b := points[i][0], points[i][1]
        for j := i + 1; j < len(points); j++ {
            c, d := points[j][0], points[j][1]
            if a == c || b == d {
                continue
            }
            if _, ok := m[pair{a, d}]; ok {
                if _, ok := m[pair{c, b}]; ok {
                    ans = min(ans, abs(a - c) * abs(b - d))
                }
            }
        }
    }
    if ans == math.MaxInt32 {
        return 0
    }
    return ans 
}

func abs(x int) int {
    if x < 0 {
        return -x 
    }
    return x
}