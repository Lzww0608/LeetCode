func isRobotBounded(instructions string) bool {
    x, y := 0, 0
    dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    k := 0
    for _, c := range instructions {
        if c == 'G' {
            x += dir[k][0]
            y += dir[k][1]
        } else if c == 'L' {
            k = (k + 1) % 4
        } else if c == 'R' {
            k = (k + 3) % 4
        }
    }

    return x == 0 && y == 0 || k != 0
}