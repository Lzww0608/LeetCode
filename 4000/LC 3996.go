func canReach(start []int, target []int) bool {
    x := abs(start[0] - target[0]) + abs(start[1] - target[1])

    return x & 1 == 0
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}