func longestWPI(hours []int) int {
    n, ans := len(hours), 0
    pre := make([]int, n + 1)
    m := map[int]int{}
    m[0] = 0
    for i, x := range hours {
        if x > 8 {
            pre[i+1] = pre[i] + 1
        } else {
            pre[i+1] = pre[i] - 1
        }
        if v, ok := m[min(0, pre[i+1]-1)]; ok {
            ans = max(i+1 - v, ans);
        }
        if _, ok := m[pre[i+1]]; !ok {
            m[pre[i+1]] = i+1
        }
    }
    return ans;
}