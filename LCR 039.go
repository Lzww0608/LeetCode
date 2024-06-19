func largestRectangleArea(heights []int) (ans int) {
    st := []int{-1}
    heights = append(heights, 0)
    for i, x := range heights {
        for len(st) > 1 && heights[st[len(st)-1]] > x {
            h := heights[st[len(st)-1]]
            st = st[:len(st)-1]
            d := i - st[len(st)-1] - 1
            ans = max(ans, d * h)
        }
        st = append(st, i)
    }

    return
}