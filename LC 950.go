func deckRevealedIncreasing(deck []int) []int {
    n := len(deck)
    ans := make([]int, n)
    dq := make([]int, n)
    for i := range dq {
        dq[i] = i
    }
    sort.Ints(deck)

    i := 0
    for len(dq) > 0 {
        t := dq[0]
        dq = dq[1:]
        ans[t] = deck[i]
        i++

        if len(dq) > 0 {
            tmp := dq[0]
            dq = dq[1:]
            dq = append(dq, tmp)
        } 
    }
    return ans
}