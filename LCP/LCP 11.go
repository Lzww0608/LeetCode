func expectNumber(scores []int) int {
    m := make(map[int]bool)
    for _, x := range scores {
        m[x] = true
    }

    return len(m)
}