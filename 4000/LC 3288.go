func maxPathLength(coordinates [][]int, k int) int {
    kx, ky := coordinates[k][0], coordinates[k][1]
    sort.Slice(coordinates, func(i, j int) bool {
        return coordinates[i][0] < coordinates[j][0] || coordinates[i][0] == coordinates[j][0] && coordinates[i][1] > coordinates[j][1]
    })

    g := []int{}
    for _, v := range coordinates {
        x, y := v[0], v[1]
        if x < kx && y < ky || x > kx && y > ky {
            pos := sort.SearchInts(g, y)
            if pos < len(g) {
                g[pos] = y 
            } else {
                g = append(g, y)
            }
        }
    }

    return len(g) + 1
}