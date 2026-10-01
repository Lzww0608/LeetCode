func minQueenMoves(source []int, target []int) int {
    if source[0] == target[0] && source[1] == target[1] {
        return 0
    }
    if source[0] == target[0] || source[1] == target[1] || abs(source[0] - target[0]) == abs(source[1] - target[1]) {
        return 1
    }

    return 2
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}