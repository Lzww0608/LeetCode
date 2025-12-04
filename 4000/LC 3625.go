func countTrapezoids(points [][]int) (ans int) {
    cnt1 := make(map[float64][]float64)
    cnt2 := make(map[int][]float64)

    for i, v := range points {
        x1, y1 := v[0], v[1]
        for j := i + 1; j < len(points); j++ {
            x2, y2 := points[j][0], points[j][1]
            dx, dy := x1 - x2, y1 - y2
            var k float64 = math.MaxFloat64
            var b float64 = float64(x1)
            if dx != 0 {
                k = float64(dy) / float64(dx)
                // y = kx + b, y = dy / dx * x + b 
                // y * dx = dy * x + b * dx 
                // b = (y * dx - x * dy) / dx
                b = float64(y1 * dx - x1 * dy) / float64(dx)
            }
            cnt1[k] = append(cnt1[k], b)
            mid := (x1 + x2 + 2000) << 16 + (y1 + y2 + 2000)
            cnt2[mid] = append(cnt2[mid], k)
        }
    }

    cnt := make(map[float64]int)
    for _, v := range cnt1 {
        if len(v) <= 1 {
            continue
        }

        clear(cnt)
        for _, x := range v {
            cnt[x]++
        }
        cur := 0
        for _, x := range cnt {
            ans += cur * x
            cur += x
        }
    }

    for _, v := range cnt2 {
        if len(v) <= 1 {
            continue
        }

        clear(cnt)
        for _, x := range v {
            cnt[x]++
        }
        cur := 0
        for _, x := range cnt {
            ans -= cur * x
            cur += x
        }
    }

    return
}