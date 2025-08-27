func minMoves(rooks [][]int) (ans int) {
    sort.Slice(rooks, func(i, j int) bool {
        return rooks[i][0] < rooks[j][0]
    })
    for i, v := range rooks {
        ans += abs(v[0] - i)
        rooks[i][0] = i
    }

    sort.Slice(rooks, func(i, j int) bool {
        return rooks[i][1] < rooks[j][1]
    })
    for i, v := range rooks {
        ans += abs(v[1] - i)
    }

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}