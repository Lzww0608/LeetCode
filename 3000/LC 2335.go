func fillCups(amount []int) (ans int) {
    a := slices.Max(amount)
    sum := amount[0] + amount[1] + amount[2] 

    if a > sum - a {
        return a
    }

    return (sum + 1) / 2
}