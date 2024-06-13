func minSwaps(grid [][]int) (ans int) {
    n := len(grid)

    a := make([]int, n)
    for i := range grid {
        pos := n
        for j := n - 1; j >= 0; j-- {
            if grid[i][j] == 0 {
                pos--
            } else {
                break
            }
        }

        a[i] = pos
    }

    for i, x := range a {
        if x <= i + 1 {
            continue
        }

        j := i + 1
        for j < n && a[j] > i + 1 {
            j++
        }

        if j == n {
            return -1
        }

        for t := j; t > i; t-- {
            a[t], a[t-1] = a[t-1], a[t]
            ans++
        }
    }

    return

}
