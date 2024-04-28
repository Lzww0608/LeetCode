func diagonalSort(mat [][]int) [][]int {
    m, n := len(mat), len(mat[0])
    a := make([]int, 0, n + m)

    store := func(i, j int) {
        for i < m && j < n {
            a = append(a, mat[i][j])
            i += 1
            j += 1
        }
        return 
    }

    update := func(i, j int) {
        k := 0
        for i < m && j < n {
            mat[i][j] = a[k]
            i, j, k = i + 1, j + 1, k + 1
        }
        return 
    }

    for i := range mat {
        a = make([]int, 0, n + m)
        store(i, 0)
        sort.Ints(a)
        update(i, 0)
    }

    for j := 1; j < n; j++ {
        a = make([]int, 0, n + m)
        store(0, j)
        sort.Ints(a)
        update(0, j)
    }

    return mat
}