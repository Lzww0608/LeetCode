func dominantIndex(nums []int) int {
    a, b := -1, -1
    idx := -1
    for i, x := range nums {
        if x > a {
            a, b = x, a
            idx = i
        } else if x > b {
            b = x
        }
    }

    if a >= 2 * b {
        return idx
    }
    return -1
}