func canTransform(source []int, target []int) bool {
    sum := 0
    for i := range source {
        sum += source[i] - target[i]
    }

    return sum == 0
}