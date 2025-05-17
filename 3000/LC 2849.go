func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    if sx == fx && sy == fy {
        return t != 1
    }

    return max(abs(sx - fx), abs(sy - fy)) <= t
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}