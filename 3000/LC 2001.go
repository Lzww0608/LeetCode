func interchangeableRectangles(rectangles [][]int) int64 {
    m := map[float64]int{}
    ans := 0

    for _, e := range rectangles {
        a, b := e[0], e[1]
        x := float64(a)/float64(b)
        ans += m[x]
        m[x]++
    }

    return int64(ans)
}