func minimumAddedCoins(coins []int, target int) (ans int) {
    sort.Ints(coins)
    s, i := 1, 0
    for s <= target {
        if i < len(coins) && coins[i] <= s {
            s += coins[i]
            i++
        } else {
            s *= 2
            ans++
        }
    }
    return
}