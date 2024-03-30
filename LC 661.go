func imageSmoother(img [][]int) [][]int {
    m, n := len(img), len(img[0])
    pre := make([][]int, m + 1)
    for i := range pre {
        pre[i] = make([]int, n + 1)
    }
    for i, row := range img {
        for j, v := range row {
            pre[i+1][j+1] = v + pre[i+1][j] + pre[i][j+1] - pre[i][j] 
        }
    }

    for i, row := range img {
        for j := range row {
            a, b, c, d := max(0, i-1), max(0, j-1), min(m, i+2), min(n, j+2)
            cnt := (c - a) * (d - b)
            sum := pre[c][d] + pre[a][b] - pre[a][d] - pre[c][b]
            img[i][j] = sum / cnt
        }
    }

    return img
}