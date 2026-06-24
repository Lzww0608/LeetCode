func maxDistance(moves string) int {
    x, y, cnt := 0, 0, 0
    for _, c := range moves {
        if c == 'U' {
            y++
        } else if c == 'D' {
            y--
        } else if c == 'L' {
            x--
        } else if c == 'R' {
            x++
        } else {
            cnt++
        }
    }

    return abs(x) + abs(y) + cnt
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}