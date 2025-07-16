func minimumCoins(prices []int) int {
    n := len(prices)
    type pair struct {
        i, x int 
    }

    q := []pair{{n + 1, 0}}
    for i := n; i > 0; i-- {
        for q[0].i > i * 2 + 1 {
            q = q[1:]
        }

        x := prices[i - 1] + q[0].x
        for len(q) > 0 && x <= q[len(q)-1].x {
            q = q[:len(q)-1]
        }
        q = append(q, pair{i, x})
    }

    return q[len(q)-1].x
}