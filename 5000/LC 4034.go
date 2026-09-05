func minBishopMoves(source []int, target []int) int {
    x := abs(source[0] - target[0])
    y := abs(source[1] - target[1])
    if x == y {
        return 1
    } else if abs(x - y) % 2 == 0 {
        return 2
    }
    return -1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}