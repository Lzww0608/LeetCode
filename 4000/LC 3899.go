func internalAngles(sides []int) []float64 {
    sort.Ints(sides)
    a, b, c := sides[0], sides[1], sides[2]
    if a + b <= c {
        return nil
    }

    radius := 180 / math.Pi
    A := math.Acos(float64(b * b + c * c - a * a) / float64(b * c * 2)) * radius
    B := math.Acos(float64(a * a + c * c - b * b) / float64(a * c * 2)) * radius
    return []float64{A, B, 180 - A - B}
}