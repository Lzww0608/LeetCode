func bestHand(ranks []int, suits []byte) string {
    ans := "High Card"
    rankCnt := make(map[int]int)
    suitCnt := make(map[byte]int)
    for i := range len(ranks) {
        x := ranks[i]
        y := suits[i]
        rankCnt[x]++
        suitCnt[y]++
        if suitCnt[y] >= 5 {
            return "Flush"
        }
        if rankCnt[x] >= 3 {
            ans = "Three of a Kind"
        } else if ans != "Three of a Kind" && rankCnt[x] >= 2 {
            ans = "Pair"
        }
    }

    return ans
}