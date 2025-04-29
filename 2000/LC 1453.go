func numPoints(darts [][]int, r int) int {
    type Point struct {
        x, y float64
    }

    getDistance := func(a, b, c, d float64) float64 {
        return math.Sqrt((a - c) * (a - c) + (b - d) * (b - d))
    }

    solve := func(a, b Point) Point {
        mid := Point{(a.x + b.x) / 2, (a.y + b.y) / 2} 
        d := getDistance(a.x, a.y, mid.x, mid.y)
        h := math.Sqrt(float64(r * r) - d * d)
        ba := Point{b.x - a.x, b.y - a.y}
        hd := Point{-ba.y, ba.x}
        length := math.Sqrt(hd.x * hd.x + hd.y * hd.y)
        hd.x /= length
        hd.y /= length
        hd.x *= h 
        hd.y *= h 
        return Point{mid.x + hd.x, mid.y + hd.y}
    }

    n := len(darts)
    ans := 1

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            a := Point{float64(darts[i][0]), float64(darts[i][1])}
            b := Point{float64(darts[j][0]), float64(darts[j][1])}
            d := getDistance(a.x, a.y, b.x, b.y) 
            if d > float64(r) * 2 {
                continue
            }

            c := solve(a, b)
            cur := 0
            for _, v := range darts {
                d = getDistance(float64(v[0]), float64(v[1]), c.x, c.y)
                if d < float64(r) || math.Abs(d - float64(r)) < 1e-5 {
                    cur++
                }
            }
            ans = max(ans, cur)
        }
    }

    return ans
}