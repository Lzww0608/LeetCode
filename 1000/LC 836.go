func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    x0, y0, x1, y1 := rec1[0], rec1[1], rec1[2], rec1[3]
    x2, y2, x3, y3 := rec2[0], rec2[1], rec2[2], rec2[3]
    if x2 >= x1 || y2 >= y1 || x3 <= x0 || y3 <= y0 {
        return false
    }
    return true
}
