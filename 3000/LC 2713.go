func maxIncreasingCells(mat [][]int) int {
    type pair struct {
        x, y int
    }
    g := map[int][]pair{}
    for i := range mat {
        for j, x := range mat[i] {
            g[x] = append(g[x], pair{i, j})
        }
    }

    a := make([]int, 0, len(g))
    for k := range g {
        a = append(a, k)
    }
    slices.Sort(a)

    rowMax := make([]int, len(mat))
    colMax := make([]int, len(mat[0]))
    for _, x := range a {
        pos := g[x]
        mx := make([]int, len(pos))
        for i, p := range pos {
            mx[i] = max(rowMax[p.x], colMax[p.y]) + 1
        }
        for i, p := range pos {
            rowMax[p.x] = max(rowMax[p.x], mx[i])
            colMax[p.y] = max(colMax[p.y], mx[i])
        }
    }

    return slices.Max(rowMax)
}