func countGoodRectangles(rectangles [][]int) (ans int) {
    mx := 0
    for _, v := range rectangles {
        x := min(v[0], v[1])
        if x > mx {
            mx = x 
            ans = 1
        } else if x == mx {
            ans++
        }
    }

    return
}