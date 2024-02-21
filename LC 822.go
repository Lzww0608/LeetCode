func flipgame(fronts []int, backs []int) int {
    m := map[int]int{}
    for i := range fronts {
        if fronts[i] == backs[i] {
            m[fronts[i]] = math.MaxInt32
        }
    }
    ans := math.MaxInt32
    for i := range fronts {
        if backs[i] < ans && m[backs[i]] != math.MaxInt32 {
            ans = backs[i]
        }
        if fronts[i] < ans && m[fronts[i]] != math.MaxInt32 {
            ans = fronts[i]
        }
    }
    if ans == math.MaxInt32 {
        return 0
    }
    return ans
}