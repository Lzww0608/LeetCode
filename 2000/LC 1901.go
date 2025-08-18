func findPeakGrid(mat [][]int) []int {
    m := len(mat)
    l, r := -1, m - 1

    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        j := calc(mat[mid])
        if mat[mid][j] < mat[mid + 1][j] {
            l = mid
        } else {
            r = mid
        }
    }

    return []int{r, calc(mat[r])}
}

func calc(a []int) (j int) {
    for i, x := range a {
        if x > a[j] {
            j = i
        }
    }

    return 
}