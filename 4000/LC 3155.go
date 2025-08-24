func maxUpgrades(count []int, upgrade []int, sell []int, money []int) []int {
    n := len(count)
    ans := make([]int, n)
    // x + y == count
    // x * upgrade <= money + sell * y
    // x * upgrade <= money + sell * count - sell * x
    // x * (upgrade + sell) <= money + sell * count
    // x <= (money + sell * count) / (upgrade + sell)
    for i := 0; i < n; i++ {
        x := (money[i] + sell[i] * count[i]) / (upgrade[i] + sell[i])
        ans[i] = min(x, count[i])
    }

    return ans
}