func bestLine(points [][]int) (ans []int) {
    n, mx := len(points), 0
    for i := range points {
        for j := i + 1; j < n; j++ {
            x1, y1 := points[i][0] - points[j][0], points[i][1] - points[j][1]
            cnt := 2
            for k := j + 1; k < n; k++ {
                x2, y2 := points[k][0] - points[j][0], points[k][1] - points[j][1]
                if x1 * y2 == x2 * y1 {
                    cnt++
                }
            }
            
            if cnt > mx {
                mx = cnt
                ans = []int{i, j}
            }
        }
        
    }

    return
}