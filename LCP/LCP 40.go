func maxmiumScore(cards []int, cnt int) int {
    sort.Slice(cards, func(i, j int) bool {
        return cards[i] > cards[j]
    })

    sum, n := 0, len(cards)
    for i := 0; i < cnt; i++ {
        sum += cards[i]
    }

    if sum & 1 == 0 {
        return sum
    }

    solve := func(x int) int {
        t, i := sum, cnt - 1
        for i = cnt - 1; i >= 0; i-- {
            if cards[i] & 1 == x {
                t -= cards[i]
                break
            }
        }
        if i == -1 {
            return 0
        }
        for i = cnt; i < n; i++ {
            if cards[i] & 1 == 1 - x {
                t += cards[i]
                return t
            }
        }
        return 0
    }

    
    return max(solve(0), solve(1))
}