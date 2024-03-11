func nextGreaterElement(a []int, b []int) []int {
    n := len(a)
    m := map[int]int{}
    ans := make([]int, n)
    for i := range ans {
        ans[i] = -1
        m[a[i]] = i + 1
    }
    

    st := []int{}
    for i := range b {
        for len(st) > 0 && b[i] > b[st[len(st)-1]] {
            if m[b[st[len(st)-1]]] > 0 && ans[m[b[st[len(st)-1]]] - 1] == -1 {
                ans[m[b[st[len(st)-1]]] - 1] = b[i]
            }
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return ans
}