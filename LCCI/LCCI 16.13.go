const eps float64 = 1e-7
func cutSquares(square1 []int, square2 []int) []float64 {
    x1 := float64(square1[0]) + float64(square1[2]) / 2
    y1 := float64(square1[1]) + float64(square1[2]) / 2
    x2 := float64(square2[0]) + float64(square2[2]) / 2
    y2 := float64(square2[1]) + float64(square2[2]) / 2

    x_max := float64(max(square1[0] + square1[2], square2[0] + square2[2]))
    x_min := float64(min(square1[0], square2[0]))
    y_max := float64(max(square1[1] + square1[2], square2[1] + square2[2]))
    y_min := float64(min(square1[1], square2[1]))

    if math.Abs(x2 - x1) < eps {
        return []float64{x1, y_min, x1, y_max}
    } else if math.Abs(y2 - y1) < eps {
        return []float64{x_min, y1, x_max, y2}
    }

    k := (y2 - y1) / (x2 - x1)
    if k > 1 {
        return []float64{x2 - (y2 - y_min) / k, y_min, x2 - (y2 - y_max) / k, y_max}
    } else if k < -1 {
        return []float64{x2 - (y2 - y_max) / k, y_max, x2 - (y2 - y_min) / k, y_min}
    }
    return []float64{x_min,  y2 - k * (x2 - x_min), x_max, y2 - k * (x2 - x_max)}
}
