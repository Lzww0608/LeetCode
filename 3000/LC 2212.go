func maximumBobPoints(numArrows int, aliceArrows []int) []int {
    n := len(aliceArrows)
    maxScore := 0
    ans := make([]int, n)
    for i := 1; i < (1 << n); i++ {
        score_Bob := 0
        cur := 0
        tmp := make([]int, n)
        for j, x := range aliceArrows {
            if i & (1 << j) != 0 {
                cur += x + 1
                tmp[j] = x + 1
                score_Bob += j
            } 
        }
        if cur <= numArrows && score_Bob > maxScore {
            tmp[0] += numArrows - cur
            maxScore = score_Bob
            copy(ans, tmp)
        }
    }

    return ans
}