func maxCoins(piles []int) (ans int) {
    sort.Slice(piles, func(i, j int) bool {
        return piles[i] > piles[j]
    })
    for i := 1; i < len(piles) - i / 2; i += 2 {
        ans += piles[i]
    }

    return
}