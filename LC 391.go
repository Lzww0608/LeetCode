

func isRectangleCover(rectangles [][]int) bool {
    var sum int64 = 0
    min_x, min_y, max_a, max_b := math.MaxInt32, math.MaxInt32, 0, 0
    m := map[[2]int]int{} // 修改为map存储点
    for _, v := range rectangles {
        x, y, a, b := v[0], v[1], v[2], v[3]
        max_a = max(max_a, a)
        max_b = max(max_b, b)
        min_x = min(min_x, x)
        min_y = min(min_y, y)
        
        // 计算面积
        sum += int64(a - x) * int64(b - y)
        
        // 记录顶点
        points := [][2]int{{x, y}, {x, b}, {a, y}, {a, b}}
        for _, p := range points {
            m[p]++
            if m[p] % 2 == 0 {
                delete(m, p)
            }
        }
    }
    
    // 检查顶点数目
    if len(m) != 4 {
        return false
    }
    
    // 检查剩余的4个顶点是否是大矩形的顶点
    expectedPoints := [][2]int{{min_x, min_y}, {min_x, max_b}, {max_a, min_y}, {max_a, max_b}}
    for _, p := range expectedPoints {
        if m[p] != 1 {
            return false
        }
    }
    
    // 检查面积
    expectedArea := int64(max_a - min_x) * int64(max_b - min_y)
    return sum == expectedArea
}


