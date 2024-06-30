func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    m, n := len(image), len(image[0])
    origin := image[sr][sc]

    type pair struct {
        x, y int
    }

    q := []pair{pair{sr, sc}}
    for len(q) > 0 {
        i, j := q[0].x, q[0].y
        q = q[1:]
        image[i][j] = color
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            if x >= 0 && x < m && y >= 0 && y < n && image[x][y] == origin && image[x][y] != color {
                q = append(q, pair{x, y})
            } 
        }
    }

    return image
}
