func cycleLengthQueries(n int, queries [][]int) []int {
    results := make([]int, len(queries))
    for i, query := range queries {
        ai, bi := query[0], query[1]
        if ai > bi {
            ai, bi = bi, ai
        }
        lca := findLCA(ai, bi) + 1
        results[i] = lca
    }
    return results
}

func findLCA(x, y int) int {
    res := 0
    for x != y {
        if x > y {
            x /= 2
        } else {
            y /= 2
        }
        res++
    }
    return res
}

