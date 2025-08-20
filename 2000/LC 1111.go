func maxDepthAfterSplit(seq string) []int {
    n := len(seq)
    ans := make([]int, n)
    a, b := 0, 0
    for i := range seq {
        if seq[i] == '(' {
            if a > b {
                ans[i] = 1
                b++
            } else {
                a++
            }
        } else {
            if a > b {
                a--
            } else {
                ans[i] = 1
                b--
            }
        }
    }

    return ans
}