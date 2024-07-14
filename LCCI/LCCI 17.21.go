func trap(height []int) (ans int) {
    st := []int{}
    for i, x := range height {
        for len(st) > 0 && height[st[len(st)-1]] <= x {
            bottom := height[st[len(st)-1]]
            st = st[:len(st)-1]
            if len(st) == 0 {
                break
            }
            l := i - st[len(st)-1] - 1
            h := min(height[st[len(st)-1]], x) - bottom
            ans += h * l
        }
        st = append(st, i)
    }

    return
}