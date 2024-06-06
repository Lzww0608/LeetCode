func circularPermutation(n int, start int) []int {
    ans := make([]int, 1 << n)
    ans[0] = 0
    j := 0
    for i := 1; i < len(ans); i++ {
        ans[i] = i ^ (i >> 1)
        if ans[i] == start {
            j = i
        }
    }
    ans = append(ans, ans[:j]...)

    return ans[j:]
}