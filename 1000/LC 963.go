
import "math"
func minAreaFreeRect(points [][]int) float64 {
    ans := math.MaxFloat32
    n := len(points)
    m := make(map[[3]float64][][]int)
    
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            x1, y1 := points[i][0], points[i][1]
            x2, y2 := points[j][0], points[j][1]
            mid_x := float64(x1 + x2) / 2
            mid_y := float64(y1 + y2) / 2
            d := float64((x2 - x1) * (x2 - x1) + (y2 - y1) * (y2 - y1))
            t := [3]float64{mid_x, mid_y, d}
            if p, ok := m[t]; ok {
                for _, v := range p {
                    x := math.Sqrt(float64((v[0] - x1) * (v[0] - x1) + (v[1] - y1) * (v[1] - y1)))
                    y := math.Sqrt(float64((v[0] - x2) * (v[0] - x2) + (v[1] - y2) * (v[1] - y2)))
                    ans = min(ans, x * y)
                }
                m[t] = append(m[t], []int{x1, y1})
            } else {
                m[t] = [][]int{[]int{x1, y1}}
            }
            
        }
    }

    if ans == math.MaxFloat32 {
        return 0
    }

    return ans
}