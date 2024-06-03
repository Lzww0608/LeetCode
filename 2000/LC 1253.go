func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
    n := len(colsum)
    ans := make([][]int, 2)
    ans[0], ans[1] = make([]int, n), make([]int, n)
    for i, x := range colsum {
        if x == 2 {
            ans[0][i], ans[1][i] = 1, 1
            upper--
            lower--
        } else if x == 1 {
            if upper > lower {
                ans[0][i] = 1
                upper--
            } else {
                lower--
                ans[1][i] = 1
            }
        }
    }


    if upper != 0 || lower != 0 {
        return nil
    }

    return ans
}
