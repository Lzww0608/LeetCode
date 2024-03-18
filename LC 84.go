func largestRectangleArea(heights []int) (ans int) {
    heights = append(heights, 0)
    n := len(heights)
    st := []int{}

    for i := 0; i < n; i++ {
        if len(st) == 0 || heights[i] >= heights[st[len(st)-1]] {
            st = append(st, i)
        } else {
            l, h := 0, heights[st[len(st)-1]]
            st = st[:len(st)-1]
            if len(st) == 0 {
                l = i
            } else {
                l = i - 1 - st[len(st)-1]
            }
            ans = max(ans, l * h)
            i--
        }

    }


    return
}