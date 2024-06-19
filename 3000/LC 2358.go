func maximumGroups(grades []int) int {
    n := len(grades)

    t := n * 2
    g := int(math.Sqrt(float64(t)))

    if g * (g + 1) > t {
        return g - 1
    }

    return g
}