func escapeGhosts(ghosts [][]int, target []int) bool {
    d := abs(target[0]) + abs(target[1])
    for _, g := range ghosts {
        if abs(g[1] - target[1]) + abs(g[0] - target[0]) <= d {
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
