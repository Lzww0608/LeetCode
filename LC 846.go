func isNStraightHand(hand []int, groupSize int) bool {
    n, k := len(hand), groupSize
    if n % k != 0 {
        return false
    }
    sort.Ints(hand)
    m := map[int]int{}
    for _, x := range hand {
        m[x]++
    }
    for i := 0; i < n; i++ {
        if m[hand[i]] == 0 {
            continue
        }
        for j := 0; j < k; j++ {
            if m[hand[i]+j] == 0 {
                return false
            }
            m[hand[i]+j]--
        }
    }
    return true
}