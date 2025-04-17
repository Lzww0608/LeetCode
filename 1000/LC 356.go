
import "math"
func isReflected(points [][]int) bool {
    mn, mx := math.MaxInt, math.MinInt
    m := make(map[[2]int]bool)
    for _, v := range points {
        mn = min(mn, v[0])
        mx = max(mx, v[0])
        m[[2]int{v[0], v[1]}] = true
    }

    mid := mn + mx
    for _, v := range points {
        tmp := [2]int{mid - v[0], v[1]}
        if !m[tmp] {
            return false
        }
    }

    return true
}