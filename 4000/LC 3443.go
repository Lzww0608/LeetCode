func maxDistance(s string, k int) (ans int) {
    x, y := 0, 0
    for i, c := range s {
        if c == 'N' {
            y++
        } else if c == 'S' {
            y--
        } else if c == 'E' {
            x++
        } else {
            x--
        }

        ans = max(ans, min(abs(x) + abs(y) + k * 2, i + 1))
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}