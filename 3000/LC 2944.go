func minimumCoins(prices []int) int {
    n := len(prices)
    q := [][]int{{n + 1, 0}}
    for i := n; i > 0; i-- {
        for q[0][0] > i * 2 + 1 {
            q = q[1:]
        }
        f := prices[i-1] + q[0][1]
        for f <= q[len(q)-1][1] {  
            q = q[:len(q)-1]
        }
        q = append(q, []int{i, f})
    }
    return q[len(q)-1][1]
}