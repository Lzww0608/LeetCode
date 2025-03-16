func maxWeight(pizzas []int) int64 {
    sort.Ints(pizzas)
    n := len(pizzas)
    ans := 0

    i := n - 1
    for j := 0; j < n / 4; j += 2 {
        ans += pizzas[i]
        i--
    }

    for j := 0; j < n / 8; j++ {
        ans += pizzas[i-1]
        i -= 2
    }

    return int64(ans)
}