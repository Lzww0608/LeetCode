func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)
    st := []int{}
    for i, x := range temperatures {
        for len(st) > 0 && x > temperatures[st[len(st)-1]] {
            ans[st[len(st)-1]] = i - st[len(st)-1]
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return ans
}