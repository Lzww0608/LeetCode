func matrixReshape(mat [][]int, r int, c int) [][]int {
    if r * c != len(mat) * len(mat[0]) {
        return mat
    }

    ans := make([][]int, r)
    for i := range ans {
        ans[i] = make([]int, c)
    }

    i, j := 0, 0
    for _, row := range mat {
        for _, x := range row {
            ans[i][j] = x
            j++
            if j == c {
                i, j = i + 1, 0
            }
        }
    }

    return ans
}