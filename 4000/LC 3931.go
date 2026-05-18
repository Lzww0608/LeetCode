func isAdjacentDiffAtMostTwo(s string) bool {
    for i := 1; i < len(s); i++ {
        x := int(s[i] - '0')
        y := int(s[i - 1] - '0')
        if abs(x - y) > 2 {
            return false
        }
    }

    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}