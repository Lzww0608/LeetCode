func constructRectangle(area int) []int {
    k := int(math.Sqrt(float64(area)))

    for area % k > 0 {
        k--
    }
    return []int{area / k, k}
}
