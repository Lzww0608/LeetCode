func robot(command string, obstacles [][]int, x int, y int) bool {
    a, b := 0, 0
    for _, c := range command {
        if c == 'U' {
            b++
        } else {
            a++
        }
    }

    check := func(i, j int) bool {
        n := min(i/a, j/b)
        i -= n * a
        j -= n * b
        for _, c := range command {
            if i == 0 && j == 0 {
                return true
            }
            if c == 'U' {
                j--
            } else if c == 'R' {
                i--
            }
            if i < 0 || j < 0 {
                return false
            }
        }
        return i == 0 && j == 0
    }

    if !check(x, y) {
        return false
    }

    for _, obstacle := range obstacles {
        ox, oy := obstacle[0], obstacle[1]
        if ox <= x && oy <= y && check(ox, oy) {
            return false
        }
    }

    return true
}

