func twoCitySchedCost(costs [][]int) (ans int) {
    n := len(costs)
    dif := make([]int, n)
    for i, v := range costs {
        a, b := v[0], v[1]
        ans += b
        dif[i] = a - b
    }
    sort.Ints(dif)
    for _, x := range dif[:n / 2] {
        ans += x
    }

    return 
}