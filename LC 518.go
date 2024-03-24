func change(amount int, coins []int) int {
    f := make([]int, amount + 1)
    f[0] = 1


    for _, x := range coins {
        for j := x; j <= amount; j++ {
                f[j] += f[j - x]
        }
    }

    return f[amount]
}