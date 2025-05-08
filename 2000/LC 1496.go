func isPathCrossing(path string) bool {
    m := map[[2]int]bool{}
    x, y := 0, 0
    m[[2]int{0, 0}] = true
    for i := range path {
        if path[i] == 'N' {
            y += 1
        } else if path[i] == 'S' {
            y -= 1
        } else if path[i] == 'E' {
            x += 1
        } else {
            x -= 1
        }

        t := [2]int{x, y}
        if m[t] {
            return true
        }
        m[t] = true
    }

    return false
}