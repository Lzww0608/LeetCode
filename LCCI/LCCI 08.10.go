func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    oldColor := image[sr][sc]

    n, m := len(image), len(image[0])
    var dfs func(i, j int) 
    dfs = func(i, j int) {
        if i < 0 || i >= n || j < 0 || j >= m || image[i][j] != oldColor {
            return
        }
        image[i][j] = newColor
        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            dfs(x, y)
        }
    }
    if newColor != oldColor {
        dfs(sr, sc)
    }
    

    return image
}