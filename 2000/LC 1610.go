import (
    "math"
    "sort"
)

const EPS float64 = 1e-8
func visiblePoints(points [][]int, angle int, location []int) (ans int) {
    same := 0
    a := []float64{}
    for _, p := range points {
        x, y := p[0], p[1]
        if x == location[0] && y == location[1] {
            same++
        } else {
            a = append(a, math.Atan2(float64(y - location[1]), float64(x - location[0])) * 180.0 / math.Pi)
        }
    }
    sort.Float64s(a)
    n := len(a)
    for i := 0; i < n; i++ {
        a = append(a, a[i] + 360.0)
    }

    for l, r := 0, 0; l < len(a); l++ {
        for r < len(a) && a[r] - a[l] <= float64(angle) + EPS {
            r++
        }

        ans = max(ans, r - l)
    }

    return ans + same
}