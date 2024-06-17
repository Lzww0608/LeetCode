func shortestSequence(rolls []int, k int) int {
    m := make([]int, k + 1)
    ans, l := 1, k
    for _, v := range rolls {
        if m[v] < ans {
            m[v] = ans
            l--
            if l == 0 {
                l = k
                ans++
            }
        }
    }

    return ans
}
