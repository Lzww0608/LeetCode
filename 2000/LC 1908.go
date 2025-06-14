func nimGame(piles []int) bool {
    x := 0
    for _, pile := range piles {
        x ^= pile
    }

    return x != 0
}